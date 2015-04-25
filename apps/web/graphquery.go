package main

import (
	"github.com/ReSc/godel/core/graph"
)

type (
	graphDataSource struct {
		graph *graph.Graph
	}

	Query interface {
		From(func(WhereClause)) PathQuery
		PathQuery
		Result() DataSource
	}

	PathQuery interface {
		In(func(WhereClause)) PathQuery
		Out(func(WhereClause)) PathQuery
	}

	WhereClause interface {
		Not() WhereClause
		And() WhereClause
		Or() WhereClause

		Tag(tag string) WhereClause
		TagMatches(pattern string) WhereClause

		Attr(name, value string) WhereClause
		AttrMatches(name, valuePattern string) WhereClause

		Select() SelectClause
	}

	SelectClause interface {
		Subject() AsClause
		Predicate() AsClause
		Object() AsClause
	}

	AsClause interface {
		As(name string)
	}
)

func TestQuery() {
	q := NewQuery()
	q.From(func(where WhereClause) {
		where.
			Tag("type").And().AttrMatches("name", "Service$").
			Select().Subject().As("serviceType")
	}).In(func(where WhereClause) {
		where.
			Tag("instance-of").
			Select().Subject().As("serviceInstance")
	}).Out(func(where WhereClause) {
		where.
			Tag("parameter").
			Select().Object().As("parameter")
	}).Out(func(where WhereClause) {
		where.
			Tag("value").
			Select().Object().As("value")
	})

	ds := q.Result()
	ds.Len()
}

func NewQuery() Query {
	return &query{}
}

type AstNode interface{}
type AstNodes struct{}
type query struct {
	graph *graph.Graph
	root  AstNode
}

type (
	Iterable interface {
		Range() Iterator
	}

	Iterator interface {
		Next() bool
		Node() *graph.Node
	}
)

func (q *query) From(where func(WhereClause)) PathQuery {
	node := &whereNode{q, &AstNodes{}}
	q.root = node
	where(node)
	return q
}

func (q *query) In(where func(WhereClause)) PathQuery {
	node := &whereNode{q, &AstNodes{}}
	q.root = node
	where(node)
	return q
}
func (q *query) Out(where func(WhereClause)) PathQuery {
	node := &whereNode{q, &AstNodes{}}
	q.root = node
	where(node)
	return q
}

func (q *query) Result() DataSource {
	return nil
}

type whereNode struct {
	q        *query
	children *AstNodes
}

func (w *whereNode) Children() *AstNodes                          { return w.children }
func (w *whereNode) And() WhereClause                             { return w }
func (w *whereNode) Or() WhereClause                              { return w }
func (w *whereNode) Not() WhereClause                             { return w }
func (w *whereNode) Attr(name, value string) WhereClause          { return w }
func (w *whereNode) Tag(tag string) WhereClause                   { return w }
func (w *whereNode) AttrMatches(name, pattern string) WhereClause { return w }
func (w *whereNode) TagMatches(tag string) WhereClause            { return w }
func (w *whereNode) Select() SelectClause                         { return &selectClause{w} }

type selectClause struct {
	w *whereNode
}

func (s *selectClause) Subject() AsClause   { return nil }
func (s *selectClause) Predicate() AsClause { return nil }
func (s *selectClause) Object() AsClause    { return nil }

func NewGraphDataSource(g *graph.Graph, q Query) DataSource {
	return newGraphDataSource(g, q)
}

func newGraphDataSource(g *graph.Graph, q Query) DataSource {
	return &graphDataSource{}
}
func (gds *graphDataSource) Id() int64                                 { return 0 }
func (gds *graphDataSource) Len() int                                  { return 0 }
func (gds *graphDataSource) Name() string                              { return "" }
func (gds *graphDataSource) Fields() Fields                            { return nil }
func (gds *graphDataSource) Range(start, count int) DataSourceIterator { return nil }
