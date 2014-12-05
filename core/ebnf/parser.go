package ebnf

import (
	"github.com/ReSc/fmt"
	"strings"
	_ "unicode"
	"unicode/utf8"
)

type Emitter interface {
	Emit(value interface{})
}

func NewScope() Scope {
	return &scope{}
}

type scope struct {
	kinds  []ProductionKind
	names  []string
	values []string
}

func (this *scope) Push(kind ProductionKind, name string, value string) {
	fmt.Printline("[%d] %s >> %s", this.Len(), name, value)
	this.kinds = append(this.kinds, kind)
	this.names = append(this.names, name)
	this.values = append(this.values, value)
}

func (this *scope) PushAll(kind []ProductionKind, name []string, value []string) Scope {
	if len(kind) != len(name) || len(name) != len(value) {
		panic("unbalanced push")
	}
	this.kinds = append(this.kinds, kind...)
	this.names = append(this.names, name...)
	this.values = append(this.values, value...)
	return this
}

func (this *scope) Len() int {
	return len(this.kinds)
}

func (this *scope) Peek() (kind ProductionKind, name string, value string) {
	top := this.Len() - 1
	if top == 0 {
		return this.kinds[top], this.names[top], this.values[top]
	}
	return 0, "", ""
}

func (this *scope) Pop() (kind ProductionKind, name string, value string) {
	kind, name, value = this.Peek()
	top := this.Len() - 1
	if top >= 0 {
		this.kinds = this.kinds[:top]
		this.names = this.names[:top]
		this.values = this.values[:top]
	}
	return kind, name, value
}

// PopUntil pops the scope stack from the mark to the top, i.e. returns an slice with mark at index 0.
// the returnvalues are valid until the next Push()
func (this *scope) PopUntil(mark int) (kind []ProductionKind, name []string, value []string) {
	l := mark - this.Mark()
	if l < 0 || mark >= this.Len() {
		return nil, nil, nil
	}

	kind = this.kinds[mark:]
	name = this.names[mark:]
	value = this.values[mark:]

	this.kinds = this.kinds[:mark]
	this.names = this.names[:mark]
	this.values = this.values[:mark]

	return kind, name, value
}

func (this *scope) Mark() int {
	return this.Len()
}

func (this *scope) String() string {
	output := ""
	for i := range this.kinds {
		output += fmt.String("[%d] %s: %s\n", this.kinds[i], this.names[i], this.values[i])
	}
	if output == "" {
		return "<EMPTY>"
	}
	return output
}

func (this *parser) LookAhead(input string) bool {
	return this.Parser.LookAhead(input)
}

func (this *parser) Parse(scope Scope, input string) (string, bool) {
	mark := scope.Mark()
	output, ok := this.Parser.Parse(scope, input)
	if ok {
		l := len(input) - len(output)
		if l > 0 {
			if this.IsTerminal {
				scope.Push(this.Kind, this.Name, input[:l])
			} else {

			}
		} else {
			fmt.Printline("[%d] %s => NO CONSUME", this.Kind, this.Name)
		}
	} else {
		scope.PopUntil(mark)
	}
	return output, ok
}

func (this *parser) String() string {
	return this.Name
}

func (this *parser) Grammar() string {
	return this.Name + " = " + this.Parser.String() + " ."
}

func (g *Ebnf) alt(p ...Parser) Parser { return &alt{p} }

type alt struct{ inner []Parser }

func (this *alt) LookAhead(input string) bool {
	for _, inner := range this.inner {
		if inner.LookAhead(input) {
			return true
		}
	}
	return false
}

func (this *alt) Parse(scope Scope, input string) (string, bool) {
	var result Scope
	output, ok := input, false
	for _, inner := range this.inner {
		if !inner.LookAhead(input) {
			continue
		}
		altScope := NewScope()
		if s, sok := inner.Parse(altScope, input); !sok {
			continue
		} else {
			if len(output) > len(s) {
				ok = true
				output = s
				result = altScope
			}
		}
	}
	if ok {
		if result != nil {
			scope.PushAll(result.PopUntil(0))
		} else {
			fmt.Printline("noresult!")
		}
		return output, ok
	}
	return input, false
}
func (this *alt) String() string {
	if len(this.inner) == 0 {
		return ""
	}
	s := this.inner[0].String()

	for _, inner := range this.inner[1:] {
		s = s + " | " + inner.String()
	}
	return s
}

