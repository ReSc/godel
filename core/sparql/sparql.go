package sparql

import ()

type Parser interface {
	LookAhead(s rune) (ok bool)
	Parse(s string) (unparsed string, ok bool)
	String() (s string)
}

// Production is a struct
type Production struct {
	Kind   ProductionKind
	Name   string
	Parser Parser
}

// NewProduction creates a new instance of Production
func NewProduction() *Production {
	return &Production{}
}

type ProductionKind int

type Sparql struct {
	myAdd                       *Add
	myAdditiveExpression        *AdditiveExpression
	myAggregate                 *Aggregate
	myArgList                   *ArgList
	myAskQuery                  *AskQuery
	myBaseDecl                  *BaseDecl
	myBind                      *Bind
	myBlankNode                 *BlankNode
	myBlankNodePropertyList     *BlankNodePropertyList
	myBlankNodePropertyListPath *BlankNodePropertyListPath
	myBooleanLiteral            *BooleanLiteral
	myBrackettedExpression      *BrackettedExpression
	myBuiltInCall               *BuiltInCall
	myClear                     *Clear
	myCollection                *Collection
	myCollectionPath            *CollectionPath
	myConditionalAndExpression  *ConditionalAndExpression
	myConditionalOrExpression   *ConditionalOrExpression
	myConstraint                *Constraint
	myConstructQuery            *ConstructQuery
	myConstructTemplate         *ConstructTemplate
	myConstructTriples          *ConstructTriples
	myCopy                      *Copy
	myCreate                    *Create
	myDataBlock                 *DataBlock
	myDataBlockValue            *DataBlockValue
	myDatasetClause             *DatasetClause
	myDefaultGraphClause        *DefaultGraphClause
	myDeleteClause              *DeleteClause
	myDeleteData                *DeleteData
	myDeleteWhere               *DeleteWhere
	myDescribeQuery             *DescribeQuery
	myDrop                      *Drop
	myExistsFunc                *ExistsFunc
	myExpression                *Expression
	myExpressionList            *ExpressionList
	myFilter                    *Filter
	myFunctionCall              *FunctionCall
	myGrammar                   *Grammar
	myGraphGraphPattern         *GraphGraphPattern
	myGraphNode                 *GraphNode
	myGraphNodePath             *GraphNodePath
	myGraphOrDefault            *GraphOrDefault
	myGraphPatternNotTriples    *GraphPatternNotTriples
	myGraphRef                  *GraphRef
	myGraphRefAll               *GraphRefAll
	myGraphTerm                 *GraphTerm
	myGroupClause               *GroupClause
	myGroupCondition            *GroupCondition
	myGroupGraphPattern         *GroupGraphPattern
	myGroupGraphPatternSub      *GroupGraphPatternSub
	myGroupOrUnionGraphPattern  *GroupOrUnionGraphPattern
	myHavingClause              *HavingClause
	myHavingCondition           *HavingCondition
	myInlineData                *InlineData
	myInlineDataFull            *InlineDataFull
	myInlineDataOneVar          *InlineDataOneVar
	myInsertClause              *InsertClause
	myInsertData                *InsertData
	myIri                       *Iri
	myIriOrFunction             *IriOrFunction
	myLimitClause               *LimitClause
	myLimitOffsetClauses        *LimitOffsetClauses
	myLoad                      *Load
	myMinusGraphPattern         *MinusGraphPattern
	myModify                    *Modify
	myMove                      *Move
	myMultiplicativeExpression  *MultiplicativeExpression
	myNamedGraphClause          *NamedGraphClause
	myNotExistsFunc             *NotExistsFunc
	myNumericExpression         *NumericExpression
	myNumericLiteral            *NumericLiteral
	myNumericLiteralNegative    *NumericLiteralNegative
	myNumericLiteralPositive    *NumericLiteralPositive
	myNumericLiteralUnsigned    *NumericLiteralUnsigned
	myObject                    *Object
	myObjectList                *ObjectList
	myObjectListPath            *ObjectListPath
	myObjectPath                *ObjectPath
	myOffsetClause              *OffsetClause
	myOptionalGraphPattern      *OptionalGraphPattern
	myOrderClause               *OrderClause
	myOrderCondition            *OrderCondition
	myPath                      *Path
	myPathAlternative           *PathAlternative
	myPathElt                   *PathElt
	myPathEltOrInverse          *PathEltOrInverse
	myPathMod                   *PathMod
	myPathNegatedPropertySet    *PathNegatedPropertySet
	myPathOneInPropertySet      *PathOneInPropertySet
	myPathPrimary               *PathPrimary
	myPathSequence              *PathSequence
	myPrefixDecl                *PrefixDecl
	myPrefixedName              *PrefixedName
	myPrimaryExpression         *PrimaryExpression
	myPrologue                  *Prologue
	myPropertyList              *PropertyList
	myPropertyListNotEmpty      *PropertyListNotEmpty
	myPropertyListPath          *PropertyListPath
	myPropertyListPathNotEmpty  *PropertyListPathNotEmpty
	myQuadData                  *QuadData
	myQuadPattern               *QuadPattern
	myQuads                     *Quads
	myQuadsNotTriples           *QuadsNotTriples
	myQuery                     *Query
	myQueryUnit                 *QueryUnit
	myRDFLiteral                *RDFLiteral
	myRegexExpression           *RegexExpression
	myRelationalExpression      *RelationalExpression
	mySelectClause              *SelectClause
	mySelectQuery               *SelectQuery
	myServiceGraphPattern       *ServiceGraphPattern
	mySolutionModifier          *SolutionModifier
	mySourceSelector            *SourceSelector
	myStrReplaceExpression      *StrReplaceExpression
	myString                    *String
	mySubSelect                 *SubSelect
	mySubstringExpression       *SubstringExpression
	myTriplesBlock              *TriplesBlock
	myTriplesNode               *TriplesNode
	myTriplesNodePath           *TriplesNodePath
	myTriplesSameSubject        *TriplesSameSubject
	myTriplesSameSubjectPath    *TriplesSameSubjectPath
	myTriplesTemplate           *TriplesTemplate
	myUnaryExpression           *UnaryExpression
	myUpdate                    *Update
	myUpdate1                   *Update1
	myUpdateUnit                *UpdateUnit
	myUsingClause               *UsingClause
	myValueLogical              *ValueLogical
	myValuesClause              *ValuesClause
	myVar                       *Var
	myVarOrIri                  *VarOrIri
	myVarOrTerm                 *VarOrTerm
	myVerb                      *Verb
	myVerbPath                  *VerbPath
	myVerbSimple                *VerbSimple
	myWhereClause               *WhereClause
	myANON                      *ANON
	myBLANK_NODE_LABEL          *BLANK_NODE_LABEL
	myDECIMAL                   *DECIMAL
	myDECIMAL_NEGATIVE          *DECIMAL_NEGATIVE
	myDECIMAL_POSITIVE          *DECIMAL_POSITIVE
	myDOUBLE                    *DOUBLE
	myDOUBLE_NEGATIVE           *DOUBLE_NEGATIVE
	myDOUBLE_POSITIVE           *DOUBLE_POSITIVE
	myECHAR                     *ECHAR
	myEXPONENT                  *EXPONENT
	myHEX                       *HEX
	myINTEGER                   *INTEGER
	myINTEGER_NEGATIVE          *INTEGER_NEGATIVE
	myINTEGER_POSITIVE          *INTEGER_POSITIVE
	myIRIREF                    *IRIREF
	myLANGTAG                   *LANGTAG
	myNIL                       *NIL
	myPERCENT                   *PERCENT
	myPLX                       *PLX
	myPNAME_LN                  *PNAME_LN
	myPNAME_NS                  *PNAME_NS
	myPN_CHARS                  *PN_CHARS
	myPN_CHARS_BASE             *PN_CHARS_BASE
	myPN_CHARS_U                *PN_CHARS_U
	myPN_LOCAL                  *PN_LOCAL
	myPN_LOCAL_ESC              *PN_LOCAL_ESC
	myPN_PREFIX                 *PN_PREFIX
	mySTRING_LITERAL1           *STRING_LITERAL1
	mySTRING_LITERAL2           *STRING_LITERAL2
	mySTRING_LITERAL_LONG1      *STRING_LITERAL_LONG1
	mySTRING_LITERAL_LONG2      *STRING_LITERAL_LONG2
	myVAR1                      *VAR1
	myVAR2                      *VAR2
	myVARNAME                   *VARNAME
	myWS                        *WS
}

func (this *Sparql) Syntax() string {

	output := ""
	output += this.Add().Syntax() + "\n"
	output += this.AdditiveExpression().Syntax() + "\n"
	output += this.Aggregate().Syntax() + "\n"
	output += this.ArgList().Syntax() + "\n"
	output += this.AskQuery().Syntax() + "\n"
	output += this.BaseDecl().Syntax() + "\n"
	output += this.Bind().Syntax() + "\n"
	output += this.BlankNode().Syntax() + "\n"
	output += this.BlankNodePropertyList().Syntax() + "\n"
	output += this.BlankNodePropertyListPath().Syntax() + "\n"
	output += this.BooleanLiteral().Syntax() + "\n"
	output += this.BrackettedExpression().Syntax() + "\n"
	output += this.BuiltInCall().Syntax() + "\n"
	output += this.Clear().Syntax() + "\n"
	output += this.Collection().Syntax() + "\n"
	output += this.CollectionPath().Syntax() + "\n"
	output += this.ConditionalAndExpression().Syntax() + "\n"
	output += this.ConditionalOrExpression().Syntax() + "\n"
	output += this.Constraint().Syntax() + "\n"
	output += this.ConstructQuery().Syntax() + "\n"
	output += this.ConstructTemplate().Syntax() + "\n"
	output += this.ConstructTriples().Syntax() + "\n"
	output += this.Copy().Syntax() + "\n"
	output += this.Create().Syntax() + "\n"
	output += this.DataBlock().Syntax() + "\n"
	output += this.DataBlockValue().Syntax() + "\n"
	output += this.DatasetClause().Syntax() + "\n"
	output += this.DefaultGraphClause().Syntax() + "\n"
	output += this.DeleteClause().Syntax() + "\n"
	output += this.DeleteData().Syntax() + "\n"
	output += this.DeleteWhere().Syntax() + "\n"
	output += this.DescribeQuery().Syntax() + "\n"
	output += this.Drop().Syntax() + "\n"
	output += this.ExistsFunc().Syntax() + "\n"
	output += this.Expression().Syntax() + "\n"
	output += this.ExpressionList().Syntax() + "\n"
	output += this.Filter().Syntax() + "\n"
	output += this.FunctionCall().Syntax() + "\n"
	output += this.Grammar().Syntax() + "\n"
	output += this.GraphGraphPattern().Syntax() + "\n"
	output += this.GraphNode().Syntax() + "\n"
	output += this.GraphNodePath().Syntax() + "\n"
	output += this.GraphOrDefault().Syntax() + "\n"
	output += this.GraphPatternNotTriples().Syntax() + "\n"
	output += this.GraphRef().Syntax() + "\n"
	output += this.GraphRefAll().Syntax() + "\n"
	output += this.GraphTerm().Syntax() + "\n"
	output += this.GroupClause().Syntax() + "\n"
	output += this.GroupCondition().Syntax() + "\n"
	output += this.GroupGraphPattern().Syntax() + "\n"
	output += this.GroupGraphPatternSub().Syntax() + "\n"
	output += this.GroupOrUnionGraphPattern().Syntax() + "\n"
	output += this.HavingClause().Syntax() + "\n"
	output += this.HavingCondition().Syntax() + "\n"
	output += this.InlineData().Syntax() + "\n"
	output += this.InlineDataFull().Syntax() + "\n"
	output += this.InlineDataOneVar().Syntax() + "\n"
	output += this.InsertClause().Syntax() + "\n"
	output += this.InsertData().Syntax() + "\n"
	output += this.Iri().Syntax() + "\n"
	output += this.IriOrFunction().Syntax() + "\n"
	output += this.LimitClause().Syntax() + "\n"
	output += this.LimitOffsetClauses().Syntax() + "\n"
	output += this.Load().Syntax() + "\n"
	output += this.MinusGraphPattern().Syntax() + "\n"
	output += this.Modify().Syntax() + "\n"
	output += this.Move().Syntax() + "\n"
	output += this.MultiplicativeExpression().Syntax() + "\n"
	output += this.NamedGraphClause().Syntax() + "\n"
	output += this.NotExistsFunc().Syntax() + "\n"
	output += this.NumericExpression().Syntax() + "\n"
	output += this.NumericLiteral().Syntax() + "\n"
	output += this.NumericLiteralNegative().Syntax() + "\n"
	output += this.NumericLiteralPositive().Syntax() + "\n"
	output += this.NumericLiteralUnsigned().Syntax() + "\n"
	output += this.Object().Syntax() + "\n"
	output += this.ObjectList().Syntax() + "\n"
	output += this.ObjectListPath().Syntax() + "\n"
	output += this.ObjectPath().Syntax() + "\n"
	output += this.OffsetClause().Syntax() + "\n"
	output += this.OptionalGraphPattern().Syntax() + "\n"
	output += this.OrderClause().Syntax() + "\n"
	output += this.OrderCondition().Syntax() + "\n"
	output += this.Path().Syntax() + "\n"
	output += this.PathAlternative().Syntax() + "\n"
	output += this.PathElt().Syntax() + "\n"
	output += this.PathEltOrInverse().Syntax() + "\n"
	output += this.PathMod().Syntax() + "\n"
	output += this.PathNegatedPropertySet().Syntax() + "\n"
	output += this.PathOneInPropertySet().Syntax() + "\n"
	output += this.PathPrimary().Syntax() + "\n"
	output += this.PathSequence().Syntax() + "\n"
	output += this.PrefixDecl().Syntax() + "\n"
	output += this.PrefixedName().Syntax() + "\n"
	output += this.PrimaryExpression().Syntax() + "\n"
	output += this.Prologue().Syntax() + "\n"
	output += this.PropertyList().Syntax() + "\n"
	output += this.PropertyListNotEmpty().Syntax() + "\n"
	output += this.PropertyListPath().Syntax() + "\n"
	output += this.PropertyListPathNotEmpty().Syntax() + "\n"
	output += this.QuadData().Syntax() + "\n"
	output += this.QuadPattern().Syntax() + "\n"
	output += this.Quads().Syntax() + "\n"
	output += this.QuadsNotTriples().Syntax() + "\n"
	output += this.Query().Syntax() + "\n"
	output += this.QueryUnit().Syntax() + "\n"
	output += this.RDFLiteral().Syntax() + "\n"
	output += this.RegexExpression().Syntax() + "\n"
	output += this.RelationalExpression().Syntax() + "\n"
	output += this.SelectClause().Syntax() + "\n"
	output += this.SelectQuery().Syntax() + "\n"
	output += this.ServiceGraphPattern().Syntax() + "\n"
	output += this.SolutionModifier().Syntax() + "\n"
	output += this.SourceSelector().Syntax() + "\n"
	output += this.StrReplaceExpression().Syntax() + "\n"
	output += this.String().Syntax() + "\n"
	output += this.SubSelect().Syntax() + "\n"
	output += this.SubstringExpression().Syntax() + "\n"
	output += this.TriplesBlock().Syntax() + "\n"
	output += this.TriplesNode().Syntax() + "\n"
	output += this.TriplesNodePath().Syntax() + "\n"
	output += this.TriplesSameSubject().Syntax() + "\n"
	output += this.TriplesSameSubjectPath().Syntax() + "\n"
	output += this.TriplesTemplate().Syntax() + "\n"
	output += this.UnaryExpression().Syntax() + "\n"
	output += this.Update().Syntax() + "\n"
	output += this.Update1().Syntax() + "\n"
	output += this.UpdateUnit().Syntax() + "\n"
	output += this.UsingClause().Syntax() + "\n"
	output += this.ValueLogical().Syntax() + "\n"
	output += this.ValuesClause().Syntax() + "\n"
	output += this.Var().Syntax() + "\n"
	output += this.VarOrIri().Syntax() + "\n"
	output += this.VarOrTerm().Syntax() + "\n"
	output += this.Verb().Syntax() + "\n"
	output += this.VerbPath().Syntax() + "\n"
	output += this.VerbSimple().Syntax() + "\n"
	output += this.WhereClause().Syntax() + "\n"
	output += this.ANON().Syntax() + "\n"
	output += this.BLANK_NODE_LABEL().Syntax() + "\n"
	output += this.DECIMAL().Syntax() + "\n"
	output += this.DECIMAL_NEGATIVE().Syntax() + "\n"
	output += this.DECIMAL_POSITIVE().Syntax() + "\n"
	output += this.DOUBLE().Syntax() + "\n"
	output += this.DOUBLE_NEGATIVE().Syntax() + "\n"
	output += this.DOUBLE_POSITIVE().Syntax() + "\n"
	output += this.ECHAR().Syntax() + "\n"
	output += this.EXPONENT().Syntax() + "\n"
	output += this.HEX().Syntax() + "\n"
	output += this.INTEGER().Syntax() + "\n"
	output += this.INTEGER_NEGATIVE().Syntax() + "\n"
	output += this.INTEGER_POSITIVE().Syntax() + "\n"
	output += this.IRIREF().Syntax() + "\n"
	output += this.LANGTAG().Syntax() + "\n"
	output += this.NIL().Syntax() + "\n"
	output += this.PERCENT().Syntax() + "\n"
	output += this.PLX().Syntax() + "\n"
	output += this.PNAME_LN().Syntax() + "\n"
	output += this.PNAME_NS().Syntax() + "\n"
	output += this.PN_CHARS().Syntax() + "\n"
	output += this.PN_CHARS_BASE().Syntax() + "\n"
	output += this.PN_CHARS_U().Syntax() + "\n"
	output += this.PN_LOCAL().Syntax() + "\n"
	output += this.PN_LOCAL_ESC().Syntax() + "\n"
	output += this.PN_PREFIX().Syntax() + "\n"
	output += this.STRING_LITERAL1().Syntax() + "\n"
	output += this.STRING_LITERAL2().Syntax() + "\n"
	output += this.STRING_LITERAL_LONG1().Syntax() + "\n"
	output += this.STRING_LITERAL_LONG2().Syntax() + "\n"
	output += this.VAR1().Syntax() + "\n"
	output += this.VAR2().Syntax() + "\n"
	output += this.VARNAME().Syntax() + "\n"
	output += this.WS().Syntax() + "\n"
	return output

}

