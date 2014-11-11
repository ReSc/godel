package main

import (
	"github.com/ReSc/fmt"
	. "github.com/ReSc/godel/core/reflect"
	_ "sort"
	_ "strings"
)

func init() {
	RegisterRendererFactory("set", NewSetRenderer)
}

type SetRenderer struct {
	*fmt.Writer
}

func NewSetRenderer(w *fmt.Writer) Renderer {
	return &SetRenderer{w}
}

func (r *SetRenderer) Id() string {
	return "set"
}

func (r *SetRenderer) Imports() []*Import {
	return nil
}

func (r *SetRenderer) RenderType(t *Type) {
	r.RenderSetTypeDeclaration(t)
	r.RenderSetConstructor(t)
	r.RenderSetMethods(t)
}

func (r *SetRenderer) RenderSetTypeDeclaration(t *Type) {
	r.Writeln("// %v is a %v", t.Name, t.Meta.Name)
	r.Writeln("type %v %v", t.Name, t.Meta.Name)
	r.Writeln("")
}

func (r *SetRenderer) RenderSetConstructor(t *Type) {
	r.Writeln("// New%[1]v creates a new instance of %[1]v", t.Name)
	r.Writefmt("func New%v", t.Name)
	r.Writeln("() %v {", t.Name)
	r.Writeln("return make(%v)", t.Name)
	r.Writeln("}")
	r.Writeln("")
}

func (r *SetRenderer) RenderSetMethods(t *Type) {
	r.Writeln("func (this %v) Contains(key %v) bool {", t.Name, t.Meta.KeyType)
	r.Writeln("_, ok := this[key]")
	r.Writeln("return ok")
	r.Writeln("}")
	r.Writeln("")
	r.Writeln("func (this %v) Len() int {", t.Name)
	r.Writeln("return len(this)")
	r.Writeln("}")
	r.Writeln("")
	r.Writeln("func (this %v) Add(key %v) bool {", t.Name, t.Meta.KeyType)
	r.Writeln("if !this.Contains(key) {")
	r.Writeln("this[key]=true")
	r.Writeln("return true")
	r.Writeln("}")
	r.Writeln("return false")
	r.Writeln("}")
	r.Writeln("")
	r.Writeln("func (this %v) Del(key %v) bool {", t.Name, t.Meta.KeyType)
	r.Writeln("if this.Contains(key) {")
	r.Writeln("delete(this,key)")
	r.Writeln("return true")
	r.Writeln("}")
	r.Writeln("return false")
	r.Writeln("}")

	r.Writeln("// Each iterates over all items")
	r.Writeln("func (this %v) Each(f func(%v)) {", t.Name, t.Meta.KeyType)
	r.Writeln("for key, _ := range this {")
	r.Writeln("f(key)")
	r.Writeln("}")
	r.Writeln("}")
	r.Writeln("")
}
