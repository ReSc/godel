package mvc

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
	"unicode"
)

type DispatcherConfig struct {
	PathPrefix            string
	ViewRootPath          string
	Router                *mux.Router
	RecognizedHttpMethods []string
	Log                   *log.Logger
}

type Dispatcher struct {
	mutex       sync.Mutex
	controllers map[string]*controllerConfig
	config      DispatcherConfig
	router      *mux.Router
	log         *log.Logger
	templates   *template.Template
}

func NewDispatcher(config DispatcherConfig) *Dispatcher {
	d := &Dispatcher{
		mutex:       sync.Mutex{},
		controllers: make(map[string]*controllerConfig),
		config:      config,
		router:      config.Router,
	}

	d.initLogger()
	// and init the rest
	d.initViewEngine()
	d.initSupportedHttpMethods()
	d.initRouter()

	return d
}

func (d *Dispatcher) initLogger() {
	if d.config.Log == nil {
		d.config.Log = log.New(os.Stdout, "mvc - ", log.LstdFlags|log.Lmicroseconds)
	}
	d.log = d.config.Log
}

func (d *Dispatcher) initViewEngine() {
	d.config.ViewRootPath = "./views"
}
func (d *Dispatcher) initSupportedHttpMethods() {
	if len(d.config.RecognizedHttpMethods) == 0 {
		d.config.RecognizedHttpMethods = []string{
			"GET",
			"PUT",
			"POST",
			"HEAD",
			"PATCH",
			"TRACE",
			"DELETE",
			"OPTIONS",
			"CONNECT",
		}
	} else {
		for i, m := range d.config.RecognizedHttpMethods {
			d.config.RecognizedHttpMethods[i] = strings.ToUpper(m)
		}
	}
}

func (d *Dispatcher) initRouter() {
	if d.config.Router == nil {
		d.config.Router = mux.NewRouter()
		d.router = d.config.Router
	}

	prefix := d.config.PathPrefix
	parts := strings.Split("/{controller}/{action}/{id}", "/")
	l := len(parts)
	for i := range parts {
		path := strings.Join(parts[:l-i], "/")
		d.router.PathPrefix(prefix).Path(path).HandlerFunc(d.dispatch)
	}
}

func (d *Dispatcher) ListenAndServe(addr string) error {
	d.log.Println("Listening on: ", addr)
	return http.ListenAndServe(addr, d.router)
}

func (d *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.router.ServeHTTP(w, r)
}

func (d *Dispatcher) Register(factory ControllerFactory) {
	cc := &controllerConfig{
		factory: factory,
		actions: newActionMap(),
	}

	d.register(cc)
	d.log.Println("Registered controller ", cc.name)
}

type controllerConfig struct {
	name         string
	actions      *actionMap
	factory      ControllerFactory
	viewBasePath string
}

func (d *Dispatcher) register(cc *controllerConfig) {
	c := cc.factory()
	name := getControllerName(c)
	cc.name = name
	cc.viewBasePath = d.config.ViewRootPath + "/" + name

	methods := getControllerMethods(c)
	for _, m := range methods {
		for _, httpMethod := range d.config.RecognizedHttpMethods {
			upperName := strings.ToUpper(m.Name)
			if strings.HasPrefix(upperName, httpMethod) {
				name := m.Name[len(httpMethod):]
				name = toSpinalCase(name)
				if cc.actions.Add(c, httpMethod, name, m.Func) {
					if name != "" {
						name = "/" + name
					}
					d.log.Printf("Registered %s %s/%s%s\n", httpMethod, d.config.PathPrefix, cc.name, name)
				}
			}
		}
	}

	d.mutex.Lock()
	d.controllers[cc.name] = cc
	d.mutex.Unlock()
}

func (d *Dispatcher) dispatch(w http.ResponseWriter, r *http.Request) {
	tstart := time.Now()

	v := mux.Vars(r)

	d.mutex.Lock()
	ctrl, ok := d.controllers[v["controller"]]
	d.mutex.Unlock()

	if !ok {
		d.log.Println("404 controller not found ", v["controller"])
		http.NotFound(w, r)
		return
	}

	action, ok := ctrl.actions.Get(r.Method, v["action"])
	if !ok {
		d.log.Println("404 action not found ", ctrl.name, " ", v["action"])
		http.NotFound(w, r)
		return
	}

	c := ctrl.factory()
	v["__prefix"] = d.config.PathPrefix
	v["__path"] = d.config.PathPrefix + "/" + ctrl.name

	rw := NewResponseWriter(w)

	c.init(ctrl, v, rw, r)

	tresolved := time.Now()

	result := action.Invoke(c)

	tinvoked := time.Now()

	result.ServeHTTP(rw, r)

	tserved := time.Now()

	d.log.Printf("[%s - %d - %s] total:%v; resolve:%v; invoke:%v; serve:%v;",
		r.Method,
		rw.Status(),
		r.URL.Path,
		tserved.Sub(tstart),
		tresolved.Sub(tstart),
		tinvoked.Sub(tresolved),
		tserved.Sub(tinvoked))
}

var (
	handlerInterface = reflect.TypeOf(new(http.Handler)).Elem()
)

func getControllerName(c interface{}) string {
	val := reflect.ValueOf(c)
	val = reflect.Indirect(val)
	typeName := val.Type().Name()
	name := strings.TrimSuffix(typeName, "Controller")
	return toSpinalCase(name)
}

func getControllerMethods(c interface{}) []reflect.Method {
	ctrlType := reflect.TypeOf(c)

	methods := make([]reflect.Method, 0, ctrlType.NumMethod())
	for i := 0; i < ctrlType.NumMethod(); i++ {
		methods = append(methods, ctrlType.Method(i))
	}

	return methods
}

func toSpinalCase(s string) string {
	words := make([]string, 0, 4)
	for len(s) > 0 {
		word := firstWord(s)
		words = append(words, strings.ToLower(word))
		s = s[len(word):]
	}
	return strings.Join(words, "-")
}

func firstWord(s string) string {
	var prev rune
	for i, curr := range s {
		if i > 0 && unicode.IsUpper(curr) && unicode.IsLower(prev) {
			return s[:i]
		}
		prev = curr
	}
	return s
}
