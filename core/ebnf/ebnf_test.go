package ebnf

import (
	"github.com/ReSc/fmt"
	. "testing"
)

/*
func TestGrammar(t *T) {
	def := (&Ebnf{})
	prod := def.Production().Syntax()
	rest, ok := def.name().Parse(prod)
	if ok {
		t.Log("OK ", prod[:len(prod)-len(rest)])
	} else {
		t.Error("failed to parse grammar ", rest)
	}

}
*/

func TestTerminalchr(t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	input := "a"
	output, ok := def.chr("a", "z").Parse(scope, input)
	if ok {
		t.Log("OK:", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}

func TestName1(t *T) {
	input := "aName"
	testname(input, t)
}
func TestName2(t *T) {
	input := "AName"
	testname(input, t)
}
func testname(input string, t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	output, ok := def.Name().Parse(scope, input)
	if ok {
		t.Log("OK:", scope)
	} else {
		t.Error("failed", output)
	}
}
func TestToken(t *T) {
	input := "\"token\""
	scope := NewScope()
	def := (&Ebnf{})
	output, ok := def.token().Parse(scope, input)
	if ok {
		t.Log("OK:", scope)
	} else {
		t.Error("failed", output)
	}
}

func TestSeqWithRpt(t *T) {
	input := "a"
	scope := NewScope()
	def := (&Ebnf{})
	output, ok := def.seq(def._tokenchar(), def.rpt(def._tokenchar())).Parse(scope, input)
	if ok {
		t.Log("OK:", scope)
	} else {
		t.Error("failed", output)
	}
}

func TestSeqWithquoteChar(t *T) {
	input := "\""
	scope := NewScope()
	def := (&Ebnf{})
	output, ok := def.seq(def.tok("\"")).Parse(scope, input)
	if ok {
		t.Log("OK:", scope)
	} else {
		t.Error("failed", output)
	}
}

func TestSeqWithTwoQuoteChar(t *T) {
	input := "\"\""
	scope := NewScope()
	def := (&Ebnf{})
	output, ok := def.seq(def.tok("\""), def.tok("\"")).Parse(scope, input)
	if ok {
		t.Log("OK:", scope)
	} else {
		t.Error("failed", output)
	}
}
func TestSeqRptChar(t *T) {
	input := "a\""
	r, _ := firstRuneOf("\"")
	fmt.Printline("%U", r)
	scope := NewScope()
	def := (&Ebnf{})
	output, ok := def.seq(def.rpt(def._tokenchar()), def.tok("\"")).Parse(scope, input)
	if ok {
		t.Log("OK:", scope)
	} else {
		t.Error("failed", output)
	}
}

func TestTerminalchar(t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	input := "a"
	output, ok := def._tokenchar().Parse(scope, input)
	if ok {
		t.Log("OK:", scope)
	} else {
		t.Error("failed", output)
	}
}

func TestTerminalseq(t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	input := "a1"
	output, ok := def.seq(
		def.chr("a", "z"),
		def.chr("0", "9")).Parse(scope, input)
	if ok {
		t.Log("OK:", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}

func TestTerminalseqWithrpt(t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	input := "1"
	output, ok := def.seq(
		def.rpt(def.chr("a", "z")),
		def.chr("0", "9")).Parse(scope, input)
	if ok {
		t.Log("OK:", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}

func TestProduction(t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	input := "token=aname."

	output, ok := def.Production().Parse(scope, input)
	if ok {
		t.Log("OK ", scope)
	} else {
		t.Error("failed", output)
	}
}
func TestTerminaltok(t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	input := "token"

	output, ok := def.tok("token").Parse(scope, input)
	if ok {
		t.Log("OK ", scope)
	} else {
		t.Error("failed", output)
	}
}

func TestTerminalrpt(t *T) {
	scope := NewScope()
	def := (&Ebnf{})
	input := "abc"
	output, ok := def.rpt(def.chr("a", "z")).Parse(scope, input)
	if ok {
		t.Log("OK ", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}
