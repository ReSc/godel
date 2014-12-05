package ebnf

import ()

type Ebnf struct {
	myAlternation *Alternation
	myGrammar     *Grammar
	myGroup       *Group
	myName        *Name
	myOption      *Option
	myProduction  *Production
	myRange       *Range
	myRepetition  *Repetition
	mySequence    *Sequence
	myTerm        *Term
	myToken       *Token
	my_tokenchar  *_tokenchar
	myequals      *equals
	mypname       *pname
	mysingleToken *singleToken
	mytname       *tname
	mytoken       *token
}

func (this *Ebnf) Syntax() string {

	output := ""
	output += this.Alternation().Syntax() + "\n"
	output += this.Grammar().Syntax() + "\n"
	output += this.Group().Syntax() + "\n"
	output += this.Name().Syntax() + "\n"
	output += this.Option().Syntax() + "\n"
	output += this.Production().Syntax() + "\n"
	output += this.Range().Syntax() + "\n"
	output += this.Repetition().Syntax() + "\n"
	output += this.Sequence().Syntax() + "\n"
	output += this.Term().Syntax() + "\n"
	output += this.Token().Syntax() + "\n"
	output += this._tokenchar().Syntax() + "\n"
	output += this.equals().Syntax() + "\n"
	output += this.pname().Syntax() + "\n"
	output += this.singleToken().Syntax() + "\n"
	output += this.tname().Syntax() + "\n"
	output += this.token().Syntax() + "\n"
	return output

}

const (
	KindAlternation ProductionKind = iota
	KindGrammar
	KindGroup
	KindName
	KindOption
	KindProduction
	KindRange
	KindRepetition
	KindSequence
	KindTerm
	KindToken
	Kind_tokenchar
	Kindequals
	Kindpname
	KindsingleToken
	Kindtname
	Kindtoken
)

type Alternation struct {
	parser
}

func (this *Alternation) String() string {
	return this.Name
}
func (this *Alternation) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Alternation() *Alternation {
	if this.myAlternation != nil {
		return this.myAlternation
	}
	this.myAlternation = &Alternation{
		parser: parser{
			Kind:       KindAlternation,
			Name:       "Alternation",
			IsTerminal: false,
		},
	}

	this.myAlternation.Parser = this.seq(
		this.Sequence(),
		this.rpt(
			this.seq(
				this.tok("|"),
				this.Sequence())))
	return this.myAlternation
}

type Grammar struct {
	parser
}

func (this *Grammar) String() string {
	return this.Name
}
func (this *Grammar) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Grammar() *Grammar {
	if this.myGrammar != nil {
		return this.myGrammar
	}
	this.myGrammar = &Grammar{
		parser: parser{
			Kind:       KindGrammar,
			Name:       "Grammar",
			IsTerminal: false,
		},
	}

	this.myGrammar.Parser = this.seq(
		this.Production(),
		this.rpt(
			this.Production()))
	return this.myGrammar
}

type Group struct {
	parser
}

func (this *Group) String() string {
	return this.Name
}
func (this *Group) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Group() *Group {
	if this.myGroup != nil {
		return this.myGroup
	}
	this.myGroup = &Group{
		parser: parser{
			Kind:       KindGroup,
			Name:       "Group",
			IsTerminal: false,
		},
	}

	this.myGroup.Parser = this.seq(
		this.tok("("),
		this.Alternation(),
		this.tok(")"))
	return this.myGroup
}

type Name struct {
	parser
}

func (this *Name) String() string {
	return this.Name
}
func (this *Name) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Name() *Name {
	if this.myName != nil {
		return this.myName
	}
	this.myName = &Name{
		parser: parser{
			Kind:       KindName,
			Name:       "Name",
			IsTerminal: false,
		},
	}

	this.myName.Parser = this.alt(
		this.pname(),
		this.tname())
	return this.myName
}

type Option struct {
	parser
}

