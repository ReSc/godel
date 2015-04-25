package main

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"strings"
)

func TestMatch() string {
	match := NewMatch()
	s := match.
		Node(Label("CLASS")).
		Return(Func("id"), Prop("name"), Prop("created"), Prop("modified")).
		Edge(Label("SPECIALIZATION_OF")).
		Node(Label("CLASS")).
		Return(Func("id"), Prop("name"), Prop("created"), Prop("modified")).
		String()
	return s
}

func main() {
	nodes := []struct {
		neoism.Props
	}{}

	q := &neoism.CypherQuery{
		Statement: TestMatch(),
		Result:    &nodes,
	}

	fmt.Println(q.Statement)

	db, err := neoism.Connect("http://localhost:7474/")
	if err != nil {
		panic(err)
	}

	err = db.Cypher(q)
	if err != nil {
		panic(err)
	}

	for _, n := range q.Columns() {
		fmt.Println(n)
	}

	for _, n := range nodes {
		for _, k := range q.Columns() {
			fmt.Printf("%v %v\n", k, n.Props[k])
		}
	}
}

type (
	Alias string
	Label string
	Prop  string
	Func  string
	Id    int64

	Value interface{}

	Expr interface {
		String() string
	}

	Match interface {
		Node(expr ...Expr) Match
		Edge(expr ...Expr) EdgeMatch
		Return(expr ...Expr) Match
		String() string
	}

	EdgeMatch interface {
		Match
		Any() Match
		Reverse() Match
	}
)

func (f Func) String() string {
	return string(f)
}

func NewMatch() Match {
	return &matchExpr{}
}

type matchExpr struct {
	nextNodeId int
	nextEdgeId int
	stack      []Expr
	returns    []*returnExpr
}

func (m *matchExpr) push(e Expr) {
	m.stack = append(m.stack, e)
}

func (m *matchExpr) peek() Expr {
	if len(m.stack) == 0 {
		panic("stack underflow")
	}
	return m.stack[len(m.stack)-1]
}

func (m *matchExpr) pop() Expr {
	if len(m.stack) == 0 {
		panic("stack underflow")
	}

	lastIndex := len(m.stack) - 1
	e := m.stack[lastIndex]
	m.stack[lastIndex] = nil
	m.stack = m.stack[:lastIndex]

	return e
}

func (m *matchExpr) getVarName(i int) string {
	switch expr := m.stack[i].(type) {
	case *edgeExpr:
		return expr.name
	case *nodeExpr:
		return expr.name
	default:
		return ""
	}
}
func (m *matchExpr) getWheres(i int) []*whereExpr {
	switch expr := m.stack[i].(type) {
	case *edgeExpr:
		return expr.wheres
	case *nodeExpr:
		return expr.wheres
	default:
		return nil
	}
}

func (m *matchExpr) String() string {
	s := m.matchClause()
	s += m.whereClause()
	s += m.returnClause()
	return s
}

func (m *matchExpr) matchClause() string {
	s := "MATCH\n"
	for i := range m.stack {
		s += fmt.Sprintf("     %v\n", m.stack[i])
	}
	return s
}

func (m *matchExpr) whereClause() string {
	where := "WHERE \n"
	addWhere := 0
	sep := " and\n"
	for i := range m.stack {
		name := m.getVarName(i)
		wheres := m.getWheres(i)

		if len(wheres) > 0 {
			addWhere += 1
			for i := range wheres {
				where += fmt.Sprintf("      %v.%v%v", name, wheres[i], sep)
			}
		}
	}
	if addWhere > 0 {
		return where[:len(where)-len(sep)] + "\n"
	} else {
		return ""
	}

}

func (m *matchExpr) returnClause() string {
	s := ""
	if len(m.returns) > 0 {
		s += "RETURN\n"
		for i := range m.returns {
			s += fmt.Sprintf("       %v", m.returns[i])
			if i < len(m.returns)-1 {
				s += ","
			}
			s += "\n"
		}
	}
	return s
}

func (m *matchExpr) newNodeName() string {
	m.nextNodeId++
	return fmt.Sprintf("n%02d", m.nextNodeId)
}

func (m *matchExpr) newEdgeName() string {
	m.nextEdgeId++
	return fmt.Sprintf("r%02d", m.nextEdgeId)
}

func (m *matchExpr) Node(e ...Expr) Match {
	ne := &nodeExpr{
		name: m.newNodeName(),
	}

	for i := range e {
		switch expr := e[i].(type) {
		case Id:
			ne.ids = append(ne.ids, expr)
		case Label:
			for _, l := range expr.Split() {
				ne.label = ne.label.Label(l.String())
			}
		case *propValueExpr:
			ne.props = append(ne.props, expr)
		case *whereExpr:
			ne.wheres = append(ne.wheres, expr)
		default:
			panic("Unsupported node expression")
		}
	}

	m.push(ne)
	return m
}

type nodeExpr struct {
	name string

	label  Label
	ids    []Id
	wheres []*whereExpr
	props  []*propValueExpr
}

