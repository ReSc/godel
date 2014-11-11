package rest

import (
	"encoding/xml"
	"errors"
	"github.com/ReSc/fmt"
	"github.com/ReSc/godel/core/codecs"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type (
	Params map[string]interface{}

	Endpoint interface {
		Getter
		Putter
		Poster
		Deleter
		Patcher
		Tracer
		Optionser
		Resourcer
	}

	Getter interface {
		Get(w http.ResponseWriter, r *http.Request)
	}

	Putter interface {
		Put(w http.ResponseWriter, r *http.Request)
	}

	Poster interface {
		Post(w http.ResponseWriter, r *http.Request)
	}

	Deleter interface {
		Delete(w http.ResponseWriter, r *http.Request)
	}

	Patcher interface {
		Patch(w http.ResponseWriter, r *http.Request)
	}

	Tracer interface {
		Trace(w http.ResponseWriter, r *http.Request)
	}

	Optionser interface {
		Options(w http.ResponseWriter, r *http.Request)
	}

	Resourcer interface {
		GetResource() *Resource
	}

	Resource struct {
		route *mux.Route
	}

	resource struct {
		e       interface{}
		formats []string
		r       *Resource
	}
)

func newResource(e interface{}) *resource {
	formats := codecs.AvailableFormats()
	res := &resource{
		e:       e,
		formats: formats,
	}
	if r, ok := e.(Resourcer); ok && r.GetResource() != nil {
		res.r = r.GetResource()
	} else {
		panic("not a resource")
	}
	return res
}

func (res *resource) GetResource() *Resource {
	return res.r
}

func (res *resource) Get(w http.ResponseWriter, r *http.Request) {
	res.setCacheHeaders(w)
	if e, ok := res.e.(Getter); ok {
		e.Get(w, r)
	} else {
		http.Error(w, "GET not allowed", http.StatusMethodNotAllowed)
	}
}

func (res *resource) Post(w http.ResponseWriter, r *http.Request) {
	res.setCacheHeaders(w)
	if e, ok := res.e.(Poster); ok {
		e.Post(w, r)
	} else {
		http.Error(w, "POST not allowed", http.StatusMethodNotAllowed)
	}
}

func (res *resource) Put(w http.ResponseWriter, r *http.Request) {
	res.setCacheHeaders(w)
	if e, ok := res.e.(Putter); ok {
		e.Put(w, r)
	} else {
		http.Error(w, "PUT not allowed", http.StatusMethodNotAllowed)
	}
}

func (res *resource) Delete(w http.ResponseWriter, r *http.Request) {
	res.setCacheHeaders(w)
	if e, ok := res.e.(Deleter); ok {
		e.Delete(w, r)
	} else {
		http.Error(w, "DELETE not allowed", http.StatusMethodNotAllowed)
	}
}

func (res *resource) Patch(w http.ResponseWriter, r *http.Request) {
	res.setCacheHeaders(w)
	if e, ok := res.e.(Patcher); ok {
		e.Patch(w, r)
	} else {
		http.Error(w, "PATCH not allowed", http.StatusMethodNotAllowed)
	}
}

func (res *resource) Options(w http.ResponseWriter, r *http.Request) {
	res.setCacheHeaders(w)

	w.Header().Set("Server", "godel runtime:"+runtime.Version()+" compiler:"+runtime.Compiler+" os:"+runtime.GOOS+" arch:"+runtime.GOARCH)

	res.setCorsHeaders(w, r)

	w.Header().Set("Allow", strings.Join(res.getSupportedMethods(), ","))
	if e, ok := res.e.(Optionser); ok {
		e.Options(w, r)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (res *resource) Trace(w http.ResponseWriter, r *http.Request) {
	res.setCacheHeaders(w)
	if e, ok := res.e.(Tracer); ok {
		e.Trace(w, r)
	} else {
		r.Write(w)
		http.Error(w, "", http.StatusOK)
	}
}

func (res *resource) setCorsHeaders(w http.ResponseWriter, r *http.Request) {
	supportedMethods := res.getSupportedMethods()

	// CORS allow origin
	origin := r.Header.Get("Origin")
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)

		// CORS allow request method
		requestedMethod := r.Header.Get("Access-Control-Request-Method")
		for _, method := range supportedMethods {
			if requestedMethod == method {
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(supportedMethods, ","))
				break
			}
		}

		// CORS allow headers
		exposedHeaders, ok := r.Header["Access-Control-Expose-Headers"]
		if ok && len(exposedHeaders) > 0 {
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(exposedHeaders, ","))
		}

		// CORS allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

func (res *resource) setCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Cache-Control", "no-cache")
}

func (res *resource) getSupportedMethods() []string {
	methods := make([]string, 0, 8)

	if _, ok := res.e.(Getter); ok {
		methods = append(methods, "GET")
	}

	if _, ok := res.e.(Putter); ok {
		methods = append(methods, "PUT")
	}

	if _, ok := res.e.(Poster); ok {
		methods = append(methods, "POST")
	}

	if _, ok := res.e.(Patcher); ok {
		methods = append(methods, "PATCH")
	}

	if _, ok := res.e.(Deleter); ok {
		methods = append(methods, "DELETE")
	}

	methods = append(methods, "OPTIONS")
	methods = append(methods, "TRACE")
	return methods
}