const (
	KindAdd ProductionKind = iota
	KindAdditiveExpression
	KindAggregate
	KindArgList
	KindAskQuery
	KindBaseDecl
	KindBind
	KindBlankNode
	KindBlankNodePropertyList
	KindBlankNodePropertyListPath
	KindBooleanLiteral
	KindBrackettedExpression
	KindBuiltInCall
	KindClear
	KindCollection
	KindCollectionPath
	KindConditionalAndExpression
	KindConditionalOrExpression
	KindConstraint
	KindConstructQuery
	KindConstructTemplate
	KindConstructTriples
	KindCopy
	KindCreate
	KindDataBlock
	KindDataBlockValue
	KindDatasetClause
	KindDefaultGraphClause
	KindDeleteClause
	KindDeleteData
	KindDeleteWhere
	KindDescribeQuery
	KindDrop
	KindExistsFunc
	KindExpression
	KindExpressionList
	KindFilter
	KindFunctionCall
	KindGrammar
	KindGraphGraphPattern
	KindGraphNode
	KindGraphNodePath
	KindGraphOrDefault
	KindGraphPatternNotTriples
	KindGraphRef
	KindGraphRefAll
	KindGraphTerm
	KindGroupClause
	KindGroupCondition
	KindGroupGraphPattern
	KindGroupGraphPatternSub
	KindGroupOrUnionGraphPattern
	KindHavingClause
	KindHavingCondition
	KindInlineData
	KindInlineDataFull
	KindInlineDataOneVar
	KindInsertClause
	KindInsertData
	KindIri
	KindIriOrFunction
	KindLimitClause
	KindLimitOffsetClauses
	KindLoad
	KindMinusGraphPattern
	KindModify
	KindMove
	KindMultiplicativeExpression
	KindNamedGraphClause
	KindNotExistsFunc
	KindNumericExpression
	KindNumericLiteral
	KindNumericLiteralNegative
	KindNumericLiteralPositive
	KindNumericLiteralUnsigned
	KindObject
	KindObjectList
	KindObjectListPath
	KindObjectPath
	KindOffsetClause
	KindOptionalGraphPattern
	KindOrderClause
	KindOrderCondition
	KindPath
	KindPathAlternative
	KindPathElt
	KindPathEltOrInverse
	KindPathMod
	KindPathNegatedPropertySet
	KindPathOneInPropertySet
	KindPathPrimary
	KindPathSequence
	KindPrefixDecl
	KindPrefixedName
	KindPrimaryExpression
	KindPrologue
	KindPropertyList
	KindPropertyListNotEmpty
	KindPropertyListPath
	KindPropertyListPathNotEmpty
	KindQuadData
	KindQuadPattern
	KindQuads
	KindQuadsNotTriples
	KindQuery
	KindQueryUnit
	KindRDFLiteral
	KindRegexExpression
	KindRelationalExpression
	KindSelectClause
	KindSelectQuery
	KindServiceGraphPattern
	KindSolutionModifier
	KindSourceSelector
	KindStrReplaceExpression
	KindString
	KindSubSelect
	KindSubstringExpression
	KindTriplesBlock
	KindTriplesNode
	KindTriplesNodePath
	KindTriplesSameSubject
	KindTriplesSameSubjectPath
	KindTriplesTemplate
	KindUnaryExpression
	KindUpdate
	KindUpdate1
	KindUpdateUnit
	KindUsingClause
	KindValueLogical
	KindValuesClause
	KindVar
	KindVarOrIri
	KindVarOrTerm
	KindVerb
	KindVerbPath
	KindVerbSimple
	KindWhereClause
	KindANON
	KindBLANK_NODE_LABEL
	KindDECIMAL
	KindDECIMAL_NEGATIVE
	KindDECIMAL_POSITIVE
	KindDOUBLE
	KindDOUBLE_NEGATIVE
	KindDOUBLE_POSITIVE
	KindECHAR
	KindEXPONENT
	KindHEX
	KindINTEGER
	KindINTEGER_NEGATIVE
	KindINTEGER_POSITIVE
	KindIRIREF
	KindLANGTAG
	KindNIL
	KindPERCENT
	KindPLX
	KindPNAME_LN
	KindPNAME_NS
	KindPN_CHARS
	KindPN_CHARS_BASE
	KindPN_CHARS_U
	KindPN_LOCAL
	KindPN_LOCAL_ESC
	KindPN_PREFIX
	KindSTRING_LITERAL1
	KindSTRING_LITERAL2
	KindSTRING_LITERAL_LONG1
	KindSTRING_LITERAL_LONG2
	KindVAR1
	KindVAR2
	KindVARNAME
	KindWS
)

type Add struct {
	Production
}

func (this *Add) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Add) String() string {
	return this.Name
}
func (this *Add) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Add() *Add {
	if this.myAdd != nil {
		return this.myAdd
	}
	this.myAdd = &Add{
		Production: Production{
			Kind: KindAdd,
			Name: "Add",
		},
	}

	this.myAdd.Parser = this.seq(
		this.tok("ADD"),
		this.opt(
			this.tok("SILENT")),
		this.GraphOrDefault(),
		this.tok("TO"),
		this.GraphOrDefault())
	return this.myAdd
}

type AdditiveExpression struct {
	Production
}

func (this *AdditiveExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *AdditiveExpression) String() string {
	return this.Name
}
func (this *AdditiveExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) AdditiveExpression() *AdditiveExpression {
	if this.myAdditiveExpression != nil {
		return this.myAdditiveExpression
	}
	this.myAdditiveExpression = &AdditiveExpression{
		Production: Production{
			Kind: KindAdditiveExpression,
			Name: "AdditiveExpression",
		},
	}

	this.myAdditiveExpression.Parser = this.seq(
		this.MultiplicativeExpression(),
		this.rpt(
			this.alt(
				this.seq(
					this.tok("+"),
					this.MultiplicativeExpression()),
				this.seq(
					this.tok("-"),
					this.MultiplicativeExpression()),
				this.seq(
					this.grp(
						this.alt(
							this.NumericLiteralPositive(),
							this.NumericLiteralNegative())),
					this.rpt(
						this.alt(
							this.grp(
								this.seq(
									this.tok("*"),
									this.UnaryExpression())),
							this.grp(
								this.seq(
									this.tok("/"),
									this.UnaryExpression()))))))))
	return this.myAdditiveExpression
}

type Aggregate struct {
	Production
}

func (this *Aggregate) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Aggregate) String() string {
	return this.Name
}
func (this *Aggregate) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Aggregate() *Aggregate {
	if this.myAggregate != nil {
		return this.myAggregate
	}
	this.myAggregate = &Aggregate{
		Production: Production{
			Kind: KindAggregate,
			Name: "Aggregate",
		},
	}

	this.myAggregate.Parser = this.alt(
		this.seq(
			this.tok("COUNT"),
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.grp(
				this.alt(
					this.tok("*"),
					this.Expression())),
			this.tok(")")),
		this.seq(
			this.tok("SUM"),
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("MIN"),
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("MAX"),
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("AVG"),
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("SAMPLE"),
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("GROUP_CONCAT"),
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.Expression(),
			this.opt(
				this.seq(
					this.tok(";"),
					this.tok("SEPARATOR"),
					this.tok("="),
					this.String())),
			this.tok(")")))
	return this.myAggregate
}

type ArgList struct {
	Production
}

func (this *ArgList) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ArgList) String() string {
	return this.Name
}
func (this *ArgList) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ArgList() *ArgList {
	if this.myArgList != nil {
		return this.myArgList
	}
	this.myArgList = &ArgList{
		Production: Production{
			Kind: KindArgList,
			Name: "ArgList",
		},
	}

	this.myArgList.Parser = this.alt(
		this.NIL(),
		this.seq(
			this.tok("("),
			this.opt(
				this.tok("DISTINCT")),
			this.Expression(),
			this.rpt(
				this.seq(
					this.tok(","),
					this.Expression())),
			this.tok(")")))
	return this.myArgList
}

type AskQuery struct {
	Production
}

func (this *AskQuery) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *AskQuery) String() string {
	return this.Name
}
func (this *AskQuery) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) AskQuery() *AskQuery {
	if this.myAskQuery != nil {
		return this.myAskQuery
	}
	this.myAskQuery = &AskQuery{
		Production: Production{
			Kind: KindAskQuery,
			Name: "AskQuery",
		},
	}

	this.myAskQuery.Parser = this.seq(
		this.tok("ASK"),
		this.opt(
			this.rpt(
				this.DatasetClause())),
		this.WhereClause(),
		this.SolutionModifier())
	return this.myAskQuery
}

type BaseDecl struct {
	Production
}

func (this *BaseDecl) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BaseDecl) String() string {
	return this.Name
}
func (this *BaseDecl) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BaseDecl() *BaseDecl {
	if this.myBaseDecl != nil {
		return this.myBaseDecl
	}
	this.myBaseDecl = &BaseDecl{
		Production: Production{
			Kind: KindBaseDecl,
			Name: "BaseDecl",
		},
	}

	this.myBaseDecl.Parser = this.seq(
		this.tok("BASE"),
		this.IRIREF())
	return this.myBaseDecl
}

type Bind struct {
	Production
}

func (this *Bind) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Bind) String() string {
	return this.Name
}
func (this *Bind) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Bind() *Bind {
	if this.myBind != nil {
		return this.myBind
	}
	this.myBind = &Bind{
		Production: Production{
			Kind: KindBind,
			Name: "Bind",
		},
	}

	this.myBind.Parser = this.seq(
		this.tok("BIND"),
		this.tok("("),
		this.Expression(),
		this.tok("AS"),
		this.Var(),
		this.tok(")"))
	return this.myBind
}

type BlankNode struct {
	Production
}

func (this *BlankNode) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BlankNode) String() string {
	return this.Name
}
func (this *BlankNode) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BlankNode() *BlankNode {
	if this.myBlankNode != nil {
		return this.myBlankNode
	}
	this.myBlankNode = &BlankNode{
		Production: Production{
			Kind: KindBlankNode,
			Name: "BlankNode",
		},
	}

	this.myBlankNode.Parser = this.alt(
		this.BLANK_NODE_LABEL(),
		this.ANON())
	return this.myBlankNode
}

type BlankNodePropertyList struct {
	Production
}

func (this *BlankNodePropertyList) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BlankNodePropertyList) String() string {
	return this.Name
}
func (this *BlankNodePropertyList) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BlankNodePropertyList() *BlankNodePropertyList {
	if this.myBlankNodePropertyList != nil {
		return this.myBlankNodePropertyList
	}
	this.myBlankNodePropertyList = &BlankNodePropertyList{
		Production: Production{
			Kind: KindBlankNodePropertyList,
			Name: "BlankNodePropertyList",
		},
	}

	this.myBlankNodePropertyList.Parser = this.seq(
		this.tok("["),
		this.PropertyListNotEmpty(),
		this.tok("]"))
	return this.myBlankNodePropertyList
}

type BlankNodePropertyListPath struct {
	Production
}

func (this *BlankNodePropertyListPath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BlankNodePropertyListPath) String() string {
	return this.Name
}
func (this *BlankNodePropertyListPath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BlankNodePropertyListPath() *BlankNodePropertyListPath {
	if this.myBlankNodePropertyListPath != nil {
		return this.myBlankNodePropertyListPath
	}
	this.myBlankNodePropertyListPath = &BlankNodePropertyListPath{
		Production: Production{
			Kind: KindBlankNodePropertyListPath,
			Name: "BlankNodePropertyListPath",
		},
	}

	this.myBlankNodePropertyListPath.Parser = this.seq(
		this.tok("["),
		this.PropertyListPathNotEmpty(),
		this.tok("]"))
	return this.myBlankNodePropertyListPath
}

type BooleanLiteral struct {
	Production
}

func (this *BooleanLiteral) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BooleanLiteral) String() string {
	return this.Name
}
func (this *BooleanLiteral) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BooleanLiteral() *BooleanLiteral {
	if this.myBooleanLiteral != nil {
		return this.myBooleanLiteral
	}
	this.myBooleanLiteral = &BooleanLiteral{
		Production: Production{
			Kind: KindBooleanLiteral,
			Name: "BooleanLiteral",
		},
	}

	this.myBooleanLiteral.Parser = this.alt(
		this.tok("true"),
		this.tok("false"))
	return this.myBooleanLiteral
}

type BrackettedExpression struct {
	Production
}

