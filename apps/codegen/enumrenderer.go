package main

import (
	"github.com/ReSc/fmt"
	. "github.com/ReSc/godel/core/reflect"
	"sort"
	"strings"
)

func init() {
	RegisterRendererFactory("enum", NewEnumRenderer)
}

func NewEnumRenderer(w *fmt.Writer) Renderer {
	return &EnumRenderer{w}
}

type EnumRenderer struct {
	*fmt.Writer
}

func (r *EnumRenderer) Id() string {
	return "enum"
}

func (r *EnumRenderer) Imports() []*Import {
	return nil
}

func (r *EnumRenderer) RenderType(t *Type) {
	r.RenderEnumTypeDeclaration(t)
	r.RenderEnumConstructor(t)
	r.RenderEnumMethods(t)
}

func (r *EnumRenderer) RenderEnumTypeDeclaration(t *Type) {
	r.Writeln("// enum%v is the enum member type of %v", t.Name, t.Name)
	r.Writeln("type enum%v %v", t.Name, t.Meta.Name)
}

func (r *EnumRenderer) RenderEnumConstructor(t *Type) {
	r.Writeln("// %v is an enumeration", t.Name)
	r.Writeln("var %v = struct {", t.Name)
	keys := t.Fields.Keys()
	sort.Strings(keys)
	for _, key := range keys {
		p := t.Fields[key]
		r.Writeln("%v enum%v", p.Name, t.Name)
	}
	r.Writeln("} {")
	for _, key := range keys {
		p := t.Fields[key]
		r.Writeln("%v: %v,", p.Name, p.Value)
	}
	r.Writeln("}")
}

func (r *EnumRenderer) RenderEnumMethods(t *Type) {
	funcName := "String"
	r.Writeln("// %v returns the name of the %v enum member", funcName, t.Name)
	r.Writeln("func (this enum%v) %v() string {", t.Name, funcName)
	r.Writeln("switch this {")
	keys := t.Fields.Keys()
	sort.Strings(keys)
	for _, key := range keys {
		p := t.Fields[key]
		r.Writeln("case %v.%v:", t.Name, p.Name)
		r.Writeln("return \"%v\"", p.Name)
	}
	r.Writeln("default:")
	r.Writeln("panic(\"Invalid enum member for %v\")", t.Name)
	r.Writeln("}")
	r.Writeln("}")
	r.Writeln("")

	funcName = "Parse"
	r.Writeln("// %v%v returns the value for the name", funcName, t.Name)
	r.Writeln("func %v%v(name string) (enum%v,bool) {", funcName, t.Name, t.Name)
	r.Writeln("switch name {")
	for _, key := range keys {
		p := t.Fields[key]
		r.Writeln("case \"%v\":", p.Name)
		r.Writeln("return %v.%v, true", t.Name, p.Name)
	}
	for _, key := range keys {
		p := t.Fields[key]
		r.Writeln("default:")
		r.Writeln("return %v.%v, false", t.Name, p.Name)
		break
	}
	r.Writeln("}")
	r.Writeln("}")
	r.Writeln("")

	funcName = "Parse"
	r.Writeln("// %v%vCI returns the value for the name (case insensitive)", funcName, t.Name)
	r.Writeln("func %v%vCI(name string) (enum%v,bool) {", funcName, t.Name, t.Name)
	r.Writeln("name=strings.ToLower(name)")
	r.Writeln("switch name {")
	for _, key := range keys {
		p := t.Fields[key]
		name := strings.ToLower(p.Name)
		r.Writeln("case \"%v\":", name)
		r.Writeln("return %v.%v, true", t.Name, p.Name)
	}
	for _, key := range keys {
		p := t.Fields[key]
		r.Writeln("default:")
		r.Writeln("return %v.%v, false", t.Name, p.Name)
		break
	}
	r.Writeln("}")
	r.Writeln("}")
	r.Writeln("")
}
