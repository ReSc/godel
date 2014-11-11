package model

type (
	Model struct {
		nId int64
		node
	}
)

func New(name string) *Model {
	m := new(Model)
	m.nId = 1
	m.node.ctor(nil, 0, name)
	return m
}

func (m *Model) nextId() int64 {
	id := m.nId
	m.nId = id + 1
	return id
}
func (m *Model) Get(path string) (Node, bool) {
	node, ok := m.find(path)
	return node, ok
}