func (this *BrackettedExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BrackettedExpression) String() string {
	return this.Name
}
func (this *BrackettedExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BrackettedExpression() *BrackettedExpression {
	if this.myBrackettedExpression != nil {
		return this.myBrackettedExpression
	}
	this.myBrackettedExpression = &BrackettedExpression{
		Production: Production{
			Kind: KindBrackettedExpression,
			Name: "BrackettedExpression",
		},
	}

	this.myBrackettedExpression.Parser = this.seq(
		this.tok("("),
		this.Expression(),
		this.tok(")"))
	return this.myBrackettedExpression
}

type BuiltInCall struct {
	Production
}

func (this *BuiltInCall) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BuiltInCall) String() string {
	return this.Name
}
func (this *BuiltInCall) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BuiltInCall() *BuiltInCall {
	if this.myBuiltInCall != nil {
		return this.myBuiltInCall
	}
	this.myBuiltInCall = &BuiltInCall{
		Production: Production{
			Kind: KindBuiltInCall,
			Name: "BuiltInCall",
		},
	}

	this.myBuiltInCall.Parser = this.alt(
		this.Aggregate(),
		this.seq(
			this.tok("STR"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("LANG"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("LANGMATCHES"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("DATATYPE"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("BOUND"),
			this.tok("("),
			this.Var(),
			this.tok(")")),
		this.seq(
			this.tok("IRI"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("URI"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("BNODE"),
			this.grp(
				this.alt(
					this.seq(
						this.tok("("),
						this.Expression(),
						this.tok(")")),
					this.NIL()))),
		this.seq(
			this.tok("RAND"),
			this.NIL()),
		this.seq(
			this.tok("ABS"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("CEIL"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("FLOOR"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("ROUND"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("CONCAT"),
			this.ExpressionList()),
		this.SubstringExpression(),
		this.seq(
			this.tok("STRLEN"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.StrReplaceExpression(),
		this.seq(
			this.tok("UCASE"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("LCASE"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("ENCODE_FOR_URI"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("CONTAINS"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("STRSTARTS"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("STRENDS"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("STRBEFORE"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("STRAFTER"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("YEAR"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("MONTH"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("DAY"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("HOURS"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("MINUTES"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("SECONDS"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("TIMEZONE"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("TZ"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("NOW"),
			this.NIL()),
		this.seq(
			this.tok("UUID"),
			this.NIL()),
		this.seq(
			this.tok("STRUUID"),
			this.NIL()),
		this.seq(
			this.tok("MD5"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("SHA1"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("SHA256"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("SHA384"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("SHA512"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("COALESCE"),
			this.ExpressionList()),
		this.seq(
			this.tok("IF"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("STRLANG"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("STRDT"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("sameTerm"),
			this.tok("("),
			this.Expression(),
			this.tok(","),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("isIri"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("isURI"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("isBLANK"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("isLITERAL"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.seq(
			this.tok("isNUMERIC"),
			this.tok("("),
			this.Expression(),
			this.tok(")")),
		this.RegexExpression(),
		this.ExistsFunc(),
		this.NotExistsFunc())
	return this.myBuiltInCall
}

type Clear struct {
	Production
}

func (this *Clear) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Clear) String() string {
	return this.Name
}
func (this *Clear) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Clear() *Clear {
	if this.myClear != nil {
		return this.myClear
	}
	this.myClear = &Clear{
		Production: Production{
			Kind: KindClear,
			Name: "Clear",
		},
	}

	this.myClear.Parser = this.seq(
		this.tok("CLEAR"),
		this.opt(
			this.tok("SILENT")),
		this.GraphRefAll())
	return this.myClear
}

type Collection struct {
	Production
}

func (this *Collection) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Collection) String() string {
	return this.Name
}
func (this *Collection) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Collection() *Collection {
	if this.myCollection != nil {
		return this.myCollection
	}
	this.myCollection = &Collection{
		Production: Production{
			Kind: KindCollection,
			Name: "Collection",
		},
	}

	this.myCollection.Parser = this.seq(
		this.tok("("),
		this.rpt(
			this.GraphNode()),
		this.tok(")"))
	return this.myCollection
}

type CollectionPath struct {
	Production
}

func (this *CollectionPath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *CollectionPath) String() string {
	return this.Name
}
func (this *CollectionPath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) CollectionPath() *CollectionPath {
	if this.myCollectionPath != nil {
		return this.myCollectionPath
	}
	this.myCollectionPath = &CollectionPath{
		Production: Production{
			Kind: KindCollectionPath,
			Name: "CollectionPath",
		},
	}

	this.myCollectionPath.Parser = this.seq(
		this.tok("("),
		this.rpt(
			this.GraphNodePath()),
		this.tok(")"))
	return this.myCollectionPath
}

type ConditionalAndExpression struct {
	Production
}

func (this *ConditionalAndExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ConditionalAndExpression) String() string {
	return this.Name
}
func (this *ConditionalAndExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ConditionalAndExpression() *ConditionalAndExpression {
	if this.myConditionalAndExpression != nil {
		return this.myConditionalAndExpression
	}
	this.myConditionalAndExpression = &ConditionalAndExpression{
		Production: Production{
			Kind: KindConditionalAndExpression,
			Name: "ConditionalAndExpression",
		},
	}

	this.myConditionalAndExpression.Parser = this.seq(
		this.ValueLogical(),
		this.rpt(
			this.seq(
				this.tok("&&"),
				this.ValueLogical())))
	return this.myConditionalAndExpression
}

type ConditionalOrExpression struct {
	Production
}

func (this *ConditionalOrExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ConditionalOrExpression) String() string {
	return this.Name
}
func (this *ConditionalOrExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ConditionalOrExpression() *ConditionalOrExpression {
	if this.myConditionalOrExpression != nil {
		return this.myConditionalOrExpression
	}
	this.myConditionalOrExpression = &ConditionalOrExpression{
		Production: Production{
			Kind: KindConditionalOrExpression,
			Name: "ConditionalOrExpression",
		},
	}

	this.myConditionalOrExpression.Parser = this.seq(
		this.ConditionalAndExpression(),
		this.rpt(
			this.seq(
				this.tok("||"),
				this.ConditionalAndExpression())))
	return this.myConditionalOrExpression
}

type Constraint struct {
	Production
}

func (this *Constraint) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Constraint) String() string {
	return this.Name
}
func (this *Constraint) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Constraint() *Constraint {
	if this.myConstraint != nil {
		return this.myConstraint
	}
	this.myConstraint = &Constraint{
		Production: Production{
			Kind: KindConstraint,
			Name: "Constraint",
		},
	}

	this.myConstraint.Parser = this.alt(
		this.BrackettedExpression(),
		this.BuiltInCall(),
		this.FunctionCall())
	return this.myConstraint
}

type ConstructQuery struct {
	Production
}

func (this *ConstructQuery) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ConstructQuery) String() string {
	return this.Name
}
func (this *ConstructQuery) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ConstructQuery() *ConstructQuery {
	if this.myConstructQuery != nil {
		return this.myConstructQuery
	}
	this.myConstructQuery = &ConstructQuery{
		Production: Production{
			Kind: KindConstructQuery,
			Name: "ConstructQuery",
		},
	}

	this.myConstructQuery.Parser = this.seq(
		this.tok("CONSTRUCT"),
		this.grp(
			this.alt(
				this.seq(
					this.ConstructTemplate(),
					this.opt(
						this.rpt(
							this.DatasetClause())),
					this.WhereClause(),
					this.SolutionModifier()),
				this.seq(
					this.opt(
						this.rpt(
							this.DatasetClause())),
					this.tok("WHERE"),
					this.tok("{"),
					this.opt(
						this.TriplesTemplate()),
					this.tok("}"),
					this.SolutionModifier()))))
	return this.myConstructQuery
}

type ConstructTemplate struct {
	Production
}

func (this *ConstructTemplate) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ConstructTemplate) String() string {
	return this.Name
}
func (this *ConstructTemplate) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ConstructTemplate() *ConstructTemplate {
	if this.myConstructTemplate != nil {
		return this.myConstructTemplate
	}
	this.myConstructTemplate = &ConstructTemplate{
		Production: Production{
			Kind: KindConstructTemplate,
			Name: "ConstructTemplate",
		},
	}

	this.myConstructTemplate.Parser = this.seq(
		this.tok("{"),
		this.opt(
			this.ConstructTriples()),
		this.tok("}"))
	return this.myConstructTemplate
}

type ConstructTriples struct {
	Production
}

func (this *ConstructTriples) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ConstructTriples) String() string {
	return this.Name
}
func (this *ConstructTriples) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ConstructTriples() *ConstructTriples {
	if this.myConstructTriples != nil {
		return this.myConstructTriples
	}
	this.myConstructTriples = &ConstructTriples{
		Production: Production{
			Kind: KindConstructTriples,
			Name: "ConstructTriples",
		},
	}

	this.myConstructTriples.Parser = this.seq(
		this.TriplesSameSubject(),
		this.opt(
			this.seq(
				this.tok("."),
				this.opt(
					this.ConstructTriples()))))
	return this.myConstructTriples
}

type Copy struct {
	Production
}

func (this *Copy) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Copy) String() string {
	return this.Name
}
func (this *Copy) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Copy() *Copy {
	if this.myCopy != nil {
		return this.myCopy
	}
	this.myCopy = &Copy{
		Production: Production{
			Kind: KindCopy,
			Name: "Copy",
		},
	}

	this.myCopy.Parser = this.seq(
		this.tok("COPY"),
		this.opt(
			this.tok("SILENT")),
		this.GraphOrDefault(),
		this.tok("TO"),
		this.GraphOrDefault())
	return this.myCopy
}

type Create struct {
	Production
}

func (this *Create) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Create) String() string {
	return this.Name
}
func (this *Create) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Create() *Create {
	if this.myCreate != nil {
		return this.myCreate
	}
	this.myCreate = &Create{
		Production: Production{
			Kind: KindCreate,
			Name: "Create",
		},
	}

	this.myCreate.Parser = this.seq(
		this.tok("CREATE"),
		this.opt(
			this.tok("SILENT")),
		this.GraphRef())
	return this.myCreate
}

type DataBlock struct {
	Production
}

func (this *DataBlock) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DataBlock) String() string {
	return this.Name
}
func (this *DataBlock) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DataBlock() *DataBlock {
	if this.myDataBlock != nil {
		return this.myDataBlock
	}
	this.myDataBlock = &DataBlock{
		Production: Production{
			Kind: KindDataBlock,
			Name: "DataBlock",
		},
	}

	this.myDataBlock.Parser = this.alt(
		this.InlineDataOneVar(),
		this.InlineDataFull())
	return this.myDataBlock
}

type DataBlockValue struct {
	Production
}

func (this *DataBlockValue) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DataBlockValue) String() string {
	return this.Name
}
func (this *DataBlockValue) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DataBlockValue() *DataBlockValue {
	if this.myDataBlockValue != nil {
		return this.myDataBlockValue
	}
	this.myDataBlockValue = &DataBlockValue{
		Production: Production{
			Kind: KindDataBlockValue,
			Name: "DataBlockValue",
		},
	}

	this.myDataBlockValue.Parser = this.alt(
		this.Iri(),
		this.RDFLiteral(),
		this.NumericLiteral(),
		this.BooleanLiteral(),
		this.tok("UNDEF"))
	return this.myDataBlockValue
}

type DatasetClause struct {
	Production
}

func (this *DatasetClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DatasetClause) String() string {
	return this.Name
}
func (this *DatasetClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DatasetClause() *DatasetClause {
	if this.myDatasetClause != nil {
		return this.myDatasetClause
	}
	this.myDatasetClause = &DatasetClause{
		Production: Production{
			Kind: KindDatasetClause,
			Name: "DatasetClause",
		},
	}

	this.myDatasetClause.Parser = this.seq(
		this.tok("FROM"),
		this.grp(
			this.alt(
				this.DefaultGraphClause(),
				this.NamedGraphClause())))
	return this.myDatasetClause
}

type DefaultGraphClause struct {
	Production
}

func (this *DefaultGraphClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DefaultGraphClause) String() string {
	return this.Name
}
func (this *DefaultGraphClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DefaultGraphClause() *DefaultGraphClause {
	if this.myDefaultGraphClause != nil {
		return this.myDefaultGraphClause
	}
	this.myDefaultGraphClause = &DefaultGraphClause{
		Production: Production{
			Kind: KindDefaultGraphClause,
			Name: "DefaultGraphClause",
		},
	}

	this.myDefaultGraphClause.Parser = this.SourceSelector()
	return this.myDefaultGraphClause
}

type DeleteClause struct {
	Production
}

func (this *DeleteClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DeleteClause) String() string {
	return this.Name
}
func (this *DeleteClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DeleteClause() *DeleteClause {
	if this.myDeleteClause != nil {
		return this.myDeleteClause
	}
	this.myDeleteClause = &DeleteClause{
		Production: Production{
			Kind: KindDeleteClause,
			Name: "DeleteClause",
		},
	}

	this.myDeleteClause.Parser = this.seq(
		this.tok("DELETE"),
		this.QuadPattern())
	return this.myDeleteClause
}

type DeleteData struct {
	Production
}

func (this *DeleteData) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DeleteData) String() string {
	return this.Name
}
func (this *DeleteData) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DeleteData() *DeleteData {
	if this.myDeleteData != nil {
		return this.myDeleteData
	}
	this.myDeleteData = &DeleteData{
		Production: Production{
			Kind: KindDeleteData,
			Name: "DeleteData",
		},
	}

	this.myDeleteData.Parser = this.seq(
		this.tok("DELETE DATA"),
		this.QuadData())
	return this.myDeleteData
}

type DeleteWhere struct {
	Production
}

func (this *DeleteWhere) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DeleteWhere) String() string {
	return this.Name
}
func (this *DeleteWhere) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DeleteWhere() *DeleteWhere {
	if this.myDeleteWhere != nil {
		return this.myDeleteWhere
	}
	this.myDeleteWhere = &DeleteWhere{
		Production: Production{
			Kind: KindDeleteWhere,
			Name: "DeleteWhere",
		},
	}

	this.myDeleteWhere.Parser = this.seq(
		this.tok("DELETE WHERE"),
		this.QuadPattern())
	return this.myDeleteWhere
}

type DescribeQuery struct {
	Production
}

func (this *DescribeQuery) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DescribeQuery) String() string {
	return this.Name
}
func (this *DescribeQuery) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DescribeQuery() *DescribeQuery {
	if this.myDescribeQuery != nil {
		return this.myDescribeQuery
	}
	this.myDescribeQuery = &DescribeQuery{
		Production: Production{
			Kind: KindDescribeQuery,
			Name: "DescribeQuery",
		},
	}

	this.myDescribeQuery.Parser = this.seq(
		this.tok("DESCRIBE"),
		this.grp(
			this.alt(
				this.rpt(
					this.VarOrIri()),
				this.tok("*"))),
		this.opt(
			this.rpt(
				this.DatasetClause())),
		this.opt(
			this.WhereClause()),
		this.SolutionModifier())
	return this.myDescribeQuery
}

type Drop struct {
	Production
}

func (this *Drop) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Drop) String() string {
	return this.Name
}
func (this *Drop) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Drop() *Drop {
	if this.myDrop != nil {
		return this.myDrop
	}
	this.myDrop = &Drop{
		Production: Production{
			Kind: KindDrop,
			Name: "Drop",
		},
	}

	this.myDrop.Parser = this.seq(
		this.tok("DROP"),
		this.opt(
			this.tok("SILENT")),
		this.GraphRefAll())
	return this.myDrop
}

type ExistsFunc struct {
	Production
}

func (this *ExistsFunc) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ExistsFunc) String() string {
	return this.Name
}
func (this *ExistsFunc) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ExistsFunc() *ExistsFunc {
	if this.myExistsFunc != nil {
		return this.myExistsFunc
	}
	this.myExistsFunc = &ExistsFunc{
		Production: Production{
			Kind: KindExistsFunc,
			Name: "ExistsFunc",
		},
	}

	this.myExistsFunc.Parser = this.seq(
		this.tok("EXISTS"),
		this.GroupGraphPattern())
	return this.myExistsFunc
}

type Expression struct {
	Production
}

func (this *Expression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Expression) String() string {
	return this.Name
}
func (this *Expression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Expression() *Expression {
	if this.myExpression != nil {
		return this.myExpression
	}
	this.myExpression = &Expression{
		Production: Production{
			Kind: KindExpression,
			Name: "Expression",
		},
	}

	this.myExpression.Parser = this.ConditionalOrExpression()
	return this.myExpression
}

type ExpressionList struct {
	Production
}

func (this *ExpressionList) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ExpressionList) String() string {
	return this.Name
}
func (this *ExpressionList) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ExpressionList() *ExpressionList {
	if this.myExpressionList != nil {
		return this.myExpressionList
	}
	this.myExpressionList = &ExpressionList{
		Production: Production{
			Kind: KindExpressionList,
			Name: "ExpressionList",
		},
	}

	this.myExpressionList.Parser = this.alt(
		this.NIL(),
		this.seq(
			this.tok("("),
			this.Expression(),
			this.rpt(
				this.seq(
					this.tok(","),
					this.Expression())),
			this.tok(")")))
	return this.myExpressionList
}

type Filter struct {
	Production
}

func (this *Filter) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Filter) String() string {
	return this.Name
}
func (this *Filter) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Filter() *Filter {
	if this.myFilter != nil {
		return this.myFilter
	}
	this.myFilter = &Filter{
		Production: Production{
			Kind: KindFilter,
			Name: "Filter",
		},
	}

	this.myFilter.Parser = this.seq(
		this.tok("FILTER"),
		this.Constraint())
	return this.myFilter
}

type FunctionCall struct {
	Production
}

func (this *FunctionCall) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *FunctionCall) String() string {
	return this.Name
}
func (this *FunctionCall) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) FunctionCall() *FunctionCall {
	if this.myFunctionCall != nil {
		return this.myFunctionCall
	}
	this.myFunctionCall = &FunctionCall{
		Production: Production{
			Kind: KindFunctionCall,
			Name: "FunctionCall",
		},
	}

	this.myFunctionCall.Parser = this.seq(
		this.Iri(),
		this.ArgList())
	return this.myFunctionCall
}

type Grammar struct {
	Production
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
func (this *Sparql) Grammar() *Grammar {
	if this.myGrammar != nil {
		return this.myGrammar
	}
	this.myGrammar = &Grammar{
		Production: Production{
			Kind: KindGrammar,
			Name: "Grammar",
		},
	}

	this.myGrammar.Parser = this.alt(
		this.QueryUnit(),
		this.UpdateUnit())
	return this.myGrammar
}

type GraphGraphPattern struct {
	Production
}

func (this *GraphGraphPattern) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphGraphPattern) String() string {
	return this.Name
}
func (this *GraphGraphPattern) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphGraphPattern() *GraphGraphPattern {
	if this.myGraphGraphPattern != nil {
		return this.myGraphGraphPattern
	}
	this.myGraphGraphPattern = &GraphGraphPattern{
		Production: Production{
			Kind: KindGraphGraphPattern,
			Name: "GraphGraphPattern",
		},
	}

	this.myGraphGraphPattern.Parser = this.seq(
		this.tok("GRAPH"),
		this.VarOrIri(),
		this.GroupGraphPattern())
	return this.myGraphGraphPattern
}

type GraphNode struct {
	Production
}

func (this *GraphNode) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphNode) String() string {
	return this.Name
}
func (this *GraphNode) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphNode() *GraphNode {
	if this.myGraphNode != nil {
		return this.myGraphNode
	}
	this.myGraphNode = &GraphNode{
		Production: Production{
			Kind: KindGraphNode,
			Name: "GraphNode",
		},
	}

	this.myGraphNode.Parser = this.alt(
		this.VarOrTerm(),
		this.TriplesNode())
	return this.myGraphNode
}

type GraphNodePath struct {
	Production
}

func (this *GraphNodePath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphNodePath) String() string {
	return this.Name
}
func (this *GraphNodePath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphNodePath() *GraphNodePath {
	if this.myGraphNodePath != nil {
		return this.myGraphNodePath
	}
	this.myGraphNodePath = &GraphNodePath{
		Production: Production{
			Kind: KindGraphNodePath,
			Name: "GraphNodePath",
		},
	}

	this.myGraphNodePath.Parser = this.alt(
		this.VarOrTerm(),
		this.TriplesNodePath())
	return this.myGraphNodePath
}

type GraphOrDefault struct {
	Production
}

func (this *GraphOrDefault) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphOrDefault) String() string {
	return this.Name
}
func (this *GraphOrDefault) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphOrDefault() *GraphOrDefault {
	if this.myGraphOrDefault != nil {
		return this.myGraphOrDefault
	}
	this.myGraphOrDefault = &GraphOrDefault{
		Production: Production{
			Kind: KindGraphOrDefault,
			Name: "GraphOrDefault",
		},
	}

	this.myGraphOrDefault.Parser = this.alt(
		this.tok("DEFAULT"),
		this.seq(
			this.opt(
				this.tok("GRAPH")),
			this.Iri()))
	return this.myGraphOrDefault
}

type GraphPatternNotTriples struct {
	Production
}

func (this *GraphPatternNotTriples) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphPatternNotTriples) String() string {
	return this.Name
}
func (this *GraphPatternNotTriples) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphPatternNotTriples() *GraphPatternNotTriples {
	if this.myGraphPatternNotTriples != nil {
		return this.myGraphPatternNotTriples
	}
	this.myGraphPatternNotTriples = &GraphPatternNotTriples{
		Production: Production{
			Kind: KindGraphPatternNotTriples,
			Name: "GraphPatternNotTriples",
		},
	}

	this.myGraphPatternNotTriples.Parser = this.alt(
		this.GroupOrUnionGraphPattern(),
		this.OptionalGraphPattern(),
		this.MinusGraphPattern(),
		this.GraphGraphPattern(),
		this.ServiceGraphPattern(),
		this.Filter(),
		this.Bind(),
		this.InlineData())
	return this.myGraphPatternNotTriples
}

type GraphRef struct {
	Production
}

func (this *GraphRef) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphRef) String() string {
	return this.Name
}
func (this *GraphRef) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphRef() *GraphRef {
	if this.myGraphRef != nil {
		return this.myGraphRef
	}
	this.myGraphRef = &GraphRef{
		Production: Production{
			Kind: KindGraphRef,
			Name: "GraphRef",
		},
	}

	this.myGraphRef.Parser = this.seq(
		this.tok("GRAPH"),
		this.Iri())
	return this.myGraphRef
}

type GraphRefAll struct {
	Production
}

func (this *GraphRefAll) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphRefAll) String() string {
	return this.Name
}
func (this *GraphRefAll) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphRefAll() *GraphRefAll {
	if this.myGraphRefAll != nil {
		return this.myGraphRefAll
	}
	this.myGraphRefAll = &GraphRefAll{
		Production: Production{
			Kind: KindGraphRefAll,
			Name: "GraphRefAll",
		},
	}

	this.myGraphRefAll.Parser = this.alt(
		this.GraphRef(),
		this.tok("DEFAULT"),
		this.tok("NAMED"),
		this.tok("ALL"))
	return this.myGraphRefAll
}

type GraphTerm struct {
	Production
}

func (this *GraphTerm) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GraphTerm) String() string {
	return this.Name
}
func (this *GraphTerm) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GraphTerm() *GraphTerm {
	if this.myGraphTerm != nil {
		return this.myGraphTerm
	}
	this.myGraphTerm = &GraphTerm{
		Production: Production{
			Kind: KindGraphTerm,
			Name: "GraphTerm",
		},
	}

	this.myGraphTerm.Parser = this.alt(
		this.Iri(),
		this.RDFLiteral(),
		this.NumericLiteral(),
		this.BooleanLiteral(),
		this.BlankNode(),
		this.NIL())
	return this.myGraphTerm
}

type GroupClause struct {
	Production
}

func (this *GroupClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GroupClause) String() string {
	return this.Name
}
func (this *GroupClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GroupClause() *GroupClause {
	if this.myGroupClause != nil {
		return this.myGroupClause
	}
	this.myGroupClause = &GroupClause{
		Production: Production{
			Kind: KindGroupClause,
			Name: "GroupClause",
		},
	}

	this.myGroupClause.Parser = this.seq(
		this.tok("GROUP"),
		this.tok("BY"),
		this.rpt(
			this.GroupCondition()))
	return this.myGroupClause
}

type GroupCondition struct {
	Production
}

func (this *GroupCondition) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GroupCondition) String() string {
	return this.Name
}
func (this *GroupCondition) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GroupCondition() *GroupCondition {
	if this.myGroupCondition != nil {
		return this.myGroupCondition
	}
	this.myGroupCondition = &GroupCondition{
		Production: Production{
			Kind: KindGroupCondition,
			Name: "GroupCondition",
		},
	}

	this.myGroupCondition.Parser = this.alt(
		this.BuiltInCall(),
		this.FunctionCall(),
		this.seq(
			this.tok("("),
			this.Expression(),
			this.opt(
				this.grp(
					this.seq(
						this.tok("AS"),
						this.Var()))),
			this.tok(")")),
		this.Var())
	return this.myGroupCondition
}

type GroupGraphPattern struct {
	Production
}

func (this *GroupGraphPattern) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GroupGraphPattern) String() string {
	return this.Name
}
func (this *GroupGraphPattern) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GroupGraphPattern() *GroupGraphPattern {
	if this.myGroupGraphPattern != nil {
		return this.myGroupGraphPattern
	}
	this.myGroupGraphPattern = &GroupGraphPattern{
		Production: Production{
			Kind: KindGroupGraphPattern,
			Name: "GroupGraphPattern",
		},
	}

	this.myGroupGraphPattern.Parser = this.seq(
		this.tok("{"),
		this.grp(
			this.alt(
				this.SubSelect(),
				this.GroupGraphPatternSub())),
		this.tok("}"))
	return this.myGroupGraphPattern
}

type GroupGraphPatternSub struct {
	Production
}

func (this *GroupGraphPatternSub) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GroupGraphPatternSub) String() string {
	return this.Name
}
func (this *GroupGraphPatternSub) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GroupGraphPatternSub() *GroupGraphPatternSub {
	if this.myGroupGraphPatternSub != nil {
		return this.myGroupGraphPatternSub
	}
	this.myGroupGraphPatternSub = &GroupGraphPatternSub{
		Production: Production{
			Kind: KindGroupGraphPatternSub,
			Name: "GroupGraphPatternSub",
		},
	}

	this.myGroupGraphPatternSub.Parser = this.seq(
		this.opt(
			this.TriplesBlock()),
		this.opt(
			this.rpt(
				this.seq(
					this.GraphPatternNotTriples(),
					this.opt(
						this.tok(".")),
					this.opt(
						this.TriplesBlock())))))
	return this.myGroupGraphPatternSub
}

type GroupOrUnionGraphPattern struct {
	Production
}

func (this *GroupOrUnionGraphPattern) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *GroupOrUnionGraphPattern) String() string {
	return this.Name
}
func (this *GroupOrUnionGraphPattern) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) GroupOrUnionGraphPattern() *GroupOrUnionGraphPattern {
	if this.myGroupOrUnionGraphPattern != nil {
		return this.myGroupOrUnionGraphPattern
	}
	this.myGroupOrUnionGraphPattern = &GroupOrUnionGraphPattern{
		Production: Production{
			Kind: KindGroupOrUnionGraphPattern,
			Name: "GroupOrUnionGraphPattern",
		},
	}

	this.myGroupOrUnionGraphPattern.Parser = this.seq(
		this.GroupGraphPattern(),
		this.rpt(
			this.seq(
				this.tok("UNION"),
				this.GroupGraphPattern())))
	return this.myGroupOrUnionGraphPattern
}

type HavingClause struct {
	Production
}

func (this *HavingClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *HavingClause) String() string {
	return this.Name
}
func (this *HavingClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) HavingClause() *HavingClause {
	if this.myHavingClause != nil {
		return this.myHavingClause
	}
	this.myHavingClause = &HavingClause{
		Production: Production{
			Kind: KindHavingClause,
			Name: "HavingClause",
		},
	}

	this.myHavingClause.Parser = this.seq(
		this.tok("HAVING"),
		this.rpt(
			this.HavingCondition()))
	return this.myHavingClause
}

type HavingCondition struct {
	Production
}

func (this *HavingCondition) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *HavingCondition) String() string {
	return this.Name
}
func (this *HavingCondition) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) HavingCondition() *HavingCondition {
	if this.myHavingCondition != nil {
		return this.myHavingCondition
	}
	this.myHavingCondition = &HavingCondition{
		Production: Production{
			Kind: KindHavingCondition,
			Name: "HavingCondition",
		},
	}

	this.myHavingCondition.Parser = this.Constraint()
	return this.myHavingCondition
}

type InlineData struct {
	Production
}

func (this *InlineData) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *InlineData) String() string {
	return this.Name
}
func (this *InlineData) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) InlineData() *InlineData {
	if this.myInlineData != nil {
		return this.myInlineData
	}
	this.myInlineData = &InlineData{
		Production: Production{
			Kind: KindInlineData,
			Name: "InlineData",
		},
	}

	this.myInlineData.Parser = this.seq(
		this.tok("VALUES"),
		this.DataBlock())
	return this.myInlineData
}

type InlineDataFull struct {
	Production
}

func (this *InlineDataFull) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *InlineDataFull) String() string {
	return this.Name
}
func (this *InlineDataFull) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) InlineDataFull() *InlineDataFull {
	if this.myInlineDataFull != nil {
		return this.myInlineDataFull
	}
	this.myInlineDataFull = &InlineDataFull{
		Production: Production{
			Kind: KindInlineDataFull,
			Name: "InlineDataFull",
		},
	}

	this.myInlineDataFull.Parser = this.seq(
		this.grp(
			this.alt(
				this.NIL(),
				this.seq(
					this.tok("("),
					this.opt(
						this.rpt(
							this.Var())),
					this.tok(")")))),
		this.tok("{"),
		this.opt(
			this.rpt(
				this.alt(
					this.seq(
						this.tok("("),
						this.opt(
							this.rpt(
								this.DataBlockValue())),
						this.tok(")")),
					this.NIL()))),
		this.tok("}"))
	return this.myInlineDataFull
}

type InlineDataOneVar struct {
	Production
}

func (this *InlineDataOneVar) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *InlineDataOneVar) String() string {
	return this.Name
}
func (this *InlineDataOneVar) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) InlineDataOneVar() *InlineDataOneVar {
	if this.myInlineDataOneVar != nil {
		return this.myInlineDataOneVar
	}
	this.myInlineDataOneVar = &InlineDataOneVar{
		Production: Production{
			Kind: KindInlineDataOneVar,
			Name: "InlineDataOneVar",
		},
	}

	this.myInlineDataOneVar.Parser = this.seq(
		this.Var(),
		this.tok("{"),
		this.opt(
			this.rpt(
				this.DataBlockValue())),
		this.tok("}"))
	return this.myInlineDataOneVar
}

type InsertClause struct {
	Production
}

func (this *InsertClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *InsertClause) String() string {
	return this.Name
}
func (this *InsertClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) InsertClause() *InsertClause {
	if this.myInsertClause != nil {
		return this.myInsertClause
	}
	this.myInsertClause = &InsertClause{
		Production: Production{
			Kind: KindInsertClause,
			Name: "InsertClause",
		},
	}

	this.myInsertClause.Parser = this.seq(
		this.tok("INSERT"),
		this.QuadPattern())
	return this.myInsertClause
}

type InsertData struct {
	Production
}

func (this *InsertData) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *InsertData) String() string {
	return this.Name
}
func (this *InsertData) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) InsertData() *InsertData {
	if this.myInsertData != nil {
		return this.myInsertData
	}
	this.myInsertData = &InsertData{
		Production: Production{
			Kind: KindInsertData,
			Name: "InsertData",
		},
	}

	this.myInsertData.Parser = this.seq(
		this.tok("INSERT DATA"),
		this.QuadData())
	return this.myInsertData
}

type Iri struct {
	Production
}

func (this *Iri) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Iri) String() string {
	return this.Name
}
func (this *Iri) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Iri() *Iri {
	if this.myIri != nil {
		return this.myIri
	}
	this.myIri = &Iri{
		Production: Production{
			Kind: KindIri,
			Name: "Iri",
		},
	}

	this.myIri.Parser = this.alt(
		this.IRIREF(),
		this.PrefixedName())
	return this.myIri
}

type IriOrFunction struct {
	Production
}

func (this *IriOrFunction) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *IriOrFunction) String() string {
	return this.Name
}
func (this *IriOrFunction) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) IriOrFunction() *IriOrFunction {
	if this.myIriOrFunction != nil {
		return this.myIriOrFunction
	}
	this.myIriOrFunction = &IriOrFunction{
		Production: Production{
			Kind: KindIriOrFunction,
			Name: "IriOrFunction",
		},
	}

	this.myIriOrFunction.Parser = this.seq(
		this.Iri(),
		this.opt(
			this.ArgList()))
	return this.myIriOrFunction
}

type LimitClause struct {
	Production
}

func (this *LimitClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *LimitClause) String() string {
	return this.Name
}
func (this *LimitClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) LimitClause() *LimitClause {
	if this.myLimitClause != nil {
		return this.myLimitClause
	}
	this.myLimitClause = &LimitClause{
		Production: Production{
			Kind: KindLimitClause,
			Name: "LimitClause",
		},
	}

	this.myLimitClause.Parser = this.seq(
		this.tok("LIMIT"),
		this.INTEGER())
	return this.myLimitClause
}

type LimitOffsetClauses struct {
	Production
}

func (this *LimitOffsetClauses) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *LimitOffsetClauses) String() string {
	return this.Name
}
func (this *LimitOffsetClauses) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) LimitOffsetClauses() *LimitOffsetClauses {
	if this.myLimitOffsetClauses != nil {
		return this.myLimitOffsetClauses
	}
	this.myLimitOffsetClauses = &LimitOffsetClauses{
		Production: Production{
			Kind: KindLimitOffsetClauses,
			Name: "LimitOffsetClauses",
		},
	}

	this.myLimitOffsetClauses.Parser = this.alt(
		this.seq(
			this.LimitClause(),
			this.opt(
				this.OffsetClause())),
		this.seq(
			this.OffsetClause(),
			this.opt(
				this.LimitClause())))
	return this.myLimitOffsetClauses
}

type Load struct {
	Production
}

func (this *Load) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Load) String() string {
	return this.Name
}
func (this *Load) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Load() *Load {
	if this.myLoad != nil {
		return this.myLoad
	}
	this.myLoad = &Load{
		Production: Production{
			Kind: KindLoad,
			Name: "Load",
		},
	}

	this.myLoad.Parser = this.seq(
		this.tok("LOAD"),
		this.opt(
			this.tok("SILENT")),
		this.Iri(),
		this.opt(
			this.seq(
				this.tok("INTO"),
				this.GraphRef())))
	return this.myLoad
}

type MinusGraphPattern struct {
	Production
}

func (this *MinusGraphPattern) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *MinusGraphPattern) String() string {
	return this.Name
}
func (this *MinusGraphPattern) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) MinusGraphPattern() *MinusGraphPattern {
	if this.myMinusGraphPattern != nil {
		return this.myMinusGraphPattern
	}
	this.myMinusGraphPattern = &MinusGraphPattern{
		Production: Production{
			Kind: KindMinusGraphPattern,
			Name: "MinusGraphPattern",
		},
	}

	this.myMinusGraphPattern.Parser = this.seq(
		this.tok("MINUS"),
		this.GroupGraphPattern())
	return this.myMinusGraphPattern
}

type Modify struct {
	Production
}

func (this *Modify) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Modify) String() string {
	return this.Name
}
func (this *Modify) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Modify() *Modify {
	if this.myModify != nil {
		return this.myModify
	}
	this.myModify = &Modify{
		Production: Production{
			Kind: KindModify,
			Name: "Modify",
		},
	}

	this.myModify.Parser = this.seq(
		this.opt(
			this.seq(
				this.tok("WITH"),
				this.Iri())),
		this.grp(
			this.alt(
				this.seq(
					this.DeleteClause(),
					this.opt(
						this.InsertClause())),
				this.InsertClause())),
		this.rpt(
			this.UsingClause()),
		this.tok("WHERE"),
		this.GroupGraphPattern())
	return this.myModify
}

