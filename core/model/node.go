package model

import (
	"strings"
)

type (
	Node interface {
		Ident
		Observable
		FullName() string
		Parent() Node
		Children() NodeMap
		Tags() Tags
		Attributes() Attributes
		self() *node
	}

	node struct {
		ident
		subject
		parent     *node
		tags       tags
		children   nodemap
		attributes attributes
	}
)

func (n *node) ctor(parent *node, id int64, name string) *node {
	n.ident.ctor(id, name)
	n.subject.ctor(n)
	n.parent = parent
	n.children = make(nodemap)
	n.attributes = make(attributes)
	n.tags = make(tags)
	return n
}

func (node *node) self() *node {
	return node
}

func (node *node) Id() int64 {
	return node.ident.id
}

func (node *node) Name() string {
	return node.name
}

func (node *node) FullName() string {
	prefix := ""
	if node.parent != nil {
		prefix = node.parent.FullName()
	}
	return prefix + Separator + node.name
}

func (node *node) Parent() Node {
	return node.parent
}

func (node *node) Children() NodeMap {
	return node.children
}

func (node *node) Attributes() Attributes {
	return node.attributes
}

func (node *node) Tags() Tags {
	return node.tags
}

func (node *node) root() *node {
	for node.parent != nil {
		node = node.parent
	}
	return node
}

func (n *node) find(path string) (*node, bool) {
	name := ""
	if strings.HasPrefix(path, Separator) {
		n = n.root()
		name, path = popRoot(path)
		if n.name != name {
			return nil, false
		}
	}

	for len(path) > 0 {
		name, path = popRoot(path)
		if child, ok := n.children[name]; ok {
			n = child
		} else {
			n = nil
			break
		}
	}
	return n, n != nil
}
