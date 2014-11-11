package main

import (
	"github.com/ReSc/fmt"
	. "github.com/ReSc/godel/core/reflect"
	"io"
	"sort"
	_ "text/template"
)

type (
	modelRenderer struct {
		*fmt.Writer
	}

	Renderer interface {
		Id() string
		Imports() []*Import
		RenderType(t *Type)
	}

	RendererFactory interface {
		Id() string
		NewRenderer(*fmt.Writer) Renderer
	}

	rendererFactory struct {
		id          string
		newRenderer func(*fmt.Writer) Renderer
	}
)

func (fac *rendererFactory) Id() string {
	return fac.id
}

func (fac *rendererFactory) NewRenderer(w *fmt.Writer) Renderer {
	return fac.newRenderer(w)
}

var (
	typeRenderers = make(map[string]RendererFactory)
)

func RegisterRendererFactory(id string, f func(*fmt.Writer) Renderer) {
	typeRenderers[id] = &rendererFactory{id: id, newRenderer: f}
}

func NewModelRenderer(w io.Writer) *modelRenderer {
	mr := &modelRenderer{
		fmt.NewWriter(w),
	}
	return mr
}

func (r *modelRenderer) Render(m *Model) {
	keys := m.Packages.Keys()
	sort.Strings(keys)
	for _, key := range keys {
		r.RenderPackage(m.Packages[key])
		r.Writeln("")
	}
}

func (r *modelRenderer) RenderPackage(p *Package) {
	r.Writeln("package %v", p.Name)
	r.ResolveRendererImports(p)
	r.RenderImports(p)
	r.RenderTypes(p)
}

func (r *modelRenderer) ResolveRendererImports(p *Package) {
	for _, t := range p.Types {
		if fac, ok := typeRenderers[t.Meta.Type]; ok {
			for _, imprt := range fac.NewRenderer(r.Writer).Imports() {
				p.Imports.Add(imprt)
			}
		}
	}
}

func (r *modelRenderer) RenderImports(p *Package) {
	r.Writeln("import (")
	keys := p.Imports.Keys()
	sort.Strings(keys)
	for _, key := range keys {
		r.RenderImport(p.Imports[key])
	}
	r.Writeln(")")
}

func (r *modelRenderer) RenderImport(i *Import) {
	r.Writeln("%v \"%v\"", i.Alias, i.Path)
}

func (r *modelRenderer) RenderTypes(p *Package) {
	keys := p.Types.Keys()
	sort.Strings(keys)
	for _, key := range keys {
		t := p.Types[key]
		r.RenderTypeTemplate(t.Meta.Type, t)
	}
}

func (r *modelRenderer) RenderTypeTemplate(templateId string, t *Type) {
	if fac, ok := typeRenderers[templateId]; ok {
		fac.NewRenderer(r.Writer).RenderType(t)
	} else {
		r.Writeln("// %v %v\n", t.Meta.Type, t.Name)
	}
}