type Move struct {
	Production
}

func (this *Move) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Move) String() string {
	return this.Name
}
func (this *Move) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Move() *Move {
	if this.myMove != nil {
		return this.myMove
	}
	this.myMove = &Move{
		Production: Production{
			Kind: KindMove,
			Name: "Move",
		},
	}

	this.myMove.Parser = this.seq(
		this.tok("MOVE"),
		this.opt(
			this.tok("SILENT")),
		this.GraphOrDefault(),
		this.tok("TO"),
		this.GraphOrDefault())
	return this.myMove
}

type MultiplicativeExpression struct {
	Production
}

func (this *MultiplicativeExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *MultiplicativeExpression) String() string {
	return this.Name
}
func (this *MultiplicativeExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) MultiplicativeExpression() *MultiplicativeExpression {
	if this.myMultiplicativeExpression != nil {
		return this.myMultiplicativeExpression
	}
	this.myMultiplicativeExpression = &MultiplicativeExpression{
		Production: Production{
			Kind: KindMultiplicativeExpression,
			Name: "MultiplicativeExpression",
		},
	}

	this.myMultiplicativeExpression.Parser = this.seq(
		this.UnaryExpression(),
		this.rpt(
			this.alt(
				this.seq(
					this.tok("*"),
					this.UnaryExpression()),
				this.seq(
					this.tok("/"),
					this.UnaryExpression()))))
	return this.myMultiplicativeExpression
}

type NamedGraphClause struct {
	Production
}

func (this *NamedGraphClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NamedGraphClause) String() string {
	return this.Name
}
func (this *NamedGraphClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NamedGraphClause() *NamedGraphClause {
	if this.myNamedGraphClause != nil {
		return this.myNamedGraphClause
	}
	this.myNamedGraphClause = &NamedGraphClause{
		Production: Production{
			Kind: KindNamedGraphClause,
			Name: "NamedGraphClause",
		},
	}

	this.myNamedGraphClause.Parser = this.seq(
		this.tok("NAMED"),
		this.SourceSelector())
	return this.myNamedGraphClause
}

type NotExistsFunc struct {
	Production
}

func (this *NotExistsFunc) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NotExistsFunc) String() string {
	return this.Name
}
func (this *NotExistsFunc) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NotExistsFunc() *NotExistsFunc {
	if this.myNotExistsFunc != nil {
		return this.myNotExistsFunc
	}
	this.myNotExistsFunc = &NotExistsFunc{
		Production: Production{
			Kind: KindNotExistsFunc,
			Name: "NotExistsFunc",
		},
	}

	this.myNotExistsFunc.Parser = this.seq(
		this.tok("NOT"),
		this.tok("EXISTS"),
		this.GroupGraphPattern())
	return this.myNotExistsFunc
}

type NumericExpression struct {
	Production
}

