package codegen

import (
	graph "github.com/ReSc/godel/core/graph"
	reflect "github.com/ReSc/godel/core/reflect"
)

func (this *ModelGraphBuilder) doVisitModel(m *reflect.Model) {
	root := this.Graph.Nodes[0]
	p := this.Graph.NewPredicate("model")
	n := this.Graph.NewNode(m.Name)
	n.Tags.Add("model")
	n.Attrs.Set("path", m.Path)
	this.Graph.AddEdge(root, p, n)
}

func (this *ModelGraphBuilder) doVisitPackage(p *reflect.Package) {
	root := this.Graph.Nodes[0]
	p := this.Graph.NewPredicate("model")
	n := this.Graph.NewNode(m.Name)
	n.Tags.Add("model")
	n.Attrs.Set("path", m.Path)
	this.Graph.AddEdge(root, p, n)
}

func (this *ModelGraphBuilder) doVisitType(t *reflect.Type) {

}

func (this *ModelGraphBuilder) doVisitField(f *reflect.Field) {

}

func (this *ModelGraphBuilder) doVisitMethod(m *reflect.Method) {

}

func (g *ModelGraphBuilder) AddNamespace(parent *graph.Node, name string) *graph.Node {
	if parent.Tags.Contains("namespace") || parent.Tags.Contains("root") {

		prd := g.Graph.NewPredicate("namespace")

		ns := g.Graph.NewNode(name)
		ns.Tags.Add("namespace")

		g.Graph.AddEdge(parent, ns, prd)
		return ns
	}

	panic("invalid parent")
	return nil
}