func Register(r *mux.Router, res interface{}, idFormat string) {
	register(r, res, idFormat)
}

func register(r *mux.Router, res interface{}, idFormat string) {
	resourceName := getResourceName(res)
	e := newResource(res)
	// collection endpoint
	registerMethods(r, "/"+resourceName, e)

	r = r.PathPrefix("/" + resourceName + "/").Subrouter()
	// element endpoint
	registerMethods(r, "/{"+idFormat+"}", e)
}

func registerMethods(r *mux.Router, path string, res Endpoint) *mux.Router {
	res.GetResource().route = r.Path(path).BuildOnly()
	r.HandleFunc(path, res.Get).Methods("GET")
	r.HandleFunc(path, res.Put).Methods("PUT")
	r.HandleFunc(path, res.Post).Methods("POST")
	r.HandleFunc(path, res.Patch).Methods("PATCH")
	r.HandleFunc(path, res.Trace).Methods("TRACE")
	r.HandleFunc(path, res.Delete).Methods("DELETE")
	r.HandleFunc(path, res.Options).Methods("OPTIONS")

	return r
}

func (res *Resource) GetResource() *Resource {
	return res
}

func (res *Resource) Error(w http.ResponseWriter, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
}

func (res *Resource) Path(params Params) (string, error) {
	pairs := make([]string, 0, len(params)*2)
	for k, v := range params {
		pairs = append(pairs, k)
		pairs = append(pairs, fmt.String("%v", v))
	}

	url, err := res.route.URLPath(pairs...)
	if err != nil {
		return "", err
	}

	return url.Path, err
}

func (res *Resource) Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (res *Resource) Ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (res *Resource) Found(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}

func (res *Resource) SeeOther(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusSeeOther)
}

func (res *Resource) DecodeBody(w http.ResponseWriter, r *http.Request, body interface{}) error {
	contentType, _ := getContentTypeHeader(r)

	if strings.Contains(contentType, "application/json") {
		return decodeRequest(w, r, "json", body)
	}

	if strings.Contains(contentType, "application/xml") {
		return decodeRequest(w, r, "xml", body)
	}

	msg := "Content type not supported " + contentType
	res.Error(w, msg, http.StatusUnsupportedMediaType)
	return errors.New(msg)
}

func (res *Resource) Return(w http.ResponseWriter, r *http.Request, response interface{}) error {
	contentType, hasWildcard := getAcceptHeader(r)

	if hasWildcard || strings.Contains(contentType, "application/json") {
		return encodeResponse(w, "json", response)
	}

	if strings.Contains(contentType, "application/xml") {
		return encodeResponse(w, "xml", response)
	}

	msg := "No acceptable mime-type in " + contentType
	res.Error(w, msg, http.StatusNotAcceptable)
	return errors.New(msg)
}

func getAcceptHeader(r *http.Request) (string, bool) {
	contentType, hasWildcard := "", true
	format := r.URL.Query().Get("format")
	switch format {
	case "":
		contentType = r.Header.Get("Accept")
		hasWildcard = isWildcardAccept(contentType)
	default:
		contentType = format
		hasWildcard = false
	}

	return contentType, hasWildcard
}

func getContentTypeHeader(r *http.Request) (string, bool) {
	contentType, queryOverride := "", false
	format := r.URL.Query().Get("format")
	switch format {
	case "":
		contentType = r.Header.Get("Content-Type")
		queryOverride = false
	default:
		contentType = format
		queryOverride = true
	}

	return contentType, queryOverride
}

func getResourceName(res interface{}) string {
	x := reflect.Indirect(reflect.ValueOf(res))
	name := x.Type().Name()
	name = strings.TrimSuffix(name, "Resource")
	return name
}

func decodeRequest(w http.ResponseWriter, r *http.Request, format string, body interface{}) error {
	err := error(nil)
	if factory, ok := codecs.Format(format); ok {
		err = factory.NewDecoder(r.Body).Decode(body)
		if err != nil {
			msg := stripSource(err)
			http.Error(w, msg, http.StatusBadRequest)
		}
	} else {
		msg := "Unsupported format: " + format
		err = errors.New(msg)
		http.Error(w, msg, http.StatusInternalServerError)
	}
	return err
}

func encodeResponse(w http.ResponseWriter, format string, body interface{}) error {
	err := error(nil)
	if factory, ok := codecs.Format(format); ok {
		w.Header().Set("Content-Type", "application/"+format)
		if format == "xml" {
			type Response struct {
				XMLName xml.Name    `xml:"response"`
				Data    interface{} `xml:"body"`
			}
			body = Response{Data: body}
		}
		err = factory.NewEncoder(w).Encode(body)
		if err != nil {
			msg := stripSource(err)
			http.Error(w, msg, http.StatusInternalServerError)
		}
	} else {
		msg := "Unsupported format: " + format
		err = errors.New(msg)
		http.Error(w, msg, http.StatusInternalServerError)
	}
	return err
}

func isWildcardAccept(accept string) bool {
	return strings.Contains(accept, "application/*") ||
		strings.Contains(accept, "*/*")
}

func stripSource(e error) string {
	msg := e.Error()
	i := strings.Index(msg, ":")
	if i > 0 {
		return msg[i+len(":"):]
	}
	return msg
}