func (this *NumericExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NumericExpression) String() string {
	return this.Name
}
func (this *NumericExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NumericExpression() *NumericExpression {
	if this.myNumericExpression != nil {
		return this.myNumericExpression
	}
	this.myNumericExpression = &NumericExpression{
		Production: Production{
			Kind: KindNumericExpression,
			Name: "NumericExpression",
		},
	}

	this.myNumericExpression.Parser = this.AdditiveExpression()
	return this.myNumericExpression
}

type NumericLiteral struct {
	Production
}

func (this *NumericLiteral) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NumericLiteral) String() string {
	return this.Name
}
func (this *NumericLiteral) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NumericLiteral() *NumericLiteral {
	if this.myNumericLiteral != nil {
		return this.myNumericLiteral
	}
	this.myNumericLiteral = &NumericLiteral{
		Production: Production{
			Kind: KindNumericLiteral,
			Name: "NumericLiteral",
		},
	}

	this.myNumericLiteral.Parser = this.alt(
		this.NumericLiteralUnsigned(),
		this.NumericLiteralPositive(),
		this.NumericLiteralNegative())
	return this.myNumericLiteral
}

type NumericLiteralNegative struct {
	Production
}

func (this *NumericLiteralNegative) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NumericLiteralNegative) String() string {
	return this.Name
}
func (this *NumericLiteralNegative) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NumericLiteralNegative() *NumericLiteralNegative {
	if this.myNumericLiteralNegative != nil {
		return this.myNumericLiteralNegative
	}
	this.myNumericLiteralNegative = &NumericLiteralNegative{
		Production: Production{
			Kind: KindNumericLiteralNegative,
			Name: "NumericLiteralNegative",
		},
	}

	this.myNumericLiteralNegative.Parser = this.alt(
		this.INTEGER_NEGATIVE(),
		this.DECIMAL_NEGATIVE(),
		this.DOUBLE_NEGATIVE())
	return this.myNumericLiteralNegative
}

type NumericLiteralPositive struct {
	Production
}

func (this *NumericLiteralPositive) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NumericLiteralPositive) String() string {
	return this.Name
}
func (this *NumericLiteralPositive) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NumericLiteralPositive() *NumericLiteralPositive {
	if this.myNumericLiteralPositive != nil {
		return this.myNumericLiteralPositive
	}
	this.myNumericLiteralPositive = &NumericLiteralPositive{
		Production: Production{
			Kind: KindNumericLiteralPositive,
			Name: "NumericLiteralPositive",
		},
	}

	this.myNumericLiteralPositive.Parser = this.alt(
		this.INTEGER_POSITIVE(),
		this.DECIMAL_POSITIVE(),
		this.DOUBLE_POSITIVE())
	return this.myNumericLiteralPositive
}

type NumericLiteralUnsigned struct {
	Production
}

func (this *NumericLiteralUnsigned) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NumericLiteralUnsigned) String() string {
	return this.Name
}
func (this *NumericLiteralUnsigned) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NumericLiteralUnsigned() *NumericLiteralUnsigned {
	if this.myNumericLiteralUnsigned != nil {
		return this.myNumericLiteralUnsigned
	}
	this.myNumericLiteralUnsigned = &NumericLiteralUnsigned{
		Production: Production{
			Kind: KindNumericLiteralUnsigned,
			Name: "NumericLiteralUnsigned",
		},
	}

	this.myNumericLiteralUnsigned.Parser = this.alt(
		this.INTEGER(),
		this.DECIMAL(),
		this.DOUBLE())
	return this.myNumericLiteralUnsigned
}

type Object struct {
	Production
}

func (this *Object) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Object) String() string {
	return this.Name
}
func (this *Object) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Object() *Object {
	if this.myObject != nil {
		return this.myObject
	}
	this.myObject = &Object{
		Production: Production{
			Kind: KindObject,
			Name: "Object",
		},
	}

	this.myObject.Parser = this.GraphNode()
	return this.myObject
}

type ObjectList struct {
	Production
}

func (this *ObjectList) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ObjectList) String() string {
	return this.Name
}
func (this *ObjectList) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ObjectList() *ObjectList {
	if this.myObjectList != nil {
		return this.myObjectList
	}
	this.myObjectList = &ObjectList{
		Production: Production{
			Kind: KindObjectList,
			Name: "ObjectList",
		},
	}

	this.myObjectList.Parser = this.seq(
		this.Object(),
		this.rpt(
			this.seq(
				this.tok(","),
				this.Object())))
	return this.myObjectList
}

type ObjectListPath struct {
	Production
}

func (this *ObjectListPath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ObjectListPath) String() string {
	return this.Name
}
func (this *ObjectListPath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ObjectListPath() *ObjectListPath {
	if this.myObjectListPath != nil {
		return this.myObjectListPath
	}
	this.myObjectListPath = &ObjectListPath{
		Production: Production{
			Kind: KindObjectListPath,
			Name: "ObjectListPath",
		},
	}

	this.myObjectListPath.Parser = this.seq(
		this.ObjectPath(),
		this.rpt(
			this.seq(
				this.tok(","),
				this.ObjectPath())))
	return this.myObjectListPath
}

type ObjectPath struct {
	Production
}

func (this *ObjectPath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ObjectPath) String() string {
	return this.Name
}
func (this *ObjectPath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ObjectPath() *ObjectPath {
	if this.myObjectPath != nil {
		return this.myObjectPath
	}
	this.myObjectPath = &ObjectPath{
		Production: Production{
			Kind: KindObjectPath,
			Name: "ObjectPath",
		},
	}

	this.myObjectPath.Parser = this.GraphNodePath()
	return this.myObjectPath
}

type OffsetClause struct {
	Production
}

func (this *OffsetClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *OffsetClause) String() string {
	return this.Name
}
func (this *OffsetClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) OffsetClause() *OffsetClause {
	if this.myOffsetClause != nil {
		return this.myOffsetClause
	}
	this.myOffsetClause = &OffsetClause{
		Production: Production{
			Kind: KindOffsetClause,
			Name: "OffsetClause",
		},
	}

	this.myOffsetClause.Parser = this.seq(
		this.tok("OFFSET"),
		this.INTEGER())
	return this.myOffsetClause
}

type OptionalGraphPattern struct {
	Production
}

func (this *OptionalGraphPattern) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *OptionalGraphPattern) String() string {
	return this.Name
}
func (this *OptionalGraphPattern) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) OptionalGraphPattern() *OptionalGraphPattern {
	if this.myOptionalGraphPattern != nil {
		return this.myOptionalGraphPattern
	}
	this.myOptionalGraphPattern = &OptionalGraphPattern{
		Production: Production{
			Kind: KindOptionalGraphPattern,
			Name: "OptionalGraphPattern",
		},
	}

	this.myOptionalGraphPattern.Parser = this.seq(
		this.tok("OPTIONAL"),
		this.GroupGraphPattern())
	return this.myOptionalGraphPattern
}

type OrderClause struct {
	Production
}

func (this *OrderClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *OrderClause) String() string {
	return this.Name
}
func (this *OrderClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) OrderClause() *OrderClause {
	if this.myOrderClause != nil {
		return this.myOrderClause
	}
	this.myOrderClause = &OrderClause{
		Production: Production{
			Kind: KindOrderClause,
			Name: "OrderClause",
		},
	}

	this.myOrderClause.Parser = this.seq(
		this.tok("ORDER"),
		this.tok("BY"),
		this.rpt(
			this.OrderCondition()))
	return this.myOrderClause
}

type OrderCondition struct {
	Production
}

func (this *OrderCondition) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *OrderCondition) String() string {
	return this.Name
}
func (this *OrderCondition) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) OrderCondition() *OrderCondition {
	if this.myOrderCondition != nil {
		return this.myOrderCondition
	}
	this.myOrderCondition = &OrderCondition{
		Production: Production{
			Kind: KindOrderCondition,
			Name: "OrderCondition",
		},
	}

	this.myOrderCondition.Parser = this.alt(
		this.grp(
			this.seq(
				this.grp(
					this.alt(
						this.tok("ASC"),
						this.tok("DESC"))),
				this.BrackettedExpression())),
		this.grp(
			this.alt(
				this.Constraint(),
				this.Var())))
	return this.myOrderCondition
}

type Path struct {
	Production
}

func (this *Path) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Path) String() string {
	return this.Name
}
func (this *Path) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Path() *Path {
	if this.myPath != nil {
		return this.myPath
	}
	this.myPath = &Path{
		Production: Production{
			Kind: KindPath,
			Name: "Path",
		},
	}

	this.myPath.Parser = this.PathAlternative()
	return this.myPath
}

type PathAlternative struct {
	Production
}

func (this *PathAlternative) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathAlternative) String() string {
	return this.Name
}
func (this *PathAlternative) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathAlternative() *PathAlternative {
	if this.myPathAlternative != nil {
		return this.myPathAlternative
	}
	this.myPathAlternative = &PathAlternative{
		Production: Production{
			Kind: KindPathAlternative,
			Name: "PathAlternative",
		},
	}

	this.myPathAlternative.Parser = this.seq(
		this.PathSequence(),
		this.rpt(
			this.seq(
				this.tok("|"),
				this.PathSequence())))
	return this.myPathAlternative
}

type PathElt struct {
	Production
}

func (this *PathElt) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathElt) String() string {
	return this.Name
}
func (this *PathElt) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathElt() *PathElt {
	if this.myPathElt != nil {
		return this.myPathElt
	}
	this.myPathElt = &PathElt{
		Production: Production{
			Kind: KindPathElt,
			Name: "PathElt",
		},
	}

	this.myPathElt.Parser = this.seq(
		this.PathPrimary(),
		this.opt(
			this.PathMod()))
	return this.myPathElt
}

type PathEltOrInverse struct {
	Production
}

func (this *PathEltOrInverse) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathEltOrInverse) String() string {
	return this.Name
}
func (this *PathEltOrInverse) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathEltOrInverse() *PathEltOrInverse {
	if this.myPathEltOrInverse != nil {
		return this.myPathEltOrInverse
	}
	this.myPathEltOrInverse = &PathEltOrInverse{
		Production: Production{
			Kind: KindPathEltOrInverse,
			Name: "PathEltOrInverse",
		},
	}

	this.myPathEltOrInverse.Parser = this.alt(
		this.PathElt(),
		this.seq(
			this.tok("^"),
			this.PathElt()))
	return this.myPathEltOrInverse
}

type PathMod struct {
	Production
}

func (this *PathMod) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathMod) String() string {
	return this.Name
}
func (this *PathMod) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathMod() *PathMod {
	if this.myPathMod != nil {
		return this.myPathMod
	}
	this.myPathMod = &PathMod{
		Production: Production{
			Kind: KindPathMod,
			Name: "PathMod",
		},
	}

	this.myPathMod.Parser = this.alt(
		this.tok("?"),
		this.tok("*"),
		this.tok("+"))
	return this.myPathMod
}

type PathNegatedPropertySet struct {
	Production
}

func (this *PathNegatedPropertySet) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathNegatedPropertySet) String() string {
	return this.Name
}
func (this *PathNegatedPropertySet) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathNegatedPropertySet() *PathNegatedPropertySet {
	if this.myPathNegatedPropertySet != nil {
		return this.myPathNegatedPropertySet
	}
	this.myPathNegatedPropertySet = &PathNegatedPropertySet{
		Production: Production{
			Kind: KindPathNegatedPropertySet,
			Name: "PathNegatedPropertySet",
		},
	}

	this.myPathNegatedPropertySet.Parser = this.alt(
		this.PathOneInPropertySet(),
		this.seq(
			this.tok("("),
			this.opt(
				this.seq(
					this.PathOneInPropertySet(),
					this.rpt(
						this.seq(
							this.tok("|"),
							this.PathOneInPropertySet())))),
			this.tok(")")))
	return this.myPathNegatedPropertySet
}

type PathOneInPropertySet struct {
	Production
}

func (this *PathOneInPropertySet) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathOneInPropertySet) String() string {
	return this.Name
}
func (this *PathOneInPropertySet) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathOneInPropertySet() *PathOneInPropertySet {
	if this.myPathOneInPropertySet != nil {
		return this.myPathOneInPropertySet
	}
	this.myPathOneInPropertySet = &PathOneInPropertySet{
		Production: Production{
			Kind: KindPathOneInPropertySet,
			Name: "PathOneInPropertySet",
		},
	}

	this.myPathOneInPropertySet.Parser = this.alt(
		this.Iri(),
		this.tok("a"),
		this.seq(
			this.tok("^"),
			this.grp(
				this.alt(
					this.Iri(),
					this.tok("a")))))
	return this.myPathOneInPropertySet
}

type PathPrimary struct {
	Production
}

func (this *PathPrimary) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathPrimary) String() string {
	return this.Name
}
func (this *PathPrimary) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathPrimary() *PathPrimary {
	if this.myPathPrimary != nil {
		return this.myPathPrimary
	}
	this.myPathPrimary = &PathPrimary{
		Production: Production{
			Kind: KindPathPrimary,
			Name: "PathPrimary",
		},
	}

	this.myPathPrimary.Parser = this.alt(
		this.Iri(),
		this.tok("a"),
		this.seq(
			this.tok("!"),
			this.PathNegatedPropertySet()),
		this.seq(
			this.tok("("),
			this.Path(),
			this.tok(")")))
	return this.myPathPrimary
}

type PathSequence struct {
	Production
}

func (this *PathSequence) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PathSequence) String() string {
	return this.Name
}
func (this *PathSequence) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PathSequence() *PathSequence {
	if this.myPathSequence != nil {
		return this.myPathSequence
	}
	this.myPathSequence = &PathSequence{
		Production: Production{
			Kind: KindPathSequence,
			Name: "PathSequence",
		},
	}

	this.myPathSequence.Parser = this.seq(
		this.PathEltOrInverse(),
		this.rpt(
			this.seq(
				this.tok("/"),
				this.PathEltOrInverse())))
	return this.myPathSequence
}

type PrefixDecl struct {
	Production
}

func (this *PrefixDecl) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PrefixDecl) String() string {
	return this.Name
}
func (this *PrefixDecl) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PrefixDecl() *PrefixDecl {
	if this.myPrefixDecl != nil {
		return this.myPrefixDecl
	}
	this.myPrefixDecl = &PrefixDecl{
		Production: Production{
			Kind: KindPrefixDecl,
			Name: "PrefixDecl",
		},
	}

	this.myPrefixDecl.Parser = this.seq(
		this.tok("PREFIX"),
		this.PNAME_NS(),
		this.IRIREF())
	return this.myPrefixDecl
}

type PrefixedName struct {
	Production
}

func (this *PrefixedName) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PrefixedName) String() string {
	return this.Name
}
func (this *PrefixedName) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PrefixedName() *PrefixedName {
	if this.myPrefixedName != nil {
		return this.myPrefixedName
	}
	this.myPrefixedName = &PrefixedName{
		Production: Production{
			Kind: KindPrefixedName,
			Name: "PrefixedName",
		},
	}

	this.myPrefixedName.Parser = this.alt(
		this.PNAME_LN(),
		this.PNAME_NS())
	return this.myPrefixedName
}

type PrimaryExpression struct {
	Production
}

func (this *PrimaryExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PrimaryExpression) String() string {
	return this.Name
}
func (this *PrimaryExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PrimaryExpression() *PrimaryExpression {
	if this.myPrimaryExpression != nil {
		return this.myPrimaryExpression
	}
	this.myPrimaryExpression = &PrimaryExpression{
		Production: Production{
			Kind: KindPrimaryExpression,
			Name: "PrimaryExpression",
		},
	}

	this.myPrimaryExpression.Parser = this.alt(
		this.BrackettedExpression(),
		this.BuiltInCall(),
		this.IriOrFunction(),
		this.RDFLiteral(),
		this.NumericLiteral(),
		this.BooleanLiteral(),
		this.Var())
	return this.myPrimaryExpression
}

type Prologue struct {
	Production
}

func (this *Prologue) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Prologue) String() string {
	return this.Name
}
func (this *Prologue) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Prologue() *Prologue {
	if this.myPrologue != nil {
		return this.myPrologue
	}
	this.myPrologue = &Prologue{
		Production: Production{
			Kind: KindPrologue,
			Name: "Prologue",
		},
	}

	this.myPrologue.Parser = this.opt(
		this.rpt(
			this.alt(
				this.BaseDecl(),
				this.PrefixDecl())))
	return this.myPrologue
}

type PropertyList struct {
	Production
}

func (this *PropertyList) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PropertyList) String() string {
	return this.Name
}
func (this *PropertyList) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PropertyList() *PropertyList {
	if this.myPropertyList != nil {
		return this.myPropertyList
	}
	this.myPropertyList = &PropertyList{
		Production: Production{
			Kind: KindPropertyList,
			Name: "PropertyList",
		},
	}

	this.myPropertyList.Parser = this.opt(
		this.PropertyListNotEmpty())
	return this.myPropertyList
}

type PropertyListNotEmpty struct {
	Production
}

func (this *PropertyListNotEmpty) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PropertyListNotEmpty) String() string {
	return this.Name
}
func (this *PropertyListNotEmpty) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PropertyListNotEmpty() *PropertyListNotEmpty {
	if this.myPropertyListNotEmpty != nil {
		return this.myPropertyListNotEmpty
	}
	this.myPropertyListNotEmpty = &PropertyListNotEmpty{
		Production: Production{
			Kind: KindPropertyListNotEmpty,
			Name: "PropertyListNotEmpty",
		},
	}

	this.myPropertyListNotEmpty.Parser = this.seq(
		this.Verb(),
		this.ObjectList(),
		this.rpt(
			this.seq(
				this.tok(";"),
				this.opt(
					this.seq(
						this.Verb(),
						this.ObjectList())))))
	return this.myPropertyListNotEmpty
}

type PropertyListPath struct {
	Production
}

