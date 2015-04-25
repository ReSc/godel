package graph

import (
	"regexp"
)

type (
	Iterable interface {
		Iterator() Iterator
	}

	Iterator interface {
		Next() bool
		Edge() *Edge
	}

	EdgeClause interface {
		Subject() WhereClause
		Predicate() WhereClause
		Object() WhereClause

		Iterator() Iterator
	}

	WhereClause interface {
		Tag(tag string) WhereClause
		TagMatches(pattern string) WhereClause

		Attr(key, value string) WhereClause
		AttrMatches(key, pattern string) WhereClause

		Iterator() Iterator
	}

	graphIterable struct {
		g *Graph
	}

	edgeIterator struct {
		edges []*Edge
		edge  *Edge
	}
)

func test() {
	g := NewGraph()
	edges := GetIterable(g)
	types := NewEdgeClause(g, edges).Subject().Tag("type")
	for i := types.Iterator(); i.Next(); {

	}
}

func GetIterable(g *Graph) Iterable {
	return &graphIterable{g}
}

func (gi *graphIterable) Iterator() Iterator {
	return &edgeIterator{gi.g.Edges.Values(), nil}
}

func (i *edgeIterator) Next() bool {
	if len(i.edges) > 0 {
		i.edge = i.edges[0]
		i.edges = i.edges[1:]
		return true
	}

	i.edge = nil
	return false
}

func (i *edgeIterator) Edge() *Edge {
	return i.edge
}

func Where(i Iterable, filter func(*Edge) bool) Iterable {
	return &whereIterable{i, filter}
}

type whereIterable struct {
	inner  Iterable
	filter func(*Edge) bool
}

func (wi *whereIterable) Iterator() Iterator {
	return &whereIterator{wi.inner.Iterator(), wi.filter}
}

type whereIterator struct {
	inner  Iterator
	filter func(*Edge) bool
}

func (i *whereIterator) Next() bool {
	for i.inner.Next() {
		if i.filter(i.inner.Edge()) {
			return true
		}
	}

	return false
}

func (i *whereIterator) Edge() *Edge {
	return i.inner.Edge()
}

type edgeClause struct {
	g     *Graph
	inner Iterable
}

func NewEdgeClause(g *Graph, source Iterable) EdgeClause {
	return &edgeClause{g, source}
}

func (ec *edgeClause) Subject() WhereClause {
	pc := NewSubjectClause(ec.g, ec.inner)
	ec.inner = pc
	return pc
}

func (ec *edgeClause) Predicate() WhereClause {
	pc := NewPredicateClause(ec.g, ec.inner)
	ec.inner = pc
	return pc
}

func (ec *edgeClause) Object() WhereClause {
	pc := NewObjectClause(ec.g, ec.inner)
	ec.inner = pc
	return pc
}

func (ec *edgeClause) Iterator() Iterator {
	return ec.inner.Iterator()
}

type predicateClause struct {
	g     *Graph
	inner Iterable
}

func NewPredicateClause(g *Graph, source Iterable) WhereClause {
	return &predicateClause{g, source}
}

func (s *predicateClause) Iterator() Iterator {
	return s.inner.Iterator()
}

func (s *predicateClause) Tag(tag string) WhereClause {
	s.inner = Where(s.inner, func(e *Edge) bool {
		prd := s.g.Nodes[e.Prd]
		return prd.Tags.Contains(tag)
	})
	return s
}

func (s *predicateClause) Attr(key, value string) WhereClause {
	s.inner = Where(s.inner, func(e *Edge) bool {
		prd := s.g.Nodes[e.Prd]
		val, ok := prd.Attrs[key]
		return ok && val.Value == value
	})
	return s
}

func (s *predicateClause) TagMatches(pattern string) WhereClause {
	matcher := regexp.MustCompile(pattern)
	s.inner = Where(s.inner, func(e *Edge) bool {
		prd := s.g.Nodes[e.Prd]
		for tag := range prd.Tags {
			if matcher.MatchString(tag) {
				return true
			}
		}
		return false
	})
	return s
}

func (s *predicateClause) AttrMatches(key, pattern string) WhereClause {
	matcher := regexp.MustCompile(pattern)
	s.inner = Where(s.inner, func(e *Edge) bool {
		v, ok := s.g.Nodes[e.Prd].Attrs[key]
		return ok && matcher.MatchString(v.Value)
	})
	return s
}

type subjectClause struct {
	g     *Graph
	inner Iterable
}

func NewSubjectClause(g *Graph, source Iterable) WhereClause {
	return &subjectClause{g, source}
}

func (s *subjectClause) Iterator() Iterator {
	return s.inner.Iterator()
}

func (s *subjectClause) Tag(tag string) WhereClause {
	s.inner = Where(s.inner, func(e *Edge) bool {
		n := s.g.Nodes[e.Sub]
		return n.Tags.Contains(tag)
	})
	return s
}

func (s *subjectClause) Attr(key, value string) WhereClause {
	s.inner = Where(s.inner, func(e *Edge) bool {
		n := s.g.Nodes[e.Sub]
		val, ok := n.Attrs[key]
		return ok && val.Value == value
	})
	return s
}

func (s *subjectClause) TagMatches(pattern string) WhereClause {
	matcher := regexp.MustCompile(pattern)
	s.inner = Where(s.inner, func(e *Edge) bool {
		n := s.g.Nodes[e.Sub]
		for tag := range n.Tags {
			if matcher.MatchString(tag) {
				return true
			}
		}
		return false
	})
	return s
}

func (s *subjectClause) AttrMatches(key, pattern string) WhereClause {
	matcher := regexp.MustCompile(pattern)
	s.inner = Where(s.inner, func(e *Edge) bool {
		v, ok := s.g.Nodes[e.Sub].Attrs[key]
		return ok && matcher.MatchString(v.Value)
	})
	return s
}

type objectClause struct {
	g     *Graph
	inner Iterable
}

func NewObjectClause(g *Graph, source Iterable) WhereClause {
	return &subjectClause{g, source}
}

func (s *objectClause) Iterator() Iterator {
	return s.inner.Iterator()
}

func (s *objectClause) Tag(tag string) WhereClause {
	s.inner = Where(s.inner, func(e *Edge) bool {
		n := s.g.Nodes[e.Sub]
		return n.Tags.Contains(tag)
	})
	return s
}

func (s *objectClause) Attr(key, value string) WhereClause {
	s.inner = Where(s.inner, func(e *Edge) bool {
		n := s.g.Nodes[e.Sub]
		val, ok := n.Attrs[key]
		return ok && val.Value == value
	})
	return s
}

func (s *objectClause) TagMatches(pattern string) WhereClause {
	matcher := regexp.MustCompile(pattern)
	s.inner = Where(s.inner, func(e *Edge) bool {
		n := s.g.Nodes[e.Sub]
		for tag := range n.Tags {
			if matcher.MatchString(tag) {
				return true
			}
		}
		return false
	})
	return s
}

func (s *objectClause) AttrMatches(key, pattern string) WhereClause {
	matcher := regexp.MustCompile(pattern)
	s.inner = Where(s.inner, func(e *Edge) bool {
		v, ok := s.g.Nodes[e.Sub].Attrs[key]
		return ok && matcher.MatchString(v.Value)
	})
	return s
}
