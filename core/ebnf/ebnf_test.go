package ebnf

import (
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
	def := (&Ebnf{})
	input := "a"
	output, ok := def.chr("a", "z").Parse(input)
	if ok {
		t.Log("OK:", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}

func TestTerminalseq(t *T) {
	def := (&Ebnf{})
	input := "a1"
	output, ok := def.seq(def.chr("a", "z"), def.chr("0", "9")).Parse(input)
	if ok {
		t.Log("OK:", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}

func TestTerminalseqWithrpt(t *T) {
	def := (&Ebnf{})
	input := "1"
	output, ok := def.seq(def.rpt(def.chr("a", "z")), def.chr("0", "9")).Parse(input)
	if ok {
		t.Log("OK:", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}

func TestTerminaltoken(t *T) {
	def := (&Ebnf{})
	input := "token"

	output, ok := def.tok("token").Parse(input)
	if ok {
		t.Log("OK ", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}

func TestTerminalrpt(t *T) {
	def := (&Ebnf{})
	input := "abc"
	output, ok := def.rpt(def.chr("a", "z")).Parse(input)
	if ok {
		t.Log("OK ", input[:len(input)-len(output)])
	} else {
		t.Error("failed", output)
	}
}
