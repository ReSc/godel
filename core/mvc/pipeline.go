package mvc

import (
	"net/http"
)

type Route struct {
	http.Handler
	Name     string
	Method   string
	Path     string
	Pipeline []string
}

type HandlerFactory interface {
	Name() string
	Build(http.Handler) http.Handler
}

type HandlerBuilder struct {
	factories map[string]HandlerFactory
}

func NewHandlerBuilder() *HandlerBuilder {
	b := &HandlerBuilder{
		factories: make(map[string]HandlerFactory),
	}
	b.Register(&logHandler{})
	return b
}

func (b *HandlerBuilder) Register(fac HandlerFactory) {
	name := fac.Name()
	if _, ok := b.factories[name]; ok {
		panic("Duplicate handler factory name: " + name)
	} else {
		b.factories[name] = fac
	}
}

func (b *HandlerBuilder) Build(r *Route) *Route {
	i := len(r.Pipeline) - 1
	h := r.Handler
	for offset := range r.Pipeline {
		name := r.Pipeline[i-offset]
		if fac, ok := b.factories[name]; ok {
			h = fac.Build(h)
		} else {
			panic("Unknown pipeline segment: " + name)
		}
	}
	r.Handler = h
	return r
}
