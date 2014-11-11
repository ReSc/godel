package main

import (
	"github.com/ReSc/fmt"
	. "github.com/ReSc/godel/core/reflect"
	"github.com/cznic/ebnf"
	"os"
	"path/filepath"
	_ "sort"
	_ "strings"
)

func init() {
	RegisterRendererFactory("ebnf", NewEbnfRenderer)
}

type EbnfRenderer struct {
	*fmt.Writer
}

func NewEbnfRenderer(w *fmt.Writer) Renderer {
	return &EbnfRenderer{w}
}

func (r *EbnfRenderer) Id() string {
	return "ebnf"
}

func (r *EbnfRenderer) Imports() []*Import {
	return nil
}

func (r *EbnfRenderer) RenderType(t *Type) {
	dir := t.Package.GetPackageDirectory()
	grammarFile := filepath.Join(dir, t.Name+".enbf")
	if _, err := os.Stat(grammarFile); os.IsNotExist(err) {
		fmt.Printline("no such file or directory: %s", filename)
		return
	}
}