func (n *nodeExpr) String() string {
	s := "("
	s += n.name
	s += ":" + n.label.String()

	if len(n.props) > 0 {
		s += " { "
		for i := range n.props {
			p := n.props[i]
			s += fmt.Sprintf("`%v`:\"%v\" ", p.prop, p.value)
		}
		s += "} "
	}
	s += ")"
	return s
}

func (m *matchExpr) Edge(e ...Expr) EdgeMatch {
	ee := &edgeExpr{
		name: m.newEdgeName(),
	}

	for i := range e {
		switch expr := e[i].(type) {
		case Id:
			ee.ids = append(ee.ids, expr)
		case Label:
			for _, l := range expr.Split() {
				ee.label = ee.label.Label(l.String())
			}
		case *propValueExpr:
			ee.props = append(ee.props, expr)
		case *whereExpr:
			ee.wheres = append(ee.wheres, expr)
		default:
			panic("Unsupported edge expression")
		}
	}

	m.push(ee)
	em := &edgeMatch{m, ee}
	em.Forward()
	return em
}

type edgeMatch struct {
	*matchExpr
	ee *edgeExpr
}

func (e edgeMatch) Any() Match {
	e.ee.dir = "-"
	return e.matchExpr
}

func (e edgeMatch) Reverse() Match {
	e.ee.dir = "<"
	return e.matchExpr
}

func (e edgeMatch) Forward() Match {
	e.ee.dir = ">"
	return e.matchExpr
}

type edgeExpr struct {
	name   string
	dir    string
	label  Label
	ids    []Id
	props  []*propValueExpr
	wheres []*whereExpr
}

func (e *edgeExpr) String() string {
	label := ""
	if len(e.label) > 0 {
		label = ":" + e.label.String()
	}

	props := ""
	if len(e.props) > 0 {
		props += " { "
		for i := range e.props {
			p := e.props[i]
			props += fmt.Sprintf("`%v`:\"%v\" ", p.prop, p.value)
		}
		props += "}"
	}

	s := fmt.Sprintf("-[%v%v%v]-", e.name, label, props)
	if e.dir == "<" {
		s = e.dir + s
	} else if e.dir == ">" {
		s = s + e.dir
	}
	return s
}

func (m *matchExpr) Return(e ...Expr) Match {
	name := ""
	switch curr := m.peek().(type) {
	case *nodeExpr:
		name = curr.name
	case *edgeExpr:
		name = curr.name
	}

	for i := range e {
		switch expr := e[i].(type) {
		case Prop:
			m.returns = append(m.returns, &returnExpr{
				varName: name,
				expr:    expr,
				alias:   Alias(name + "_" + expr.String()),
			})
		case Func:
			m.returns = append(m.returns, &returnExpr{
				varName: "",
				expr:    Func(expr.String() + "(" + name + ")"),
				alias:   Alias(name + "_" + expr.String()),
			})
		case *aliasExpr:
			m.returns = append(m.returns, &returnExpr{
				varName: name,
				expr:    expr.p,
				alias:   expr.a,
			})
		default:
			panic("Unsupported return expression")
		}
	}

	return m
}

type returnExpr struct {
	varName string
	expr    Expr
	alias   Alias
}

func (m *returnExpr) String() string {
	if len(m.varName) == 0 {
		return fmt.Sprintf("%v as %v", m.expr, m.alias)
	}
	return fmt.Sprintf("%v.%v as %v", m.varName, m.expr, m.alias)
}

func (a Alias) String() string {
	return string(a)
}

func (l Label) Label(s string) Label {
	if len(s) == 0 {
		return l
	}

	if len(l) == 0 {
		return Label(s)
	}

	return Label(string(l) + ":" + s)
}

func (l Label) Len() int {
	return strings.Count(string(l), ":") + 1
}

func (l Label) String() string {
	return string(l)
}

func (l Label) Split() []Label {
	parts := strings.Split(string(l), ":")
	labels := make([]Label, len(parts))
	for i := range parts {
		labels[i] = Label(parts[i])
	}
	return labels
}

func (p Prop) As(alias Alias) Expr {
	return &aliasExpr{p, alias}
}

type aliasExpr struct {
	p Prop
	a Alias
}

func (e *aliasExpr) String() string {
	return e.p.String() + " as " + e.a.String()
}

func (p Prop) Eq(v Value) Expr { return &propValueExpr{p, v} }

type propValueExpr struct {
	prop  Prop
	value Value
}

func (e *propValueExpr) String() string {
	return ""
}

func (p Prop) Gt(v Value) Expr    { return &whereExpr{operator: ">", prop: p, value: v} }
func (p Prop) Lt(v Value) Expr    { return &whereExpr{operator: "<", prop: p, value: v} }
func (p Prop) Gte(v Value) Expr   { return &whereExpr{operator: ">=", prop: p, value: v} }
func (p Prop) Lte(v Value) Expr   { return &whereExpr{operator: "<=", prop: p, value: v} }
func (p Prop) Match(v Value) Expr { return &whereExpr{operator: "=~", prop: p, value: v} }
func (p Prop) String() string     { return string(p) }

func (id Id) String() string { return fmt.Sprint(id) }

func (e whereExpr) String() string {
	return fmt.Sprintf("`%v` %v %#v", e.prop, e.operator, e.value)
}

type whereExpr struct {
	operator string
	prop     Prop
	value    Value
}
