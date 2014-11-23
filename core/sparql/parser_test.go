package sparql

import (
	"github.com/ReSc/fmt"
	. "testing"
)

func TestSyntax(t *T) {
	t.Error(New.Syntax())
}

func TestANONGood(t *T) {
	goodInputs := []string{"[]", "[\t]", "[\r\n]"}
	testParser(goodInputs, New.ANON(), t)
}

func TestNumberGood(t *T) {
	goodInputs := []string{"1", "0", "10a"}
	testParser(goodInputs, New.INTEGER(), t)
}

func TestNumberError(t *T) {
	badInputs := []string{"aa1", "@0", "h10g"}
	antiTestParser(badInputs, New.INTEGER(), t)
}

func testParser(inputs []string, parser SyntaxParser, t *T) {
	for _, input := range inputs {
		leftover, ok := parser.Parse(input)
		if !ok || len(leftover) != 0 {

			fmt.Printline("error parsing `%v` with production `%s` leftover: `%s`", input, parser.Syntax(), leftover)
		}
	}
}

func antiTestParser(inputs []string, parser SyntaxParser, t *T) {
	for _, input := range inputs {
		leftover, ok := parser.Parse(input)
		if ok {
			fmt.Printline("Expected an error parsing `%v` with production %s leftover: `%s`", input, parser.Syntax(), leftover)
		}
	}
}
