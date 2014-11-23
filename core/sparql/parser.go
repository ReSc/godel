package sparql

import (
	"github.com/ReSc/fmt"
	"unicode/utf8"
)

var (
	New Sparql
)

type (
	Syntaxer interface {
		Syntax() string
	}
	SyntaxParser interface {
		Syntaxer
		Parser
	}
)

func (p *Production) LookAhead(r rune) bool {
	return p.Parser.LookAhead(r)
}

func (p *Production) Parse(s string) (string, bool) {
	return p.Parser.Parse(s)
}

func (p *Sparql) seq(pr ...Parser) Parser {
	return &seq{
		inner: pr,
	}
}

type seq struct {
	inner []Parser
}

func (p *seq) LookAhead(r rune) bool {
	return len(p.inner) > 0 && p.inner[0].LookAhead(r)
}

func (p *seq) Parse(s string) (string, bool) {
	input, ok := s, true
	for _, inner := range p.inner {
		if input, ok = inner.Parse(input); !ok {
			return s, false
		}
	}
	return input, true
}

func (p *seq) String() string {
	output := ""
	for i, inner := range p.inner {
		output += inner.String()
		if i < len(p.inner)-1 {
			output += " "
		}
	}
	return output
}

func (p *Sparql) rpt(pr Parser) Parser {
	return &rpt{
		inner: pr,
	}
}

type rpt struct {
	inner Parser
}

func (p *rpt) LookAhead(r rune) bool {
	return p.inner.LookAhead(r)
}

func (p *rpt) Parse(s string) (string, bool) {
	lastInput, input, ok := s, s, true
	for {
		input, ok = p.inner.Parse(input)
		progress := len(lastInput) > len(input)
		if ok && progress {
			lastInput = input
			continue
		}
		break
	}
	if len(lastInput) < len(s) {
		return lastInput, true
	}
	return s, false
}

func (p *rpt) String() string {
	output := "{ " + p.inner.String() + " }"
	return output
}

func (p *Sparql) grp(pr Parser) Parser {
	return &grp{
		inner: pr,
	}
}

type grp struct {
	inner Parser
}

func (p *grp) LookAhead(r rune) bool {
	return p.inner.LookAhead(r)
}

func (p *grp) Parse(s string) (string, bool) {
	return p.Parse(s)
}

func (p *grp) String() string {
	output := "( " + p.inner.String() + " )"
	return output
}

func (p *Sparql) alt(pr ...Parser) Parser {
	return &alt{
		inner: pr,
	}
}

type alt struct {
	inner []Parser
}

func (p *alt) LookAhead(r rune) bool {
	for _, inner := range p.inner {
		if inner.LookAhead(r) {
			return true
		}
	}
	return false
}

func (p *alt) Parse(s string) (string, bool) {
	lastInput := s
	for _, inner := range p.inner {
		if input, ok := inner.Parse(s); ok && len(input) < len(lastInput) {
			lastInput = input
		}
	}
	return lastInput, len(lastInput) < len(s)
}

func (p *alt) String() string {
	output := " "
	sep := " | "
	if len(p.inner) > 5 {
		sep = "\n\t| "
	}
	for i, inner := range p.inner {
		output += inner.String()
		if i < len(p.inner)-1 {
			output += sep
		}
	}
	return output + " "
}

func (p *Sparql) opt(pr Parser) Parser {
	return &opt{
		inner: pr,
	}
}

type opt struct{ inner Parser }

func (p *opt) LookAhead(r rune) bool {
	return p.LookAhead(r)
}

func (p *opt) Parse(s string) (string, bool) {
	if input, ok := p.inner.Parse(s); ok {
		return input, ok
	}
	return s, true
}

func (p *opt) String() string {
	output := "[ " + p.inner.String() + " ]"
	return output
}

func (p *Sparql) tok(token string) Parser {
	count := utf8.RuneCountInString(token)
	runes := make([]rune, 0, count)
	s := token
	for r, l := firstRune(s); l > 0; r, l = firstRune(s) {
		s = s[l:]
		runes = append(runes, r)
	}
	return &tok{
		token:  token,
		runes:  runes,
		length: len(runes),
	}
}

type tok struct {
	token  string
	runes  []rune
	length int
}

func (p *tok) LookAhead(r rune) bool {
	return p.length > 0 && p.runes[0] == r
}

func (p *tok) Parse(s string) (string, bool) {
	input := s
	index := 0
	for r, l := firstRune(input); l > 0 && index < p.length; r, l = firstRune(input) {
		input = input[l:]
		if r != p.runes[index] {
			return s, false
		}
		index++
		if index == p.length {
			return input, true
		}
	}
	return s, false
}

func (p *tok) String() string {
	return fmt.String("%q", p.token)
}

func (p *Sparql) chr(start, end string) Parser {
	parser := &chr{Start: start, End: end}
	r, l := firstRune(start)
	if l != len(start) {
		panic("range start should be a single rune")
	} else {
		parser.start = r
	}
	r, l = firstRune(end)
	if l != len(end) {
		panic("range end should be a single rune")
	} else {
		parser.end = r
	}
	return parser
}

type chr struct {
	Start string
	End   string
	start rune
	end   rune
}

func (p *chr) LookAhead(r rune) bool {
	return p.start <= r && r <= p.end
}

func (p *chr) Parse(s string) (string, bool) {
	r, l := firstRune(s)
	if l > 0 && p.LookAhead(r) {
		return s[l:], true
	}
	return s, false
}

func (p *chr) String() string {
	return fmt.String("%+q … %+q", p.Start, p.End)
}

func firstRune(s string) (rune, int) {
	r, length := utf8.DecodeRuneInString(s)
	return r, length
}
