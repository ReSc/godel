package main

import (
	"github.com/ReSc/fmt"
	. "github.com/ReSc/godel/core/reflect"
	_ "sort"
	_ "strings"
	"text/template"
)

func init() {
	tpl, err := template.ParseGlob("templates\\go\\*")
	PanicIf(err)
	for _, t := range tpl.Templates() {
		fmt.Printline("imported template %v", t.Name())
		myT := t
		newRenderer := func(w *fmt.Writer) Renderer {
			return &TemplateRenderer{w, myT}
		}
		RegisterRendererFactory(myT.Name(), newRenderer)
	}
}

type TemplateRendererFactory struct {
	t *template.Template
	rendererFactory
}

type TemplateRenderer struct {
	*fmt.Writer
	*template.Template
}

func (r *TemplateRenderer) Id() string {
	return r.Name()
}

func (r *TemplateRenderer) Imports() []*Import {
	return nil
}

func (r *TemplateRenderer) RenderType(t *Type) {
	fmt.Printline("rendering %v", t.Name)
	err := r.Execute(r.Writer, t)
	PanicIf(err)
}
