package main

import (
	"github.com/ReSc/godel/core/graph"
	"github.com/ReSc/godel/core/mvc"
	"github.com/ReSc/godel/core/reflect"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", router())
}

func router() *mux.Router {
	r := mux.NewRouter()
	serveFilesFrom(r, "./public", "/app/")
	serveControllersFrom(r, "/api/v1")
	return r
}

func serveFilesFrom(r *mux.Router, fspath, prefix string) {
	dir := http.Dir(fspath)
	fileServer := http.StripPrefix(prefix, http.FileServer(dir))
	handler := mvc.NewHandlerBuilder().
		Build(&mvc.Route{
		Name:     "StaticFiles",
		Method:   "GET",
		Pipeline: []string{"log"},
		Handler:  fileServer,
	})
	r.
		Methods(handler.Method).
		PathPrefix(prefix).
		Handler(handler)
}

func serveControllersFrom(r *mux.Router, prefix string) {
	d := mvc.NewDispatcher(mvc.DispatcherConfig{
		PathPrefix: prefix,
		Router:     r,
	})

	m, err := reflect.LoadModelFile("./data/model.xml")
	if err != nil {
		panic(err)
	}

	d.Register(func() mvc.Controller {
		return &ModelController{model: m}
	})

	g := graph.NewGraph().Init()
	g.Root.Attrs.Set("type", "NODE")
	g.Root.Attrs.Set("id", "ROOT")
	d.Register(func() mvc.Controller {
		return &GraphController{graph: g}
	})
}
