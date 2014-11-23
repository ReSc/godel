package ebnf

import (
	"github.com/ReSc/fmt"
	"strings"
	"unicode/utf8"
)

func (this *parser) LookAhead(r rune) bool {
	return this.Parser.LookAhead(r)
}

func (this *parser) Parse(input string) (string, bool) {
	output, ok := this.Parser.Parse(input)
	if ok {
		l := len(input) - len(output)
		if l > 0 {
			fmt.Printline("[%d] %s => CONSUMED\n%q\\n", this.Kind, this.Name, input[:l])
		} else {
			fmt.Printline("[%d] %s => NO COMSUME", this.Kind, this.Name)
		}
	} else {
		fmt.Printline("[%d] %s => ERROR", this.Kind, this.Name)
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

func (this *alt) LookAhead(r rune) bool {
	for _, inner := range this.inner {
		if inner.LookAhead(r) {
			return true
		}
	}
	return false
}

func (this *alt) Parse(input string) (string, bool) {
	r, l := firstRuneOf(input)
	if l == 0 {
		return input, false
	}

	output, ok := input, false
	for _, inner := range this.inner {
		if !inner.LookAhead(r) {
			continue
		}
		if s, sok := inner.Parse(input); !sok {
			continue
		} else {
			ok = true
			if len(output) > len(s) {
				output = s
			}
		}
	}
	return output, ok
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

func (this *seq) LookAhead(r rune) bool {
	return len(this.inner) > 0 && this.inner[0].LookAhead(r)
}
func (this *seq) Parse(input string) (string, bool) {
	output, ok := input, false
	for _, inner := range this.inner {
		r, l := firstRuneOf(output)
		if l == 0 || !inner.LookAhead(r) {
			return input, false
		}
		if output, ok = inner.Parse(output); !ok {
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

func (this *rpt) LookAhead(r rune) bool {
	return true
}

func (this *rpt) Parse(input string) (string, bool) {
	output := input
	for {
		r, l := firstRuneOf(output)
		if l > 0 && this.inner.LookAhead(r) {
			s, ok := this.inner.Parse(output)
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

func (this *opt) LookAhead(r rune) bool {
	return true
}

func (this *opt) Parse(input string) (string, bool) {
	output, _ := this.inner.Parse(input)
	return output, true
}
func (this *opt) String() string {
	return "[ " + this.inner.String() + " ]"
}

func (g *Ebnf) grp(p Parser) Parser { return &grp{p} }

type grp struct{ inner Parser }

func (this *grp) LookAhead(r rune) bool {
	return this.inner.LookAhead(r)
}
func (this *grp) Parse(input string) (string, bool) {
	return this.inner.Parse(input)
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

func (this *chr) LookAhead(r rune) bool {
	return this.startrune <= r && r <= this.endrune
}

func (this *chr) Parse(input string) (string, bool) {
	r, l := firstRuneOf(input)
	if l > 0 && this.LookAhead(r) {
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
	firstRune, l := firstRuneOf(token)
	if l == 0 {
		panic("cannot have a zero length token")
	}
	return &tok{
		token:     token,
		firstRune: firstRune,
	}
}

type tok struct {
	token     string
	firstRune rune
}

func (this *tok) LookAhead(r rune) bool {
	return this.firstRune == r
}
func (this *tok) Parse(input string) (string, bool) {
	output := strings.TrimPrefix(input, this.token)
	return output, len(output) < len(input)
}

func (this *tok) String() string {
	return fmt.String("%+q", this.token)
}
