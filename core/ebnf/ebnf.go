package ebnf

import ()

type Ebnf struct {
	myAlternation *Alternation
	myGrammar     *Grammar
	myGroup       *Group
	myOption      *Option
	myProduction  *Production
	myRange       *Range
	myRepetition  *Repetition
	mySequence    *Sequence
	myTerm        *Term
	myToken       *Token
	mychar        *char
	myname        *name
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
	output += this.Option().Syntax() + "\n"
	output += this.Production().Syntax() + "\n"
	output += this.Range().Syntax() + "\n"
	output += this.Repetition().Syntax() + "\n"
	output += this.Sequence().Syntax() + "\n"
	output += this.Term().Syntax() + "\n"
	output += this.Token().Syntax() + "\n"
	output += this.char().Syntax() + "\n"
	output += this.name().Syntax() + "\n"
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
	KindOption
	KindProduction
	KindRange
	KindRepetition
	KindSequence
	KindTerm
	KindToken
	Kindchar
	Kindname
	Kindpname
	KindsingleToken
	Kindtname
	Kindtoken
)

type Alternation struct {
	parser
}

func (this *Alternation) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindAlternation,
			Name: "Alternation",
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

func (this *Grammar) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindGrammar,
			Name: "Grammar",
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

func (this *Group) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindGroup,
			Name: "Group",
		},
	}

	this.myGroup.Parser = this.seq(
		this.tok("("),
		this.Alternation(),
		this.tok(")"))
	return this.myGroup
}

type Option struct {
	parser
}

func (this *Option) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindOption,
			Name: "Option",
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

func (this *Production) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindProduction,
			Name: "Production",
		},
	}

	this.myProduction.Parser = this.seq(
		this.name(),
		this.tok("="),
		this.Alternation(),
		this.tok("."))
	return this.myProduction
}

type Range struct {
	parser
}

func (this *Range) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindRange,
			Name: "Range",
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

func (this *Repetition) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindRepetition,
			Name: "Repetition",
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

func (this *Sequence) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindSequence,
			Name: "Sequence",
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

func (this *Term) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindTerm,
			Name: "Term",
		},
	}

	this.myTerm.Parser = this.alt(
		this.name(),
		this.Token(),
		this.Group(),
		this.Option(),
		this.Repetition())
	return this.myTerm
}

type Token struct {
	parser
}

func (this *Token) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindToken,
			Name: "Token",
		},
	}

	this.myToken.Parser = this.alt(
		this.Range(),
		this.token())
	return this.myToken
}

type char struct {
	parser
}

func (this *char) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *char) String() string {
	return this.Name
}
func (this *char) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) char() *char {
	if this.mychar != nil {
		return this.mychar
	}
	this.mychar = &char{
		parser: parser{
			Kind: Kindchar,
			Name: "char",
		},
	}

	this.mychar.Parser = this.alt(
		this.tok("\t"),
		this.tok("\n"),
		this.tok("\r"),
		this.chr(" ", "\ud7ff"),
		this.chr("\ue000", "\ufffd"),
		this.chr("\U00010000", "\U0010ffff"))
	return this.mychar
}

type name struct {
	parser
}

func (this *name) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *name) String() string {
	return this.Name
}
func (this *name) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Ebnf) name() *name {
	if this.myname != nil {
		return this.myname
	}
	this.myname = &name{
		parser: parser{
			Kind: Kindname,
			Name: "name",
		},
	}

	this.myname.Parser = this.alt(
		this.pname(),
		this.tname())
	return this.myname
}

type pname struct {
	parser
}

func (this *pname) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: Kindpname,
			Name: "pname",
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

func (this *singleToken) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: KindsingleToken,
			Name: "singleToken",
		},
	}

	this.mysingleToken.Parser = this.seq(
		this.tok("\""),
		this.char(),
		this.tok("\""))
	return this.mysingleToken
}

type tname struct {
	parser
}

func (this *tname) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: Kindtname,
			Name: "tname",
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

func (this *token) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
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
			Kind: Kindtoken,
			Name: "token",
		},
	}

	this.mytoken.Parser = this.seq(
		this.tok("\""),
		this.char(),
		this.rpt(
			this.char()),
		this.tok("\""))
	return this.mytoken
}

type Parser interface {
	LookAhead(char rune) (ok bool)
	Parse(input string) (unparsed string, ok bool)
	String() (s string)
}
type ProductionKind int

// parser is a struct
type parser struct {
	Kind   ProductionKind
	Name   string
	Parser Parser
}

// Newparser creates a new instance of parser
func Newparser() *parser {
	return &parser{}
}
