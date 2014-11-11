package main

import (
	"github.com/ReSc/fmt"
	. "github.com/ReSc/godel/core/reflect"
	_ "sort"
	_ "strings"
)

func init() {
	RegisterRendererFactory("list", NewListRenderer)
}

type ListRenderer struct {
	*fmt.Writer
}

func NewListRenderer(w *fmt.Writer) Renderer {
	return &ListRenderer{w}
}

func (r *ListRenderer) Id() string {
	return "list"
}

func (r *ListRenderer) Imports() []*Import {
	return nil
}

func (r *ListRenderer) RenderType(t *Type) {
	r.RenderListTypeDeclaration(t)
	r.RenderListConstructor(t)
	r.RenderListMethods(t)
}

func (r *ListRenderer) RenderListTypeDeclaration(t *Type) {
	r.Writeln("// %v is a %v", t.Name, t.Meta.Name)
	r.Writeln("type %v %v", t.Name, t.Meta.Name)
	r.Writeln("")
}

func (r *ListRenderer) RenderListConstructor(t *Type) {
	r.Writeln("// New%[1]v creates a new instance of %[1]v", t.Name)
	r.Writefmt("func New%v", t.Name)
	r.Writeln("() %v {", t.Name)
	r.Writeln("return make(%v, 0, 4)", t.Name)
	r.Writeln("}")
	r.Writeln("")
	r.Writeln("// New%[1]vSized creates a new instance of %[1]v", t.Name)
	r.Writefmt("func New%vSized", t.Name)
	r.Writeln("(size,capacity int) %v {", t.Name)
	r.Writeln("return make(%v, size, capacity)", t.Name)
	r.Writeln("}")
	r.Writeln("")
}

func (r *ListRenderer) RenderListMethods(t *Type) {
	r.Writeln("func (this *%v) Len() int{", t.Name)
	r.Writeln("return len(*this)")
	r.Writeln("}")
	r.Writeln("")

	r.Writeln("func (this *%v) Contains(value %v) bool {", t.Name, t.Meta.ElementType)
	r.Writeln("return this.IndexOf(value) >= 0")
	r.Writeln("}")
	r.Writeln("")

	r.Writeln("func (this *%v) Del(value %v) bool {", t.Name, t.Meta.ElementType)
	r.Writeln("if i:= this.IndexOf(value); i >= 0 {")
	r.Writeln("l:=*this")
	r.Writeln("copy(l[i:],l[i+1:])")
	r.Writeln("var defaultValue %v", t.Meta.ElementType)
	r.Writeln("l[len(l)-1]=defaultValue")
	r.Writeln("l=l[:len(l)-1]")
	r.Writeln("*this=l")
	r.Writeln("return true")
	r.Writeln("}")
	r.Writeln("return false")
	r.Writeln("}")
	r.Writeln("")

	r.Writeln("func (this *%v) Add(value %v) bool {", t.Name, t.Meta.ElementType)
	r.Writeln("l:=*this")
	r.Writeln("l=append(l,value)")
	r.Writeln("*this=l")
	r.Writeln("return true")
	r.Writeln("}")
	r.Writeln("")

	r.Writeln("func (this *%v) IndexOf(value %v) int {", t.Name, t.Meta.ElementType)
	r.Writeln("items := *this ")
	r.Writeln("for index, item := range items {")
	r.Writeln("if item == value {")
	r.Writeln("return index")
	r.Writeln("}")
	r.Writeln("}")
	r.Writeln("return -1")
	r.Writeln("}")
	r.Writeln("")

	r.Writeln("// Each iterates over all items")
	r.Writeln("func (this *%v) Each(f func(int,%v)) {", t.Name, t.Meta.ElementType)
	r.Writeln("items := *this ")
	r.Writeln("for index, item := range items {")
	r.Writeln("f(index,item)")
	r.Writeln("}")
	r.Writeln("}")
	r.Writeln("")
}
