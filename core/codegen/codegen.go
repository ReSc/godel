package codegen

import (
	graph "github.com/ReSc/godel/core/graph"
	reflect "github.com/ReSc/godel/core/reflect"
)

// ModelGraphBuilder is a struct
type ModelGraphBuilder struct {
	Graph *graph.Graph
}

// NewModelGraphBuilder creates a new instance of ModelGraphBuilder
func NewModelGraphBuilder() *ModelGraphBuilder {
	return &ModelGraphBuilder{
		Graph: graph.NewGraph(),
	}
}

func (this *ModelGraphBuilder) VisitField(f *reflect.Field) {
	this.doVisitField(f)
}
func (this *ModelGraphBuilder) VisitMethod(m *reflect.Method) {
	this.doVisitMethod(m)
}
func (this *ModelGraphBuilder) VisitModel(m *reflect.Model) {
	this.doVisitModel(m)
}
func (this *ModelGraphBuilder) VisitPackage(p *reflect.Package) {
	this.doVisitPackage(p)
}
func (this *ModelGraphBuilder) VisitType(t *reflect.Type) {
	this.doVisitType(t)
}