func (g *Ebnf) seq(p ...Parser) Parser { return &seq{p} }

type seq struct{ inner []Parser }

func (this *seq) LookAhead(input string) bool {
	return len(this.inner) > 0 && this.inner[0].LookAhead(input)
}
func (this *seq) Parse(scope Scope, input string) (string, bool) {
	output, ok := input, false
	for _, inner := range this.inner {
		if !inner.LookAhead(output) {
			return input, false
		}
		if output, ok = inner.Parse(scope, output); !ok {
			return input, false
		}
	}
	return output, ok
}
func (this *seq) String() string {
	if len(this.inner) == 0 {
		return ""
	}

	s := this.inner[0].String()
	for _, inner := range this.inner[1:] {
		s = s + " " + inner.String()
	}

	return s
}

func (g *Ebnf) rpt(p Parser) Parser { return &rpt{p} }

type rpt struct{ inner Parser }

func (this *rpt) LookAhead(r string) bool {
	return true
}

func (this *rpt) Parse(scope Scope, input string) (string, bool) {
	output := input
	for {
		if this.inner.LookAhead(output) {
			s, ok := this.inner.Parse(scope, output)
			if ok && len(s) < len(output) {
				output = s
				continue
			}
		}
		break
	}
	return output, true
}

func (this *rpt) String() string {
	return "{ " + this.inner.String() + " }"
}

func (g *Ebnf) opt(p Parser) Parser { return &opt{p} }

type opt struct{ inner Parser }

func (this *opt) LookAhead(r string) bool {
	return true
}

func (this *opt) Parse(scope Scope, input string) (string, bool) {
	output, _ := this.inner.Parse(scope, input)
	return output, true
}
func (this *opt) String() string {
	return "[ " + this.inner.String() + " ]"
}

func (g *Ebnf) grp(p Parser) Parser { return &grp{p} }

type grp struct{ inner Parser }

func (this *grp) LookAhead(r string) bool {
	return this.inner.LookAhead(r)
}
func (this *grp) Parse(scope Scope, input string) (string, bool) {
	return this.inner.Parse(scope, input)
}
func (this *grp) String() string {
	return "( " + this.inner.String() + " )"
}

func (g *Ebnf) chr(start, end string) Parser {
	startrune, l := firstRuneOf(start)
	if l != len(start) {
		panic("start is not a single rune")
	}
	endrune, l := firstRuneOf(end)
	if l != len(end) {
		panic("end is not a single rune")
	}
	return &chr{
		start: start,
		end:   end,

		startrune: startrune,
		endrune:   endrune,
	}
}

type chr struct {
	start, end         string
	startrune, endrune rune
}

func (this *chr) LookAhead(input string) bool {
	r, l := firstRuneOf(input)
	if l == 0 {
		return false
	}
	return this.startrune <= r && r <= this.endrune
}

func (this *chr) Parse(scope Scope, input string) (string, bool) {
	r, l := firstRuneOf(input)
	if l > 0 && this.startrune <= r && r <= this.endrune {
		return input[l:], true
	}
	return input, false
}

func firstRuneOf(s string) (r rune, length int) {
	if !utf8.FullRuneInString(s) {
		return 0, 0
	}
	return utf8.DecodeRuneInString(s)
}

func (this *chr) String() string {
	return fmt.String("%+q … %+q", this.start, this.end)
}

func (g *Ebnf) tok(token string) Parser {
	_, l := firstRuneOf(token)
	if l == 0 {
		panic("cannot have a zero length token")
	}
	return &tok{
		token: token,
	}
}

type tok struct {
	token string
}

func (this *tok) LookAhead(input string) bool {
	return strings.HasPrefix(input, this.token)
}
func (this *tok) Parse(scope Scope, input string) (string, bool) {
	output := strings.TrimPrefix(input, this.token)
	return output, len(output) < len(input)
}

func (this *tok) String() string {
	return fmt.String("%+q", this.token)
}
