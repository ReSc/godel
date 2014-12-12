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
	convertModelToGraph(g, m)
	d.Register(func() mvc.Controller {
		return &GraphController{graph: g}
	})
}

func convertModelToGraph(g *graph.Graph, m *reflect.Model) {
	createModelNode(g, m)
}

func createModelNode(g *graph.Graph, m *reflect.Model) {
	n := g.NewNode(m.Name)
	n.Tags.Add("model")
	n.Attrs.Set("path", m.Path)
	g.AddNode(n)
	g.AddEdge(g.Nodes[0], getModelPredicate(g), n)

	for _, pkg := range m.Packages {
		createPackageNode(g, n, pkg)
	}
}

func createPackageNode(g *graph.Graph, modelNode *graph.Node, pkg *reflect.Package) {
	n := g.NewNode(pkg.Name)
	n.Tags.Add("package")
	n.Attrs.Set("path", pkg.GetPackagePath())
	g.AddNode(n)
	g.AddEdge(modelNode, getPackagePredicate(g), n)
	for _, t := range pkg.Types {
		createTypeNode(g, n, t)
	}
	for _, i := range pkg.Imports {
		createImportNode(g, n, i)
	}
}

func createTypeNode(g *graph.Graph, packageNode *graph.Node, typ *reflect.Type) {
	n := g.NewNode(typ.Name)
	n.Tags.Add("type")
	n.Attrs.Set("meta-type", typ.MetaType)
	g.AddNode(n)
	g.AddEdge(packageNode, getTypePredicate(g), n)

	for _, f := range typ.Fields {
		createFieldNode(g, n, f)
	}
}

func createImportNode(g *graph.Graph, packageNode *graph.Node, i *reflect.Import) {
	n := g.NewNode(i.Alias)
	n.Tags.Add("import")
	n.Attrs.Set("path", i.Path)
	g.AddNode(n)
	g.AddEdge(packageNode, getImportPredicate(g), n)
}

func createFieldNode(g *graph.Graph, typeNode *graph.Node, field *reflect.Field) {
	n := g.NewNode(field.Name)
	n.Tags.Add("field")
	n.Attrs.Set("data-type", field.DataType)
	g.AddNode(n)
	g.AddEdge(typeNode, getFieldPredicate(g), n)
}

func getFieldPredicate(g *graph.Graph) *graph.Node {
	return getPredicate(g, "field")
}

func getTypePredicate(g *graph.Graph) *graph.Node {
	return getPredicate(g, "type")
}

func getImportPredicate(g *graph.Graph) *graph.Node {
	return getPredicate(g, "import")
}

func getPackagePredicate(g *graph.Graph) *graph.Node {
	return getPredicate(g, "package")
}

func getModelPredicate(g *graph.Graph) *graph.Node {
	return getPredicate(g, "model")
}

var predicates = make(map[string]*graph.Node)

func getPredicate(g *graph.Graph, name string) *graph.Node {
	if p, ok := predicates[name]; ok {
		return p
	}

	tag := "predicate"
	/*
		for _, n := range g.Nodes {
			if n.Tags.Contains(tag) && n.Name() == name {
				predicates[name] = n
				return n
			}
		}
	*/
	newPred := g.NewNode(name)
	newPred.Tags.Add(tag)
	g.AddNode(newPred)
	predicates[name] = newPred
	return newPred
}