func (this *Option) String() string {
	return this.Name
}
func (this *Option) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Option() *Option {
	if this.myOption != nil {
		return this.myOption
	}
	this.myOption = &Option{
		parser: parser{
			Kind:       KindOption,
			Name:       "Option",
			IsTerminal: false,
		},
	}

	this.myOption.Parser = this.seq(
		this.tok("["),
		this.Alternation(),
		this.tok("]"))
	return this.myOption
}

type Production struct {
	parser
}

func (this *Production) String() string {
	return this.Name
}
func (this *Production) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Production() *Production {
	if this.myProduction != nil {
		return this.myProduction
	}
	this.myProduction = &Production{
		parser: parser{
			Kind:       KindProduction,
			Name:       "Production",
			IsTerminal: false,
		},
	}

	this.myProduction.Parser = this.seq(
		this.Name(),
		this.equals(),
		this.Alternation(),
		this.tok("."))
	return this.myProduction
}

type Range struct {
	parser
}

func (this *Range) String() string {
	return this.Name
}
func (this *Range) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Range() *Range {
	if this.myRange != nil {
		return this.myRange
	}
	this.myRange = &Range{
		parser: parser{
			Kind:       KindRange,
			Name:       "Range",
			IsTerminal: false,
		},
	}

	this.myRange.Parser = this.seq(
		this.singleToken(),
		this.tok("\u2026"),
		this.singleToken())
	return this.myRange
}

type Repetition struct {
	parser
}

func (this *Repetition) String() string {
	return this.Name
}
func (this *Repetition) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Repetition() *Repetition {
	if this.myRepetition != nil {
		return this.myRepetition
	}
	this.myRepetition = &Repetition{
		parser: parser{
			Kind:       KindRepetition,
			Name:       "Repetition",
			IsTerminal: false,
		},
	}

	this.myRepetition.Parser = this.seq(
		this.tok("{"),
		this.Alternation(),
		this.tok("}"))
	return this.myRepetition
}

type Sequence struct {
	parser
}

func (this *Sequence) String() string {
	return this.Name
}
func (this *Sequence) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Sequence() *Sequence {
	if this.mySequence != nil {
		return this.mySequence
	}
	this.mySequence = &Sequence{
		parser: parser{
			Kind:       KindSequence,
			Name:       "Sequence",
			IsTerminal: false,
		},
	}

	this.mySequence.Parser = this.seq(
		this.Term(),
		this.rpt(
			this.Term()))
	return this.mySequence
}

type Term struct {
	parser
}

func (this *Term) String() string {
	return this.Name
}
func (this *Term) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Term() *Term {
	if this.myTerm != nil {
		return this.myTerm
	}
	this.myTerm = &Term{
		parser: parser{
			Kind:       KindTerm,
			Name:       "Term",
			IsTerminal: false,
		},
	}

	this.myTerm.Parser = this.alt(
		this.Name(),
		this.Token(),
		this.Group(),
		this.Option(),
		this.Repetition())
	return this.myTerm
}

type Token struct {
	parser
}

func (this *Token) String() string {
	return this.Name
}
func (this *Token) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) Token() *Token {
	if this.myToken != nil {
		return this.myToken
	}
	this.myToken = &Token{
		parser: parser{
			Kind:       KindToken,
			Name:       "Token",
			IsTerminal: false,
		},
	}

	this.myToken.Parser = this.alt(
		this.Range(),
		this.token())
	return this.myToken
}

type _tokenchar struct {
	parser
}

func (this *_tokenchar) String() string {
	return this.Name
}
func (this *_tokenchar) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) _tokenchar() *_tokenchar {
	if this.my_tokenchar != nil {
		return this.my_tokenchar
	}
	this.my_tokenchar = &_tokenchar{
		parser: parser{
			Kind:       Kind_tokenchar,
			Name:       "_tokenchar",
			IsTerminal: false,
		},
	}

	this.my_tokenchar.Parser = this.alt(
		this.tok("\t"),
		this.tok("\n"),
		this.tok("\r"),
		this.tok(" "),
		this.tok("!"),
		this.tok("\\\""),
		this.chr("#", "\ud7ff"),
		this.chr("\ue000", "\ufffd"),
		this.chr("\U00010000", "\U0010ffff"))
	return this.my_tokenchar
}

