package model

type (
	NodeMap interface {
		Each(func(Node))
		Get(name string) (node Node, ok bool)
	}

	nodemap map[string]*node
)

func (nodes nodemap) Each(f func(Node)) {
	for _, v := range nodes {
		f(v)
	}
}

func (nodes nodemap) Get(name string) (Node, bool) {
	node, ok := nodes[name]
	return node, ok
}

func (nodes nodemap) Contains(name string) bool {
	_, ok := nodes[name]
	return ok
}
