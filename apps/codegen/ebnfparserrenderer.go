package main

import (
	"github.com/ReSc/fmt"
	. "github.com/ReSc/godel/core/reflect"
	"github.com/cznic/ebnf"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	_ "strings"
)

type ByProduction []string

func (p ByProduction) Len() int { return len(p) }
func (p ByProduction) Less(i, j int) bool {
	iT := isTerminal(p[i])
	jT := isTerminal(p[j])
	if iT != jT {
		return jT
	}
	return p[i] < p[j]
}

var terminalPattern = regexp.MustCompile("^[A-Z0-9_]+$")

func isTerminal(s string) bool {
	return terminalPattern.MatchString(s)
}

func (p ByProduction) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func init() {
	RegisterRendererFactory("ebnf", NewEbnfRenderer)
}

type EbnfRenderer struct {
	*fmt.Writer
	parserType *Type
}

func NewEbnfRenderer(w *fmt.Writer) Renderer {
	return &EbnfRenderer{w, nil}
}

func (r *EbnfRenderer) Id() string {
	return "ebnf"
}

func (r *EbnfRenderer) Imports() []*Import {
	return nil
}

func (r *EbnfRenderer) RenderType(t *Type) {
	dir := t.Package.GetPackageDirectory()
	grammarFile := filepath.Join(dir, t.Package.Name+".ebnf")
	if _, err := os.Stat(grammarFile); os.IsNotExist(err) {
		fmt.Printline("no such file or directory: %s", grammarFile)
		return
	}

	if reader, err := os.Open(grammarFile); err != nil {
		fmt.Printline("error opening %s: %s", grammarFile, err.Error())
		return
	} else {
		defer reader.Close()
		g, err := ebnf.Parse(grammarFile, reader)
		if err != nil {
			fmt.Printline("%s", err.Error())
			return
		}
		err = ebnf.Verify(g, "Grammar")
		if err != nil {
			fmt.Printline("%s", err.Error())
			return
		}

		productions := make([]string, 0, len(g))
		for k := range g {
			productions = append(productions, k)
		}
		sort.Stable(ByProduction(productions))

		r.parserType = t
		r.Writeln("type %s struct {", r.parserType.Name)
		for _, name := range productions {
			r.Writeln("my%s *%s", name, name)
		}
		r.Writeln("}\n")

		r.Writeln("func (this *%s) Syntax() string {\n", r.parserType.Name)
		r.Writeln("output:=\"\"")
		for _, name := range productions {
			r.Writeln("output +=  this.%s().Syntax() + \"\\n\"", name)
		}
		r.Writeln("return output\n")
		r.Writeln("\n}\n")

		r.Writeln("const (")
		r.Writeln("%s%s %s = iota", t.Meta.KeyName, productions[0], t.Meta.KeyType)
		for _, name := range productions[1:] {
			r.Writeln("%s%s", t.Meta.KeyName, name)
		}
		r.Writeln(")")

		for _, name := range productions {
			r.render(g[name])
		}

	}
}

func (r *EbnfRenderer) render(e ebnf.Expression) {
	switch prod := e.(type) {
	case ebnf.Alternative:
		last := len(prod) - 1
		r.WriteString("this.alt(\n")
		for i := range prod {
			r.render(prod[i])
			if i < last {
				r.Writefmt(",\n")
			}
		}
		r.WriteString(" )")
	case *ebnf.Bad:
		r.Writeln("// Parse error %s at %v", prod.Error, prod.Pos())
		break
	case *ebnf.Group:
		r.Writefmt("this.grp(\n")
		r.render(prod.Body)
		r.Writefmt(" )")
	case *ebnf.Name:
		r.Writefmt("this.")
		r.Writefmt(prod.String)
		r.Writefmt("()")
	case *ebnf.Option:
		r.Writefmt("this.opt(\n")
		r.render(prod.Body)
		r.Writefmt(" )")
	case *ebnf.Production:
		fmt.Printline("Rendering production\t%s", prod.Name.String)
		r.Writefmt("type %s struct {\n%s}\n", prod.Name.String, r.parserType.Meta.ElementTypeName)

		r.Writefmt("func (this *" + prod.Name.String + ") Parse(input string) (string,bool) {\n")
		r.Writefmt(" return this.Parser.Parse(input)\n")
		r.Writefmt("}\n")

		r.Writefmt("func (this *" + prod.Name.String + ") String() string {\n")
		r.Writefmt(" return this.Name\n")
		r.Writefmt("}\n")

		r.Writefmt("func (this *" + prod.Name.String + ") Syntax() string {\n")
		r.Writefmt(" return this.Name + \" = \" + this.Parser.String()\n")
		r.Writefmt("}\n")

		r.Writefmt("func (this *"+r.parserType.Name+") %[1]s() *%[1]s {\n", prod.Name.String)
		r.Writefmt("if this.my%s != nil {\n", prod.Name.String)
		r.Writefmt("return this.my%s\n", prod.Name.String)
		r.Writefmt("}\n")

		r.Writefmt("this.my%[1]s = &%[1]s{\n", prod.Name.String)
		r.Writefmt("%[1]s: %[1]s{\n", r.parserType.Meta.ElementTypeName)
		r.Writefmt("%s: %s%s,\n", r.parserType.Meta.KeyName, r.parserType.Meta.KeyName, prod.Name.String)
		r.Writefmt("Name: \"%s\",\n", prod.Name.String)
		r.Writeln("},\n}\n")

		r.Writefmt("this.my%s.Parser = ", prod.Name.String)
		r.render(prod.Expr)
		r.Writefmt("\nreturn this.my%s\n}\n", prod.Name.String)

	case *ebnf.Range:
		r.Write([]byte("this.chr("))
		r.WriteQuoted(prod.Begin.String)
		r.WriteString(", ")
		r.WriteQuoted(prod.End.String)
		r.Write([]byte(")"))
	case *ebnf.Repetition:
		r.Writefmt("this.rpt(\n")
		r.render(prod.Body)
		r.Writefmt(" )")
	case ebnf.Sequence:
		last := len(prod) - 1
		r.Writefmt("this.seq(\n")
		for i := range prod {
			r.render(prod[i])
			if i < last {
				r.Writefmt(",\n")
			}
		}
		r.Writefmt(" )")
	case *ebnf.Token:
		r.Writefmt("this.tok( ")
		r.WriteQuoted(prod.String)
		r.Writefmt(" )")
	default:
		r.Writefmt("this.def(%+v)", reflect.TypeOf(prod))
	}
}