type equals struct {
	parser
}

func (this *equals) String() string {
	return this.Name
}
func (this *equals) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) equals() *equals {
	if this.myequals != nil {
		return this.myequals
	}
	this.myequals = &equals{
		parser: parser{
			Kind:       Kindequals,
			Name:       "equals",
			IsTerminal: true,
		},
	}

	this.myequals.Parser = this.tok("=")
	return this.myequals
}

type pname struct {
	parser
}

func (this *pname) String() string {
	return this.Name
}
func (this *pname) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) pname() *pname {
	if this.mypname != nil {
		return this.mypname
	}
	this.mypname = &pname{
		parser: parser{
			Kind:       Kindpname,
			Name:       "pname",
			IsTerminal: true,
		},
	}

	this.mypname.Parser = this.seq(
		this.chr("A", "Z"),
		this.rpt(
			this.alt(
				this.chr("a", "z"),
				this.chr("A", "Z"),
				this.chr("0", "9"),
				this.tok("_"))))
	return this.mypname
}

type singleToken struct {
	parser
}

func (this *singleToken) String() string {
	return this.Name
}
func (this *singleToken) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) singleToken() *singleToken {
	if this.mysingleToken != nil {
		return this.mysingleToken
	}
	this.mysingleToken = &singleToken{
		parser: parser{
			Kind:       KindsingleToken,
			Name:       "singleToken",
			IsTerminal: true,
		},
	}

	this.mysingleToken.Parser = this.seq(
		this.tok("\""),
		this._tokenchar(),
		this.tok("\""))
	return this.mysingleToken
}

type tname struct {
	parser
}

func (this *tname) String() string {
	return this.Name
}
func (this *tname) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) tname() *tname {
	if this.mytname != nil {
		return this.mytname
	}
	this.mytname = &tname{
		parser: parser{
			Kind:       Kindtname,
			Name:       "tname",
			IsTerminal: true,
		},
	}

	this.mytname.Parser = this.seq(
		this.chr("a", "z"),
		this.rpt(
			this.alt(
				this.chr("a", "z"),
				this.chr("A", "Z"),
				this.chr("0", "9"),
				this.tok("_"))))
	return this.mytname
}

type token struct {
	parser
}

func (this *token) String() string {
	return this.Name
}
func (this *token) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) token() *token {
	if this.mytoken != nil {
		return this.mytoken
	}
	this.mytoken = &token{
		parser: parser{
			Kind:       Kindtoken,
			Name:       "token",
			IsTerminal: true,
		},
	}

	this.mytoken.Parser = this.seq(
		this.tok("\""),
		this._tokenchar(),
		this.rpt(
			this._tokenchar()),
		this.tok("\""))
	return this.mytoken
}

type Parser interface {
	LookAhead(input string) (ok bool)
	Parse(scope Scope, input string) (unparsed string, ok bool)
	String() (s string)
}
type ProductionKind int

type Scope interface {
	Len() (l int)
	Mark() (mark int)
	Peek() (kind ProductionKind, name string, value string)
	Pop() (kind ProductionKind, name string, value string)
	PopUntil(mark int) (kind []ProductionKind, name []string, value []string)
	Push(kind ProductionKind, name string, value string)
	PushAll(kind []ProductionKind, name []string, value []string) (self Scope)
}

// parser is a struct
type parser struct {
	IsTerminal bool
	Kind       ProductionKind
	Name       string
	Parser     Parser
}

// Newparser creates a new instance of parser
func Newparser() *parser {
	return &parser{}
}