func (this *PropertyListPath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PropertyListPath) String() string {
	return this.Name
}
func (this *PropertyListPath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PropertyListPath() *PropertyListPath {
	if this.myPropertyListPath != nil {
		return this.myPropertyListPath
	}
	this.myPropertyListPath = &PropertyListPath{
		Production: Production{
			Kind: KindPropertyListPath,
			Name: "PropertyListPath",
		},
	}

	this.myPropertyListPath.Parser = this.opt(
		this.PropertyListPathNotEmpty())
	return this.myPropertyListPath
}

type PropertyListPathNotEmpty struct {
	Production
}

func (this *PropertyListPathNotEmpty) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PropertyListPathNotEmpty) String() string {
	return this.Name
}
func (this *PropertyListPathNotEmpty) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PropertyListPathNotEmpty() *PropertyListPathNotEmpty {
	if this.myPropertyListPathNotEmpty != nil {
		return this.myPropertyListPathNotEmpty
	}
	this.myPropertyListPathNotEmpty = &PropertyListPathNotEmpty{
		Production: Production{
			Kind: KindPropertyListPathNotEmpty,
			Name: "PropertyListPathNotEmpty",
		},
	}

	this.myPropertyListPathNotEmpty.Parser = this.seq(
		this.grp(
			this.alt(
				this.VerbPath(),
				this.VerbSimple())),
		this.ObjectListPath(),
		this.rpt(
			this.seq(
				this.tok(";"),
				this.opt(
					this.seq(
						this.grp(
							this.alt(
								this.VerbPath(),
								this.VerbSimple())),
						this.ObjectList())))))
	return this.myPropertyListPathNotEmpty
}

type QuadData struct {
	Production
}

func (this *QuadData) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *QuadData) String() string {
	return this.Name
}
func (this *QuadData) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) QuadData() *QuadData {
	if this.myQuadData != nil {
		return this.myQuadData
	}
	this.myQuadData = &QuadData{
		Production: Production{
			Kind: KindQuadData,
			Name: "QuadData",
		},
	}

	this.myQuadData.Parser = this.seq(
		this.tok("{"),
		this.Quads(),
		this.tok("}"))
	return this.myQuadData
}

type QuadPattern struct {
	Production
}

func (this *QuadPattern) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *QuadPattern) String() string {
	return this.Name
}
func (this *QuadPattern) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) QuadPattern() *QuadPattern {
	if this.myQuadPattern != nil {
		return this.myQuadPattern
	}
	this.myQuadPattern = &QuadPattern{
		Production: Production{
			Kind: KindQuadPattern,
			Name: "QuadPattern",
		},
	}

	this.myQuadPattern.Parser = this.seq(
		this.tok("{"),
		this.Quads(),
		this.tok("}"))
	return this.myQuadPattern
}

type Quads struct {
	Production
}

func (this *Quads) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Quads) String() string {
	return this.Name
}
func (this *Quads) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Quads() *Quads {
	if this.myQuads != nil {
		return this.myQuads
	}
	this.myQuads = &Quads{
		Production: Production{
			Kind: KindQuads,
			Name: "Quads",
		},
	}

	this.myQuads.Parser = this.seq(
		this.opt(
			this.TriplesTemplate()),
		this.rpt(
			this.seq(
				this.QuadsNotTriples(),
				this.opt(
					this.tok(".")),
				this.opt(
					this.TriplesTemplate()))))
	return this.myQuads
}

type QuadsNotTriples struct {
	Production
}

func (this *QuadsNotTriples) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *QuadsNotTriples) String() string {
	return this.Name
}
func (this *QuadsNotTriples) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) QuadsNotTriples() *QuadsNotTriples {
	if this.myQuadsNotTriples != nil {
		return this.myQuadsNotTriples
	}
	this.myQuadsNotTriples = &QuadsNotTriples{
		Production: Production{
			Kind: KindQuadsNotTriples,
			Name: "QuadsNotTriples",
		},
	}

	this.myQuadsNotTriples.Parser = this.seq(
		this.tok("GRAPH"),
		this.VarOrIri(),
		this.tok("{"),
		this.opt(
			this.TriplesTemplate()),
		this.tok("}"))
	return this.myQuadsNotTriples
}

type Query struct {
	Production
}

func (this *Query) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Query) String() string {
	return this.Name
}
func (this *Query) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Query() *Query {
	if this.myQuery != nil {
		return this.myQuery
	}
	this.myQuery = &Query{
		Production: Production{
			Kind: KindQuery,
			Name: "Query",
		},
	}

	this.myQuery.Parser = this.seq(
		this.Prologue(),
		this.grp(
			this.alt(
				this.SelectQuery(),
				this.ConstructQuery(),
				this.DescribeQuery(),
				this.AskQuery())),
		this.ValuesClause())
	return this.myQuery
}

type QueryUnit struct {
	Production
}

func (this *QueryUnit) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *QueryUnit) String() string {
	return this.Name
}
func (this *QueryUnit) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) QueryUnit() *QueryUnit {
	if this.myQueryUnit != nil {
		return this.myQueryUnit
	}
	this.myQueryUnit = &QueryUnit{
		Production: Production{
			Kind: KindQueryUnit,
			Name: "QueryUnit",
		},
	}

	this.myQueryUnit.Parser = this.Query()
	return this.myQueryUnit
}

type RDFLiteral struct {
	Production
}

func (this *RDFLiteral) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *RDFLiteral) String() string {
	return this.Name
}
func (this *RDFLiteral) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) RDFLiteral() *RDFLiteral {
	if this.myRDFLiteral != nil {
		return this.myRDFLiteral
	}
	this.myRDFLiteral = &RDFLiteral{
		Production: Production{
			Kind: KindRDFLiteral,
			Name: "RDFLiteral",
		},
	}

	this.myRDFLiteral.Parser = this.seq(
		this.String(),
		this.opt(
			this.alt(
				this.LANGTAG(),
				this.grp(
					this.seq(
						this.tok("^^"),
						this.Iri())))))
	return this.myRDFLiteral
}

type RegexExpression struct {
	Production
}

func (this *RegexExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *RegexExpression) String() string {
	return this.Name
}
func (this *RegexExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) RegexExpression() *RegexExpression {
	if this.myRegexExpression != nil {
		return this.myRegexExpression
	}
	this.myRegexExpression = &RegexExpression{
		Production: Production{
			Kind: KindRegexExpression,
			Name: "RegexExpression",
		},
	}

	this.myRegexExpression.Parser = this.seq(
		this.tok("REGEX"),
		this.tok("("),
		this.Expression(),
		this.tok(","),
		this.Expression(),
		this.opt(
			this.seq(
				this.tok(","),
				this.Expression())),
		this.tok(")"))
	return this.myRegexExpression
}

type RelationalExpression struct {
	Production
}

func (this *RelationalExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *RelationalExpression) String() string {
	return this.Name
}
func (this *RelationalExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) RelationalExpression() *RelationalExpression {
	if this.myRelationalExpression != nil {
		return this.myRelationalExpression
	}
	this.myRelationalExpression = &RelationalExpression{
		Production: Production{
			Kind: KindRelationalExpression,
			Name: "RelationalExpression",
		},
	}

	this.myRelationalExpression.Parser = this.seq(
		this.NumericExpression(),
		this.opt(
			this.alt(
				this.seq(
					this.tok("="),
					this.NumericExpression()),
				this.seq(
					this.tok("!="),
					this.NumericExpression()),
				this.seq(
					this.tok("<"),
					this.NumericExpression()),
				this.seq(
					this.tok(">"),
					this.NumericExpression()),
				this.seq(
					this.tok("<="),
					this.NumericExpression()),
				this.seq(
					this.tok(">="),
					this.NumericExpression()),
				this.seq(
					this.tok("IN"),
					this.ExpressionList()),
				this.seq(
					this.tok("NOT"),
					this.tok("IN"),
					this.ExpressionList()))))
	return this.myRelationalExpression
}

type SelectClause struct {
	Production
}

func (this *SelectClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *SelectClause) String() string {
	return this.Name
}
func (this *SelectClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) SelectClause() *SelectClause {
	if this.mySelectClause != nil {
		return this.mySelectClause
	}
	this.mySelectClause = &SelectClause{
		Production: Production{
			Kind: KindSelectClause,
			Name: "SelectClause",
		},
	}

	this.mySelectClause.Parser = this.seq(
		this.tok("SELECT"),
		this.opt(
			this.alt(
				this.tok("DISTINCT"),
				this.tok("REDUCED"))),
		this.grp(
			this.alt(
				this.rpt(
					this.alt(
						this.Var(),
						this.grp(
							this.seq(
								this.tok("("),
								this.Expression(),
								this.tok("AS"),
								this.Var(),
								this.tok(")"))))),
				this.tok("*"))))
	return this.mySelectClause
}

type SelectQuery struct {
	Production
}

func (this *SelectQuery) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *SelectQuery) String() string {
	return this.Name
}
func (this *SelectQuery) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) SelectQuery() *SelectQuery {
	if this.mySelectQuery != nil {
		return this.mySelectQuery
	}
	this.mySelectQuery = &SelectQuery{
		Production: Production{
			Kind: KindSelectQuery,
			Name: "SelectQuery",
		},
	}

	this.mySelectQuery.Parser = this.seq(
		this.SelectClause(),
		this.opt(
			this.rpt(
				this.DatasetClause())),
		this.WhereClause(),
		this.SolutionModifier())
	return this.mySelectQuery
}

type ServiceGraphPattern struct {
	Production
}

func (this *ServiceGraphPattern) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ServiceGraphPattern) String() string {
	return this.Name
}
func (this *ServiceGraphPattern) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ServiceGraphPattern() *ServiceGraphPattern {
	if this.myServiceGraphPattern != nil {
		return this.myServiceGraphPattern
	}
	this.myServiceGraphPattern = &ServiceGraphPattern{
		Production: Production{
			Kind: KindServiceGraphPattern,
			Name: "ServiceGraphPattern",
		},
	}

	this.myServiceGraphPattern.Parser = this.seq(
		this.tok("SERVICE"),
		this.opt(
			this.tok("SILENT")),
		this.VarOrIri(),
		this.GroupGraphPattern())
	return this.myServiceGraphPattern
}

type SolutionModifier struct {
	Production
}

func (this *SolutionModifier) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *SolutionModifier) String() string {
	return this.Name
}
func (this *SolutionModifier) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) SolutionModifier() *SolutionModifier {
	if this.mySolutionModifier != nil {
		return this.mySolutionModifier
	}
	this.mySolutionModifier = &SolutionModifier{
		Production: Production{
			Kind: KindSolutionModifier,
			Name: "SolutionModifier",
		},
	}

	this.mySolutionModifier.Parser = this.seq(
		this.opt(
			this.GroupClause()),
		this.opt(
			this.HavingClause()),
		this.opt(
			this.OrderClause()),
		this.opt(
			this.LimitOffsetClauses()))
	return this.mySolutionModifier
}

type SourceSelector struct {
	Production
}

func (this *SourceSelector) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *SourceSelector) String() string {
	return this.Name
}
func (this *SourceSelector) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) SourceSelector() *SourceSelector {
	if this.mySourceSelector != nil {
		return this.mySourceSelector
	}
	this.mySourceSelector = &SourceSelector{
		Production: Production{
			Kind: KindSourceSelector,
			Name: "SourceSelector",
		},
	}

	this.mySourceSelector.Parser = this.Iri()
	return this.mySourceSelector
}

type StrReplaceExpression struct {
	Production
}

func (this *StrReplaceExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *StrReplaceExpression) String() string {
	return this.Name
}
func (this *StrReplaceExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) StrReplaceExpression() *StrReplaceExpression {
	if this.myStrReplaceExpression != nil {
		return this.myStrReplaceExpression
	}
	this.myStrReplaceExpression = &StrReplaceExpression{
		Production: Production{
			Kind: KindStrReplaceExpression,
			Name: "StrReplaceExpression",
		},
	}

	this.myStrReplaceExpression.Parser = this.seq(
		this.tok("REPLACE"),
		this.tok("("),
		this.Expression(),
		this.tok(","),
		this.Expression(),
		this.tok(","),
		this.Expression(),
		this.opt(
			this.seq(
				this.tok(","),
				this.Expression())),
		this.tok(")"))
	return this.myStrReplaceExpression
}

type String struct {
	Production
}

func (this *String) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *String) String() string {
	return this.Name
}
func (this *String) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) String() *String {
	if this.myString != nil {
		return this.myString
	}
	this.myString = &String{
		Production: Production{
			Kind: KindString,
			Name: "String",
		},
	}

	this.myString.Parser = this.alt(
		this.STRING_LITERAL1(),
		this.STRING_LITERAL2(),
		this.STRING_LITERAL_LONG1(),
		this.STRING_LITERAL_LONG2())
	return this.myString
}

type SubSelect struct {
	Production
}

func (this *SubSelect) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *SubSelect) String() string {
	return this.Name
}
func (this *SubSelect) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) SubSelect() *SubSelect {
	if this.mySubSelect != nil {
		return this.mySubSelect
	}
	this.mySubSelect = &SubSelect{
		Production: Production{
			Kind: KindSubSelect,
			Name: "SubSelect",
		},
	}

	this.mySubSelect.Parser = this.seq(
		this.SelectClause(),
		this.WhereClause(),
		this.SolutionModifier(),
		this.ValuesClause())
	return this.mySubSelect
}

type SubstringExpression struct {
	Production
}

func (this *SubstringExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *SubstringExpression) String() string {
	return this.Name
}
func (this *SubstringExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) SubstringExpression() *SubstringExpression {
	if this.mySubstringExpression != nil {
		return this.mySubstringExpression
	}
	this.mySubstringExpression = &SubstringExpression{
		Production: Production{
			Kind: KindSubstringExpression,
			Name: "SubstringExpression",
		},
	}

	this.mySubstringExpression.Parser = this.seq(
		this.tok("SUBSTR"),
		this.tok("("),
		this.Expression(),
		this.tok(","),
		this.Expression(),
		this.opt(
			this.seq(
				this.tok(","),
				this.Expression())),
		this.tok(")"))
	return this.mySubstringExpression
}

type TriplesBlock struct {
	Production
}

func (this *TriplesBlock) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *TriplesBlock) String() string {
	return this.Name
}
func (this *TriplesBlock) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) TriplesBlock() *TriplesBlock {
	if this.myTriplesBlock != nil {
		return this.myTriplesBlock
	}
	this.myTriplesBlock = &TriplesBlock{
		Production: Production{
			Kind: KindTriplesBlock,
			Name: "TriplesBlock",
		},
	}

	this.myTriplesBlock.Parser = this.seq(
		this.TriplesSameSubjectPath(),
		this.opt(
			this.seq(
				this.tok("."),
				this.opt(
					this.TriplesBlock()))))
	return this.myTriplesBlock
}

type TriplesNode struct {
	Production
}

func (this *TriplesNode) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *TriplesNode) String() string {
	return this.Name
}
func (this *TriplesNode) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) TriplesNode() *TriplesNode {
	if this.myTriplesNode != nil {
		return this.myTriplesNode
	}
	this.myTriplesNode = &TriplesNode{
		Production: Production{
			Kind: KindTriplesNode,
			Name: "TriplesNode",
		},
	}

	this.myTriplesNode.Parser = this.alt(
		this.Collection(),
		this.BlankNodePropertyList())
	return this.myTriplesNode
}

type TriplesNodePath struct {
	Production
}

func (this *TriplesNodePath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *TriplesNodePath) String() string {
	return this.Name
}
func (this *TriplesNodePath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) TriplesNodePath() *TriplesNodePath {
	if this.myTriplesNodePath != nil {
		return this.myTriplesNodePath
	}
	this.myTriplesNodePath = &TriplesNodePath{
		Production: Production{
			Kind: KindTriplesNodePath,
			Name: "TriplesNodePath",
		},
	}

	this.myTriplesNodePath.Parser = this.alt(
		this.CollectionPath(),
		this.BlankNodePropertyListPath())
	return this.myTriplesNodePath
}

type TriplesSameSubject struct {
	Production
}

func (this *TriplesSameSubject) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *TriplesSameSubject) String() string {
	return this.Name
}
func (this *TriplesSameSubject) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) TriplesSameSubject() *TriplesSameSubject {
	if this.myTriplesSameSubject != nil {
		return this.myTriplesSameSubject
	}
	this.myTriplesSameSubject = &TriplesSameSubject{
		Production: Production{
			Kind: KindTriplesSameSubject,
			Name: "TriplesSameSubject",
		},
	}

	this.myTriplesSameSubject.Parser = this.alt(
		this.seq(
			this.VarOrTerm(),
			this.PropertyListNotEmpty()),
		this.seq(
			this.TriplesNode(),
			this.PropertyList()))
	return this.myTriplesSameSubject
}

type TriplesSameSubjectPath struct {
	Production
}

func (this *TriplesSameSubjectPath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *TriplesSameSubjectPath) String() string {
	return this.Name
}
func (this *TriplesSameSubjectPath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) TriplesSameSubjectPath() *TriplesSameSubjectPath {
	if this.myTriplesSameSubjectPath != nil {
		return this.myTriplesSameSubjectPath
	}
	this.myTriplesSameSubjectPath = &TriplesSameSubjectPath{
		Production: Production{
			Kind: KindTriplesSameSubjectPath,
			Name: "TriplesSameSubjectPath",
		},
	}

	this.myTriplesSameSubjectPath.Parser = this.alt(
		this.seq(
			this.VarOrTerm(),
			this.PropertyListPathNotEmpty()),
		this.seq(
			this.TriplesNodePath(),
			this.PropertyListPath()))
	return this.myTriplesSameSubjectPath
}

type TriplesTemplate struct {
	Production
}

