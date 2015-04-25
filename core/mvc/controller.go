package mvc

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

// ControllerFactory is a function that creates a single type of Controller
//
// usage:
// type MyController struct { ControllerBase }
// func (c *MyController) GetIndex() http.Handler {
//		return c.Json([]string{"Hello","World!"})
// }
//
// 	NewDispatcher(...).Register(func() Controller { return &MyController{} })
//
type ControllerFactory func() Controller

// Controller is an interface implemented by ControllerBase.
// It has a private method, so your controller must include
// ControllerBase as an anonymous field to implement this interface.
// see ControllerFactory documentation
type Controller interface {
	init(cc *controllerConfig, vars map[string]string, w http.ResponseWriter, r *http.Request)
}

type ViewEngine interface {
	Render(name string, model interface{}) error
}

// ControllerBase is the base implementation for mvc controllers
type ControllerBase struct {
	vars     map[string]string
	config   *controllerConfig
	Response http.ResponseWriter
	Request  *http.Request
}

func (c *ControllerBase) init(config *controllerConfig, vars map[string]string, w http.ResponseWriter, r *http.Request) {
	c.Response = w
	c.Request = r
	c.config = config
	c.vars = vars
}

// PathPrefix returns the basepath for all controllers
func (c *ControllerBase) PathPrefix() string {
	return c.vars["__prefix"]
}

// Path returns the basepath for this controller
func (c *ControllerBase) Path() string {
	return c.vars["__path"]
}

// Controller returns the controller path parameter
func (c *ControllerBase) Controller() string {
	return c.vars["controller"]
}

// Action return the action path parameter
func (c *ControllerBase) Action() string {
	return c.vars["action"]
}

// Id returns the Id path parameter
func (c *ControllerBase) Id() (string, bool) {
	return c.Var("id")
}

// IdAsInt64 returns the Id path parameter as an 64bit signed integer
func (c *ControllerBase) IdAsInt64() (int64, bool) {
	if id, ok := c.Var("id"); ok {
		if val, err := strconv.ParseInt(id, 10, 64); err == nil {
			return val, true
		}
	}
	return 0, false
}

// Var can be used to retrieve path parameters
func (c *ControllerBase) Var(name string) (string, bool) {
	val, ok := c.vars[name]
	return val, ok
}

// Get is the default GET implementation
func (c *ControllerBase) Get() http.Handler {
	return c.RedirectToMyAction("index")
}

// Index is the default GET index implementation
func (c *ControllerBase) GetIndex() http.Handler {
	actions := make(map[string][]string)
	for _, a := range c.config.actions.actions {
		aa := actions[a.httpMethod]
		aa = append(aa, a.name)
		actions[a.httpMethod] = aa
	}
	return c.ViewAt("", c.Action(), actions)
}

// Redirect ==============================

// RedirectTo returns a 303 client redirect to /prefix/controller/action/id path
func (c *ControllerBase) RedirectTo(controller, action, id string) http.Handler {
	parts := []string{c.PathPrefix(), controller, action, id}
	path := strings.Join(parts, "/")
	return &redirectHandler{path}
}

// RedirectToController returns a 303 client redirect to the /prefix/controller path
func (c *ControllerBase) RedirectToController(controller string) http.Handler {
	parts := []string{c.PathPrefix(), controller}
	path := strings.Join(parts, "/")
	return &redirectHandler{path}
}

// RedirectToAction returns a 303 client redirect to the /prefix/controller/action path
func (c *ControllerBase) RedirectToAction(controller, action string) http.Handler {
	parts := []string{c.PathPrefix(), controller, action}
	path := strings.Join(parts, "/")
	return &redirectHandler{path}
}

// RedirectToMyAction returns a 303 client redirect to the action on this controller
func (c *ControllerBase) RedirectToMyAction(action string) http.Handler {
	return c.RedirectToAction(c.Controller(), action)
}

// RedirectToUrl returns a 303 client redirect to the specified url
func (c *ControllerBase) RedirectToUrl(url string) http.Handler {
	return &redirectHandler{url}
}

type redirectHandler struct {
	url string
}

func (h *redirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.url, http.StatusSeeOther)
}

// View ==============================

func (c *ControllerBase) View(model interface{}) http.Handler {
	return c.ViewAt(c.Controller(), c.Action(), model)
}

func (c *ControllerBase) ViewAt(controller, action string, model interface{}) http.Handler {
	parts := []string{c.config.viewBasePath, ".."}

	if len(controller) > 0 {
		parts = append(parts, controller)
	}

	if len(action) > 0 {
		parts = append(parts, action+viewPostFix)
	} else {
		return c.InternalServerError("Empty template")
	}

	path := filepath.Join(parts...)
	path = filepath.Clean(path)

	if s, ok := model.(string); ok && s == "" {
		return c.File(path)
	} else {
		return c.parseAndRender(path, model)
	}
}

var viewPostFix = ".html"

func (c *ControllerBase) parseAndRender(path string, model interface{}) http.Handler {
	t, err := template.ParseFiles(path)
	if err != nil {
		return c.InternalServerError(err.Error())
	}

	return &viewHandler{
		view:  t,
		model: model,
	}
}

type viewHandler struct {
	model interface{}
	view  *template.Template
}

func (h *viewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := h.view.Execute(w, h.model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// File ==================

func (c *ControllerBase) File(path string) http.Handler {
	return &fileHandler{path}
}

type fileHandler struct {
	path string
}

func (h *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, h.path)
}

// Form ==================

var formdecoder = schema.NewDecoder()

func init() {
	// set the schema tag to json
	// so the decoder reuses json tags
	formdecoder.SetAliasTag("json")
	formdecoder.IgnoreUnknownKeys(true)
}

// FormBody decodes the request form data into v and returns an error if decoding failed
func (c *ControllerBase) FormBody(v interface{}) error {
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}
	return formdecoder.Decode(v, c.Request.Form)
}

// Json ==============================

// JsonBody decodes the request body into v and returns an error if decoding failed
func (c *ControllerBase) JsonBody(v interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

// Json returns a json serialized result of the supplied value.
// Also sets the correct Content-Type header
func (c *ControllerBase) Json(v interface{}) http.Handler {
	return &jsonHandler{v}
}

type jsonHandler struct {
	value interface{}
}

func (h *jsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.value)
}

// Xml ==============================

// XmBody decodes the request body into v and returns an error if this fails
func (c *ControllerBase) XmlBody(v interface{}) error {
	return xml.NewDecoder(c.Request.Body).Decode(v)
}

// Xml returns a xml serialized result of the supplied value.
// Also sets the correct Content-Type header
func (c *ControllerBase) Xml(v interface{}) http.Handler {
	return &xmlHandler{v}
}

type xmlHandler struct {
	value interface{}
}

func (h *xmlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(h.value)
}

// NotFound ==============================

// NotFound returns a 404 not found result
func (c *ControllerBase) NotFound() http.Handler {
	return c.Status(http.StatusNotFound, "")
}

// BadRequest ==============================

// BadRequest bad request
func (c *ControllerBase) BadRequest() http.Handler {
	return c.Status(http.StatusBadRequest, "")
}

// Internal Server Error ==============================

// InternalServerError returns a 500 internal server error with a custom error message
func (c *ControllerBase) InternalServerError(err string) http.Handler {
	return c.Status(http.StatusInternalServerError, err)
}

// Status ==============================

// Status returns a custom http Status result.
func (c *ControllerBase) Status(code int, message string) http.Handler {
	if message == "" {
		message = http.StatusText(code)
	}
	return &statusHandler{code, message}
}

type statusHandler struct {
	code    int
	message string
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, h.message, h.code)
}
