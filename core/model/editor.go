package model

import (
	"errors"
	"strings"
	"sync"
)

type (
	Editor struct {
		sync.Mutex
		model *Model
	}
)

func NewEditor(m *Model) *Editor {
	if m == nil {
		panic("m is nil")
	}
	return &Editor{
		model: m,
	}
}

func (m *Editor) Get(path string) (Node, bool) {
	return m.model.Get(path)
}

func (m *Editor) SetId(n Node, id int64) {
	if id < 0 {
		panic("invalid id")
	}

	m.Lock()
	defer m.Unlock()

	n.self().id = id
}

func (m *Editor) SetName(n Node, name string) error {
	if !IsValidName(name) {
		return errors.New("Invalid node name: " + name)
	}

	m.Lock()
	defer m.Unlock()

	return m.rename(n.self(), name)
}

func (m *Editor) SetTag(n Node, tags ...string) error {
	msg := ""
	for _, tag := range tags {
		if !IsValidTag(tag) {
			msg += "Invalid tag '" + tag + "', "
		}
	}

	if len(msg) > 0 {
		return errors.New(msg[:len(msg)-len(", ")])
	}

	m.Lock()
	defer m.Unlock()

	node := n.self()
	for _, tag := range tags {
		if node.tags.set(tag) {
			node.raiseTagSet(tag)
		}
	}
	return nil
}

func (m *Editor) DelTag(n Node, tag string) bool {
	m.Lock()
	defer m.Unlock()

	node := n.self()
	if node.tags.del(tag) {
		node.raiseTagDel(tag)
		return true
	}
	return false
}

func (m *Editor) NewChild(n Node, name string) (Node, error) {
	if !IsValidName(name) {
		return nil, errors.New("invalid name")
	}

	m.Lock()
	defer m.Unlock()

	parent := n.self()
	if child, err := m.newChild(parent, name); err == nil {
		parent.raiseChildSet(child)
		return child, nil
	} else {
		return nil, err
	}
}

func (m *Editor) DelChild(n Node, name string) (Node, bool) {
	m.Lock()
	defer m.Unlock()

	parent := n.self()
	if child, deleted := m.delChild(parent, name); deleted {
		parent.raiseChildDel(child)
		return child, deleted
	}
	return nil, false
}

func (m *Editor) delChild(parent *node, name string) (*node, bool) {
	if child, present := parent.children[name]; present {
		delete(parent.children, name)
		child.parent = nil
		return child, true
	}
	return nil, false
}

func (m *Editor) newChild(parent *node, name string) (*node, error) {
	if _, ok := parent.children[name]; ok {
		return nil, errors.New("name exists: " + name)
	}
	child := new(node)
	id := m.model.nextId()
	child.ctor(parent, id, name)
	parent.children[name] = child
	return child, nil
}

func (m *Editor) rename(n *node, name string) error {
	child := n
	if child.parent != nil {
		parent := child.parent
		if _, nameExists := parent.children[name]; nameExists {
			return errors.New("name exists: " + name)
		}
		delete(parent.children, child.name)
		parent.children[name] = child
	}
	print("renaming ", child.name, " to ", name, "\n")
	child.name = name
	return nil
}

func popLeaf(path string) (string, string) {
	for {
		path = strings.TrimSuffix(path, Separator)
		i := strings.LastIndex(path, Separator)
		switch true {
		case i < 0:
			return path, ""
		case i == len(path)-len(Separator):
			continue
		case i >= 0:
			return path[i+len(Separator):], path[:i]
		}
	}
}

func popRoot(path string) (string, string) {
	for {
		path = strings.TrimPrefix(path, Separator)
		i := strings.Index(path, Separator)
		switch true {
		case i < 0:
			return path, ""
		case i == 0:
			continue
		case i > 0:
			return path[:i], path[i+len(Separator):]
		}
	}
}