func (this *TriplesTemplate) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *TriplesTemplate) String() string {
	return this.Name
}
func (this *TriplesTemplate) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) TriplesTemplate() *TriplesTemplate {
	if this.myTriplesTemplate != nil {
		return this.myTriplesTemplate
	}
	this.myTriplesTemplate = &TriplesTemplate{
		Production: Production{
			Kind: KindTriplesTemplate,
			Name: "TriplesTemplate",
		},
	}

	this.myTriplesTemplate.Parser = this.seq(
		this.TriplesSameSubject(),
		this.opt(
			this.seq(
				this.tok("."),
				this.opt(
					this.TriplesTemplate()))))
	return this.myTriplesTemplate
}

type UnaryExpression struct {
	Production
}

func (this *UnaryExpression) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *UnaryExpression) String() string {
	return this.Name
}
func (this *UnaryExpression) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) UnaryExpression() *UnaryExpression {
	if this.myUnaryExpression != nil {
		return this.myUnaryExpression
	}
	this.myUnaryExpression = &UnaryExpression{
		Production: Production{
			Kind: KindUnaryExpression,
			Name: "UnaryExpression",
		},
	}

	this.myUnaryExpression.Parser = this.alt(
		this.seq(
			this.tok("!"),
			this.PrimaryExpression()),
		this.seq(
			this.tok("+"),
			this.PrimaryExpression()),
		this.seq(
			this.tok("-"),
			this.PrimaryExpression()),
		this.PrimaryExpression())
	return this.myUnaryExpression
}

type Update struct {
	Production
}

func (this *Update) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Update) String() string {
	return this.Name
}
func (this *Update) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Update() *Update {
	if this.myUpdate != nil {
		return this.myUpdate
	}
	this.myUpdate = &Update{
		Production: Production{
			Kind: KindUpdate,
			Name: "Update",
		},
	}

	this.myUpdate.Parser = this.seq(
		this.Prologue(),
		this.opt(
			this.seq(
				this.Update1(),
				this.opt(
					this.seq(
						this.tok(";"),
						this.Update())))))
	return this.myUpdate
}

type Update1 struct {
	Production
}

func (this *Update1) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Update1) String() string {
	return this.Name
}
func (this *Update1) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Update1() *Update1 {
	if this.myUpdate1 != nil {
		return this.myUpdate1
	}
	this.myUpdate1 = &Update1{
		Production: Production{
			Kind: KindUpdate1,
			Name: "Update1",
		},
	}

	this.myUpdate1.Parser = this.alt(
		this.Load(),
		this.Clear(),
		this.Drop(),
		this.Add(),
		this.Move(),
		this.Copy(),
		this.Create(),
		this.InsertData(),
		this.DeleteData(),
		this.DeleteWhere(),
		this.Modify())
	return this.myUpdate1
}

type UpdateUnit struct {
	Production
}

func (this *UpdateUnit) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *UpdateUnit) String() string {
	return this.Name
}
func (this *UpdateUnit) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) UpdateUnit() *UpdateUnit {
	if this.myUpdateUnit != nil {
		return this.myUpdateUnit
	}
	this.myUpdateUnit = &UpdateUnit{
		Production: Production{
			Kind: KindUpdateUnit,
			Name: "UpdateUnit",
		},
	}

	this.myUpdateUnit.Parser = this.Update()
	return this.myUpdateUnit
}

type UsingClause struct {
	Production
}

func (this *UsingClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *UsingClause) String() string {
	return this.Name
}
func (this *UsingClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) UsingClause() *UsingClause {
	if this.myUsingClause != nil {
		return this.myUsingClause
	}
	this.myUsingClause = &UsingClause{
		Production: Production{
			Kind: KindUsingClause,
			Name: "UsingClause",
		},
	}

	this.myUsingClause.Parser = this.seq(
		this.tok("USING"),
		this.grp(
			this.alt(
				this.Iri(),
				this.seq(
					this.tok("NAMED"),
					this.Iri()))))
	return this.myUsingClause
}

type ValueLogical struct {
	Production
}

func (this *ValueLogical) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ValueLogical) String() string {
	return this.Name
}
func (this *ValueLogical) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ValueLogical() *ValueLogical {
	if this.myValueLogical != nil {
		return this.myValueLogical
	}
	this.myValueLogical = &ValueLogical{
		Production: Production{
			Kind: KindValueLogical,
			Name: "ValueLogical",
		},
	}

	this.myValueLogical.Parser = this.RelationalExpression()
	return this.myValueLogical
}

type ValuesClause struct {
	Production
}

func (this *ValuesClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ValuesClause) String() string {
	return this.Name
}
func (this *ValuesClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ValuesClause() *ValuesClause {
	if this.myValuesClause != nil {
		return this.myValuesClause
	}
	this.myValuesClause = &ValuesClause{
		Production: Production{
			Kind: KindValuesClause,
			Name: "ValuesClause",
		},
	}

	this.myValuesClause.Parser = this.opt(
		this.seq(
			this.tok("VALUES"),
			this.DataBlock()))
	return this.myValuesClause
}

type Var struct {
	Production
}

func (this *Var) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Var) String() string {
	return this.Name
}
func (this *Var) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Var() *Var {
	if this.myVar != nil {
		return this.myVar
	}
	this.myVar = &Var{
		Production: Production{
			Kind: KindVar,
			Name: "Var",
		},
	}

	this.myVar.Parser = this.alt(
		this.VAR1(),
		this.VAR2())
	return this.myVar
}

type VarOrIri struct {
	Production
}

func (this *VarOrIri) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *VarOrIri) String() string {
	return this.Name
}
func (this *VarOrIri) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) VarOrIri() *VarOrIri {
	if this.myVarOrIri != nil {
		return this.myVarOrIri
	}
	this.myVarOrIri = &VarOrIri{
		Production: Production{
			Kind: KindVarOrIri,
			Name: "VarOrIri",
		},
	}

	this.myVarOrIri.Parser = this.alt(
		this.Var(),
		this.Iri())
	return this.myVarOrIri
}

type VarOrTerm struct {
	Production
}

func (this *VarOrTerm) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *VarOrTerm) String() string {
	return this.Name
}
func (this *VarOrTerm) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) VarOrTerm() *VarOrTerm {
	if this.myVarOrTerm != nil {
		return this.myVarOrTerm
	}
	this.myVarOrTerm = &VarOrTerm{
		Production: Production{
			Kind: KindVarOrTerm,
			Name: "VarOrTerm",
		},
	}

	this.myVarOrTerm.Parser = this.alt(
		this.Var(),
		this.GraphTerm())
	return this.myVarOrTerm
}

type Verb struct {
	Production
}

func (this *Verb) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *Verb) String() string {
	return this.Name
}
func (this *Verb) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) Verb() *Verb {
	if this.myVerb != nil {
		return this.myVerb
	}
	this.myVerb = &Verb{
		Production: Production{
			Kind: KindVerb,
			Name: "Verb",
		},
	}

	this.myVerb.Parser = this.alt(
		this.VarOrIri(),
		this.tok("a"))
	return this.myVerb
}

type VerbPath struct {
	Production
}

func (this *VerbPath) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *VerbPath) String() string {
	return this.Name
}
func (this *VerbPath) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) VerbPath() *VerbPath {
	if this.myVerbPath != nil {
		return this.myVerbPath
	}
	this.myVerbPath = &VerbPath{
		Production: Production{
			Kind: KindVerbPath,
			Name: "VerbPath",
		},
	}

	this.myVerbPath.Parser = this.Path()
	return this.myVerbPath
}

type VerbSimple struct {
	Production
}

func (this *VerbSimple) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *VerbSimple) String() string {
	return this.Name
}
func (this *VerbSimple) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) VerbSimple() *VerbSimple {
	if this.myVerbSimple != nil {
		return this.myVerbSimple
	}
	this.myVerbSimple = &VerbSimple{
		Production: Production{
			Kind: KindVerbSimple,
			Name: "VerbSimple",
		},
	}

	this.myVerbSimple.Parser = this.Var()
	return this.myVerbSimple
}

type WhereClause struct {
	Production
}

func (this *WhereClause) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *WhereClause) String() string {
	return this.Name
}
func (this *WhereClause) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) WhereClause() *WhereClause {
	if this.myWhereClause != nil {
		return this.myWhereClause
	}
	this.myWhereClause = &WhereClause{
		Production: Production{
			Kind: KindWhereClause,
			Name: "WhereClause",
		},
	}

	this.myWhereClause.Parser = this.seq(
		this.tok("WHERE"),
		this.GroupGraphPattern())
	return this.myWhereClause
}

type ANON struct {
	Production
}

func (this *ANON) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ANON) String() string {
	return this.Name
}
func (this *ANON) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ANON() *ANON {
	if this.myANON != nil {
		return this.myANON
	}
	this.myANON = &ANON{
		Production: Production{
			Kind: KindANON,
			Name: "ANON",
		},
	}

	this.myANON.Parser = this.seq(
		this.tok("["),
		this.opt(
			this.rpt(
				this.WS())),
		this.tok("]"))
	return this.myANON
}

type BLANK_NODE_LABEL struct {
	Production
}

func (this *BLANK_NODE_LABEL) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *BLANK_NODE_LABEL) String() string {
	return this.Name
}
func (this *BLANK_NODE_LABEL) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) BLANK_NODE_LABEL() *BLANK_NODE_LABEL {
	if this.myBLANK_NODE_LABEL != nil {
		return this.myBLANK_NODE_LABEL
	}
	this.myBLANK_NODE_LABEL = &BLANK_NODE_LABEL{
		Production: Production{
			Kind: KindBLANK_NODE_LABEL,
			Name: "BLANK_NODE_LABEL",
		},
	}

	this.myBLANK_NODE_LABEL.Parser = this.seq(
		this.tok("_:"),
		this.grp(
			this.alt(
				this.PN_CHARS_U(),
				this.chr("0", "9"))),
		this.opt(
			this.seq(
				this.rpt(
					this.alt(
						this.PN_CHARS(),
						this.tok("."))),
				this.PN_CHARS())))
	return this.myBLANK_NODE_LABEL
}

type DECIMAL struct {
	Production
}

func (this *DECIMAL) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DECIMAL) String() string {
	return this.Name
}
func (this *DECIMAL) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DECIMAL() *DECIMAL {
	if this.myDECIMAL != nil {
		return this.myDECIMAL
	}
	this.myDECIMAL = &DECIMAL{
		Production: Production{
			Kind: KindDECIMAL,
			Name: "DECIMAL",
		},
	}

	this.myDECIMAL.Parser = this.seq(
		this.rpt(
			this.chr("0", "9")),
		this.tok("."),
		this.rpt(
			this.chr("0", "9")))
	return this.myDECIMAL
}

type DECIMAL_NEGATIVE struct {
	Production
}

func (this *DECIMAL_NEGATIVE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DECIMAL_NEGATIVE) String() string {
	return this.Name
}
func (this *DECIMAL_NEGATIVE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DECIMAL_NEGATIVE() *DECIMAL_NEGATIVE {
	if this.myDECIMAL_NEGATIVE != nil {
		return this.myDECIMAL_NEGATIVE
	}
	this.myDECIMAL_NEGATIVE = &DECIMAL_NEGATIVE{
		Production: Production{
			Kind: KindDECIMAL_NEGATIVE,
			Name: "DECIMAL_NEGATIVE",
		},
	}

	this.myDECIMAL_NEGATIVE.Parser = this.seq(
		this.tok("-"),
		this.DECIMAL())
	return this.myDECIMAL_NEGATIVE
}

type DECIMAL_POSITIVE struct {
	Production
}

func (this *DECIMAL_POSITIVE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DECIMAL_POSITIVE) String() string {
	return this.Name
}
func (this *DECIMAL_POSITIVE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DECIMAL_POSITIVE() *DECIMAL_POSITIVE {
	if this.myDECIMAL_POSITIVE != nil {
		return this.myDECIMAL_POSITIVE
	}
	this.myDECIMAL_POSITIVE = &DECIMAL_POSITIVE{
		Production: Production{
			Kind: KindDECIMAL_POSITIVE,
			Name: "DECIMAL_POSITIVE",
		},
	}

	this.myDECIMAL_POSITIVE.Parser = this.seq(
		this.tok("+"),
		this.DECIMAL())
	return this.myDECIMAL_POSITIVE
}

type DOUBLE struct {
	Production
}

func (this *DOUBLE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DOUBLE) String() string {
	return this.Name
}
func (this *DOUBLE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DOUBLE() *DOUBLE {
	if this.myDOUBLE != nil {
		return this.myDOUBLE
	}
	this.myDOUBLE = &DOUBLE{
		Production: Production{
			Kind: KindDOUBLE,
			Name: "DOUBLE",
		},
	}

	this.myDOUBLE.Parser = this.alt(
		this.seq(
			this.rpt(
				this.chr("0", "9")),
			this.tok("."),
			this.rpt(
				this.chr("0", "9")),
			this.EXPONENT()),
		this.seq(
			this.tok("."),
			this.rpt(
				this.chr("0", "9")),
			this.EXPONENT()),
		this.seq(
			this.rpt(
				this.chr("0", "9")),
			this.EXPONENT()))
	return this.myDOUBLE
}

type DOUBLE_NEGATIVE struct {
	Production
}

func (this *DOUBLE_NEGATIVE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DOUBLE_NEGATIVE) String() string {
	return this.Name
}
func (this *DOUBLE_NEGATIVE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DOUBLE_NEGATIVE() *DOUBLE_NEGATIVE {
	if this.myDOUBLE_NEGATIVE != nil {
		return this.myDOUBLE_NEGATIVE
	}
	this.myDOUBLE_NEGATIVE = &DOUBLE_NEGATIVE{
		Production: Production{
			Kind: KindDOUBLE_NEGATIVE,
			Name: "DOUBLE_NEGATIVE",
		},
	}

	this.myDOUBLE_NEGATIVE.Parser = this.seq(
		this.tok("-"),
		this.DOUBLE())
	return this.myDOUBLE_NEGATIVE
}

type DOUBLE_POSITIVE struct {
	Production
}

func (this *DOUBLE_POSITIVE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *DOUBLE_POSITIVE) String() string {
	return this.Name
}
func (this *DOUBLE_POSITIVE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) DOUBLE_POSITIVE() *DOUBLE_POSITIVE {
	if this.myDOUBLE_POSITIVE != nil {
		return this.myDOUBLE_POSITIVE
	}
	this.myDOUBLE_POSITIVE = &DOUBLE_POSITIVE{
		Production: Production{
			Kind: KindDOUBLE_POSITIVE,
			Name: "DOUBLE_POSITIVE",
		},
	}

	this.myDOUBLE_POSITIVE.Parser = this.seq(
		this.tok("+"),
		this.DOUBLE())
	return this.myDOUBLE_POSITIVE
}

type ECHAR struct {
	Production
}

func (this *ECHAR) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *ECHAR) String() string {
	return this.Name
}
func (this *ECHAR) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) ECHAR() *ECHAR {
	if this.myECHAR != nil {
		return this.myECHAR
	}
	this.myECHAR = &ECHAR{
		Production: Production{
			Kind: KindECHAR,
			Name: "ECHAR",
		},
	}

	this.myECHAR.Parser = this.seq(
		this.tok("\\"),
		this.opt(
			this.tok("tbnrf\"'")))
	return this.myECHAR
}

type EXPONENT struct {
	Production
}

func (this *EXPONENT) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *EXPONENT) String() string {
	return this.Name
}
func (this *EXPONENT) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) EXPONENT() *EXPONENT {
	if this.myEXPONENT != nil {
		return this.myEXPONENT
	}
	this.myEXPONENT = &EXPONENT{
		Production: Production{
			Kind: KindEXPONENT,
			Name: "EXPONENT",
		},
	}

	this.myEXPONENT.Parser = this.seq(
		this.grp(
			this.alt(
				this.tok("e"),
				this.tok("E"))),
		this.opt(
			this.grp(
				this.alt(
					this.tok("+"),
					this.tok("-")))),
		this.rpt(
			this.chr("0", "9")))
	return this.myEXPONENT
}

type HEX struct {
	Production
}

func (this *HEX) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *HEX) String() string {
	return this.Name
}
func (this *HEX) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) HEX() *HEX {
	if this.myHEX != nil {
		return this.myHEX
	}
	this.myHEX = &HEX{
		Production: Production{
			Kind: KindHEX,
			Name: "HEX",
		},
	}

	this.myHEX.Parser = this.alt(
		this.chr("0", "9"),
		this.chr("A", "F"),
		this.chr("a", "f"))
	return this.myHEX
}

type INTEGER struct {
	Production
}

func (this *INTEGER) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *INTEGER) String() string {
	return this.Name
}
func (this *INTEGER) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) INTEGER() *INTEGER {
	if this.myINTEGER != nil {
		return this.myINTEGER
	}
	this.myINTEGER = &INTEGER{
		Production: Production{
			Kind: KindINTEGER,
			Name: "INTEGER",
		},
	}

	this.myINTEGER.Parser = this.rpt(
		this.chr("0", "9"))
	return this.myINTEGER
}

type INTEGER_NEGATIVE struct {
	Production
}

func (this *INTEGER_NEGATIVE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *INTEGER_NEGATIVE) String() string {
	return this.Name
}
func (this *INTEGER_NEGATIVE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) INTEGER_NEGATIVE() *INTEGER_NEGATIVE {
	if this.myINTEGER_NEGATIVE != nil {
		return this.myINTEGER_NEGATIVE
	}
	this.myINTEGER_NEGATIVE = &INTEGER_NEGATIVE{
		Production: Production{
			Kind: KindINTEGER_NEGATIVE,
			Name: "INTEGER_NEGATIVE",
		},
	}

	this.myINTEGER_NEGATIVE.Parser = this.seq(
		this.tok("-"),
		this.INTEGER())
	return this.myINTEGER_NEGATIVE
}

type INTEGER_POSITIVE struct {
	Production
}

func (this *INTEGER_POSITIVE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *INTEGER_POSITIVE) String() string {
	return this.Name
}
func (this *INTEGER_POSITIVE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) INTEGER_POSITIVE() *INTEGER_POSITIVE {
	if this.myINTEGER_POSITIVE != nil {
		return this.myINTEGER_POSITIVE
	}
	this.myINTEGER_POSITIVE = &INTEGER_POSITIVE{
		Production: Production{
			Kind: KindINTEGER_POSITIVE,
			Name: "INTEGER_POSITIVE",
		},
	}

	this.myINTEGER_POSITIVE.Parser = this.seq(
		this.tok("+"),
		this.INTEGER())
	return this.myINTEGER_POSITIVE
}

type IRIREF struct {
	Production
}

func (this *IRIREF) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *IRIREF) String() string {
	return this.Name
}
func (this *IRIREF) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) IRIREF() *IRIREF {
	if this.myIRIREF != nil {
		return this.myIRIREF
	}
	this.myIRIREF = &IRIREF{
		Production: Production{
			Kind: KindIRIREF,
			Name: "IRIREF",
		},
	}

	this.myIRIREF.Parser = this.seq(
		this.tok("<"),
		this.rpt(
			this.alt(
				this.chr("a", "z"),
				this.chr("A", "Z"),
				this.chr("0", "9"),
				this.tok(":"),
				this.tok("/"))),
		this.tok(">"))
	return this.myIRIREF
}

type LANGTAG struct {
	Production
}

func (this *LANGTAG) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *LANGTAG) String() string {
	return this.Name
}
func (this *LANGTAG) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) LANGTAG() *LANGTAG {
	if this.myLANGTAG != nil {
		return this.myLANGTAG
	}
	this.myLANGTAG = &LANGTAG{
		Production: Production{
			Kind: KindLANGTAG,
			Name: "LANGTAG",
		},
	}

	this.myLANGTAG.Parser = this.seq(
		this.tok("@"),
		this.rpt(
			this.alt(
				this.chr("a", "z"),
				this.chr("A", "Z"))),
		this.opt(
			this.seq(
				this.tok("-"),
				this.rpt(
					this.alt(
						this.chr("a", "z"),
						this.chr("A", "Z"),
						this.chr("0", "9"))))))
	return this.myLANGTAG
}

type NIL struct {
	Production
}

func (this *NIL) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *NIL) String() string {
	return this.Name
}
func (this *NIL) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) NIL() *NIL {
	if this.myNIL != nil {
		return this.myNIL
	}
	this.myNIL = &NIL{
		Production: Production{
			Kind: KindNIL,
			Name: "NIL",
		},
	}

	this.myNIL.Parser = this.seq(
		this.tok("("),
		this.opt(
			this.rpt(
				this.WS())),
		this.tok(")"))
	return this.myNIL
}

type PERCENT struct {
	Production
}

func (this *PERCENT) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PERCENT) String() string {
	return this.Name
}
func (this *PERCENT) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PERCENT() *PERCENT {
	if this.myPERCENT != nil {
		return this.myPERCENT
	}
	this.myPERCENT = &PERCENT{
		Production: Production{
			Kind: KindPERCENT,
			Name: "PERCENT",
		},
	}

	this.myPERCENT.Parser = this.seq(
		this.tok("%"),
		this.HEX(),
		this.HEX())
	return this.myPERCENT
}

type PLX struct {
	Production
}

func (this *PLX) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PLX) String() string {
	return this.Name
}
func (this *PLX) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PLX() *PLX {
	if this.myPLX != nil {
		return this.myPLX
	}
	this.myPLX = &PLX{
		Production: Production{
			Kind: KindPLX,
			Name: "PLX",
		},
	}

	this.myPLX.Parser = this.alt(
		this.PERCENT(),
		this.PN_LOCAL_ESC())
	return this.myPLX
}

type PNAME_LN struct {
	Production
}

func (this *PNAME_LN) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PNAME_LN) String() string {
	return this.Name
}
func (this *PNAME_LN) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PNAME_LN() *PNAME_LN {
	if this.myPNAME_LN != nil {
		return this.myPNAME_LN
	}
	this.myPNAME_LN = &PNAME_LN{
		Production: Production{
			Kind: KindPNAME_LN,
			Name: "PNAME_LN",
		},
	}

	this.myPNAME_LN.Parser = this.seq(
		this.PNAME_NS(),
		this.PN_LOCAL())
	return this.myPNAME_LN
}

type PNAME_NS struct {
	Production
}

func (this *PNAME_NS) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PNAME_NS) String() string {
	return this.Name
}
func (this *PNAME_NS) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PNAME_NS() *PNAME_NS {
	if this.myPNAME_NS != nil {
		return this.myPNAME_NS
	}
	this.myPNAME_NS = &PNAME_NS{
		Production: Production{
			Kind: KindPNAME_NS,
			Name: "PNAME_NS",
		},
	}

	this.myPNAME_NS.Parser = this.seq(
		this.opt(
			this.PN_PREFIX()),
		this.tok(":"))
	return this.myPNAME_NS
}

type PN_CHARS struct {
	Production
}

func (this *PN_CHARS) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PN_CHARS) String() string {
	return this.Name
}
func (this *PN_CHARS) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PN_CHARS() *PN_CHARS {
	if this.myPN_CHARS != nil {
		return this.myPN_CHARS
	}
	this.myPN_CHARS = &PN_CHARS{
		Production: Production{
			Kind: KindPN_CHARS,
			Name: "PN_CHARS",
		},
	}

	this.myPN_CHARS.Parser = this.alt(
		this.PN_CHARS_U(),
		this.tok("-"),
		this.chr("0", "9"),
		this.tok("\u00b7"),
		this.chr("\u0300", "\u036f"),
		this.chr("\u203f", "\u2040"))
	return this.myPN_CHARS
}

type PN_CHARS_BASE struct {
	Production
}

func (this *PN_CHARS_BASE) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PN_CHARS_BASE) String() string {
	return this.Name
}
func (this *PN_CHARS_BASE) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PN_CHARS_BASE() *PN_CHARS_BASE {
	if this.myPN_CHARS_BASE != nil {
		return this.myPN_CHARS_BASE
	}
	this.myPN_CHARS_BASE = &PN_CHARS_BASE{
		Production: Production{
			Kind: KindPN_CHARS_BASE,
			Name: "PN_CHARS_BASE",
		},
	}

	this.myPN_CHARS_BASE.Parser = this.alt(
		this.chr("A", "Z"),
		this.chr("a", "z"),
		this.chr("\u00c0", "\u00d6"),
		this.chr("\u00d8", "\u00f6"),
		this.chr("\u00f8", "\u02ff"),
		this.chr("\u0370", "\u037d"),
		this.chr("\u037f", "\u1fff"),
		this.chr("\u200c", "\u200d"),
		this.chr("\u2070", "\u218f"),
		this.chr("\u2c00", "\u2fef"),
		this.chr("\u3001", "\ud7ff"),
		this.chr("\uf900", "\ufdcf"),
		this.chr("\ufdf0", "\ufffd"),
		this.chr("\U00010000", "\U000effff"))
	return this.myPN_CHARS_BASE
}

type PN_CHARS_U struct {
	Production
}

func (this *PN_CHARS_U) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PN_CHARS_U) String() string {
	return this.Name
}
func (this *PN_CHARS_U) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PN_CHARS_U() *PN_CHARS_U {
	if this.myPN_CHARS_U != nil {
		return this.myPN_CHARS_U
	}
	this.myPN_CHARS_U = &PN_CHARS_U{
		Production: Production{
			Kind: KindPN_CHARS_U,
			Name: "PN_CHARS_U",
		},
	}

	this.myPN_CHARS_U.Parser = this.alt(
		this.PN_CHARS_BASE(),
		this.tok("_"))
	return this.myPN_CHARS_U
}

type PN_LOCAL struct {
	Production
}

func (this *PN_LOCAL) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PN_LOCAL) String() string {
	return this.Name
}
func (this *PN_LOCAL) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PN_LOCAL() *PN_LOCAL {
	if this.myPN_LOCAL != nil {
		return this.myPN_LOCAL
	}
	this.myPN_LOCAL = &PN_LOCAL{
		Production: Production{
			Kind: KindPN_LOCAL,
			Name: "PN_LOCAL",
		},
	}

	this.myPN_LOCAL.Parser = this.seq(
		this.grp(
			this.alt(
				this.PN_CHARS_U(),
				this.tok(":"),
				this.chr("0", "9"),
				this.PLX())),
		this.opt(
			this.seq(
				this.rpt(
					this.alt(
						this.PN_CHARS(),
						this.tok("."),
						this.tok(":"),
						this.PLX())),
				this.grp(
					this.alt(
						this.PN_CHARS(),
						this.tok(":"),
						this.PLX())))))
	return this.myPN_LOCAL
}

type PN_LOCAL_ESC struct {
	Production
}

func (this *PN_LOCAL_ESC) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PN_LOCAL_ESC) String() string {
	return this.Name
}
func (this *PN_LOCAL_ESC) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PN_LOCAL_ESC() *PN_LOCAL_ESC {
	if this.myPN_LOCAL_ESC != nil {
		return this.myPN_LOCAL_ESC
	}
	this.myPN_LOCAL_ESC = &PN_LOCAL_ESC{
		Production: Production{
			Kind: KindPN_LOCAL_ESC,
			Name: "PN_LOCAL_ESC",
		},
	}

	this.myPN_LOCAL_ESC.Parser = this.seq(
		this.tok("\\"),
		this.grp(
			this.alt(
				this.tok("_"),
				this.tok("~"),
				this.tok("."),
				this.tok("-"),
				this.tok("!"),
				this.tok("$"),
				this.tok("&"),
				this.tok("'"),
				this.tok("("),
				this.tok(")"),
				this.tok("*"),
				this.tok("+"),
				this.tok(","),
				this.tok(";"),
				this.tok("="),
				this.tok("/"),
				this.tok("?"),
				this.tok("#"),
				this.tok("@"),
				this.tok("%"))))
	return this.myPN_LOCAL_ESC
}

type PN_PREFIX struct {
	Production
}

func (this *PN_PREFIX) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *PN_PREFIX) String() string {
	return this.Name
}
func (this *PN_PREFIX) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) PN_PREFIX() *PN_PREFIX {
	if this.myPN_PREFIX != nil {
		return this.myPN_PREFIX
	}
	this.myPN_PREFIX = &PN_PREFIX{
		Production: Production{
			Kind: KindPN_PREFIX,
			Name: "PN_PREFIX",
		},
	}

	this.myPN_PREFIX.Parser = this.seq(
		this.PN_CHARS_BASE(),
		this.opt(
			this.seq(
				this.rpt(
					this.alt(
						this.PN_CHARS(),
						this.tok("."))),
				this.PN_CHARS())))
	return this.myPN_PREFIX
}

type STRING_LITERAL1 struct {
	Production
}

func (this *STRING_LITERAL1) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *STRING_LITERAL1) String() string {
	return this.Name
}
func (this *STRING_LITERAL1) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) STRING_LITERAL1() *STRING_LITERAL1 {
	if this.mySTRING_LITERAL1 != nil {
		return this.mySTRING_LITERAL1
	}
	this.mySTRING_LITERAL1 = &STRING_LITERAL1{
		Production: Production{
			Kind: KindSTRING_LITERAL1,
			Name: "STRING_LITERAL1",
		},
	}

	this.mySTRING_LITERAL1.Parser = this.seq(
		this.tok("'"),
		this.opt(
			this.rpt(
				this.alt(
					this.grp(
						this.alt(
							this.tok("^"),
							this.tok("'"),
							this.tok("\\"),
							this.tok("\n"),
							this.tok("\r"))),
					this.ECHAR()))),
		this.tok("'"))
	return this.mySTRING_LITERAL1
}

type STRING_LITERAL2 struct {
	Production
}

func (this *STRING_LITERAL2) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *STRING_LITERAL2) String() string {
	return this.Name
}
func (this *STRING_LITERAL2) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) STRING_LITERAL2() *STRING_LITERAL2 {
	if this.mySTRING_LITERAL2 != nil {
		return this.mySTRING_LITERAL2
	}
	this.mySTRING_LITERAL2 = &STRING_LITERAL2{
		Production: Production{
			Kind: KindSTRING_LITERAL2,
			Name: "STRING_LITERAL2",
		},
	}

	this.mySTRING_LITERAL2.Parser = this.seq(
		this.tok("\""),
		this.opt(
			this.rpt(
				this.alt(
					this.grp(
						this.alt(
							this.tok("^"),
							this.tok("\""),
							this.tok("\\"),
							this.tok("\n"),
							this.tok("\r"))),
					this.ECHAR()))),
		this.tok("\""))
	return this.mySTRING_LITERAL2
}

type STRING_LITERAL_LONG1 struct {
	Production
}

func (this *STRING_LITERAL_LONG1) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *STRING_LITERAL_LONG1) String() string {
	return this.Name
}
func (this *STRING_LITERAL_LONG1) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) STRING_LITERAL_LONG1() *STRING_LITERAL_LONG1 {
	if this.mySTRING_LITERAL_LONG1 != nil {
		return this.mySTRING_LITERAL_LONG1
	}
	this.mySTRING_LITERAL_LONG1 = &STRING_LITERAL_LONG1{
		Production: Production{
			Kind: KindSTRING_LITERAL_LONG1,
			Name: "STRING_LITERAL_LONG1",
		},
	}

	this.mySTRING_LITERAL_LONG1.Parser = this.seq(
		this.tok("'''"),
		this.rpt(
			this.seq(
				this.opt(
					this.alt(
						this.tok("'"),
						this.tok("''"))),
				this.grp(
					this.alt(
						this.opt(
							this.tok("^'\\")),
						this.ECHAR())))),
		this.tok("'''"))
	return this.mySTRING_LITERAL_LONG1
}

type STRING_LITERAL_LONG2 struct {
	Production
}

func (this *STRING_LITERAL_LONG2) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *STRING_LITERAL_LONG2) String() string {
	return this.Name
}
func (this *STRING_LITERAL_LONG2) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) STRING_LITERAL_LONG2() *STRING_LITERAL_LONG2 {
	if this.mySTRING_LITERAL_LONG2 != nil {
		return this.mySTRING_LITERAL_LONG2
	}
	this.mySTRING_LITERAL_LONG2 = &STRING_LITERAL_LONG2{
		Production: Production{
			Kind: KindSTRING_LITERAL_LONG2,
			Name: "STRING_LITERAL_LONG2",
		},
	}

	this.mySTRING_LITERAL_LONG2.Parser = this.seq(
		this.tok("\"\"\""),
		this.opt(
			this.rpt(
				this.seq(
					this.opt(
						this.alt(
							this.tok("\""),
							this.tok("\"\""))),
					this.grp(
						this.alt(
							this.opt(
								this.tok("^\"\\")),
							this.ECHAR()))))),
		this.tok("\"\"\""))
	return this.mySTRING_LITERAL_LONG2
}

type VAR1 struct {
	Production
}

func (this *VAR1) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *VAR1) String() string {
	return this.Name
}
func (this *VAR1) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) VAR1() *VAR1 {
	if this.myVAR1 != nil {
		return this.myVAR1
	}
	this.myVAR1 = &VAR1{
		Production: Production{
			Kind: KindVAR1,
			Name: "VAR1",
		},
	}

	this.myVAR1.Parser = this.seq(
		this.tok("?"),
		this.VARNAME())
	return this.myVAR1
}

type VAR2 struct {
	Production
}

func (this *VAR2) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *VAR2) String() string {
	return this.Name
}
func (this *VAR2) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) VAR2() *VAR2 {
	if this.myVAR2 != nil {
		return this.myVAR2
	}
	this.myVAR2 = &VAR2{
		Production: Production{
			Kind: KindVAR2,
			Name: "VAR2",
		},
	}

	this.myVAR2.Parser = this.seq(
		this.tok("$"),
		this.VARNAME())
	return this.myVAR2
}

type VARNAME struct {
	Production
}

func (this *VARNAME) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *VARNAME) String() string {
	return this.Name
}
func (this *VARNAME) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) VARNAME() *VARNAME {
	if this.myVARNAME != nil {
		return this.myVARNAME
	}
	this.myVARNAME = &VARNAME{
		Production: Production{
			Kind: KindVARNAME,
			Name: "VARNAME",
		},
	}

	this.myVARNAME.Parser = this.seq(
		this.grp(
			this.alt(
				this.PN_CHARS_U(),
				this.chr("0", "9"))),
		this.rpt(
			this.alt(
				this.PN_CHARS_U(),
				this.chr("0", "9"),
				this.tok("\u00b7"),
				this.chr("\u0300", "\u036f"),
				this.chr("\u203f", "\u2040"))))
	return this.myVARNAME
}

type WS struct {
	Production
}

func (this *WS) Parse(input string) (string, bool) {
	return this.Parser.Parse(input)
}
func (this *WS) String() string {
	return this.Name
}
func (this *WS) Syntax() string {
	return this.Name + " = " + this.Parser.String()
}
func (this *Sparql) WS() *WS {
	if this.myWS != nil {
		return this.myWS
	}
	this.myWS = &WS{
		Production: Production{
			Kind: KindWS,
			Name: "WS",
		},
	}

	this.myWS.Parser = this.alt(
		this.tok(" "),
		this.tok("\t"),
		this.tok("\r"),
		this.tok("\n"))
	return this.myWS
}
