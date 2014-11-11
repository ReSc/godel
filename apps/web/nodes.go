package main

import (
	"encoding/xml"
	"sort"
)

type (
	Node struct {
		XMLName    xml.Name `json:"-" xml:"node"`
		Id         int64    `json:"id" xml:"id,attr"`
		ParentId   int64    `json:"parentId" xml:"parentId,attr"`
		Provenance int64    `json:"prv" xml:"prv,attr"`
		Name       string   `json:"name" xml:"name,attr,omitempty"`
		Tags       Tags     `json:"tags,omitempty" xml:"tags"`
		Attrs      Attrs    `json:"attrs,omitempty" xml:"attr"`
		Rels       Rels     `json:"rels,omitempty" xml:"rel"`
	}

	Nodes []*Node
)

func (nodes Nodes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		return e.Encode(nodes[0])
	default:
		sortedNodes := make(Nodes, len(nodes))
		copy(sortedNodes, nodes)
		sort.Stable(NodesById(sortedNodes))

		for _, node := range sortedNodes {
			if err := e.Encode(node); err != nil {
				return err
			}
		}
		return nil
	}
}

func (nodes Nodes) Each(action func(*Node)) {
	for _, n := range nodes {
		action(n)
	}
}

func (nodes Nodes) FindAll(predicate func(*Node) bool) []*Node {
	nn := make([]*Node, 0, 8)
	for _, n := range nodes {
		if predicate(n) {
			nn = append(nn, n)
		}
	}
	return nn
}

func (nodes Nodes) FindFirst(predicate func(*Node) bool) (*Node, bool) {
	for _, n := range nodes {
		if predicate(n) {
			return n, true
		}
	}
	return nil, false
}

func (nodes Nodes) FindLast(predicate func(*Node) bool) (*Node, bool) {
	offset := len(nodes) - 1
	for i := range nodes {
		n := nodes[offset-i]
		if predicate(n) {
			return n, true
		}
	}
	return nil, false
}

type NodesById Nodes

func (nodes NodesById) Len() int           { return len(nodes) }
func (nodes NodesById) Less(i, j int) bool { return nodes[i].Id < nodes[j].Id }
func (nodes NodesById) Swap(i, j int)      { nodes[j], nodes[i] = nodes[i], nodes[j] }

type NodesByName Nodes

func (nodes NodesByName) Len() int           { return len(nodes) }
func (nodes NodesByName) Less(i, j int) bool { return nodes[i].Name < nodes[j].Name }
func (nodes NodesByName) Swap(i, j int)      { nodes[j], nodes[i] = nodes[i], nodes[j] }

func (nodes *Nodes) append(n *Node) {
	nn := *nodes
	nn = append(nn, n)
	*nodes = nn
}

func (nodes *Nodes) del(id int64) (*Node, bool) {
	nn := *nodes
	if i := nn.indexOf(id); i >= 0 {
		oldNode := nn[i]
		l := len(nn) - 1
		copy(nn[i:], nn[i+1:])
		nn[l], nn = nil, nn[:l]

		*nodes = nn

		return oldNode, true
	}
	return nil, false
}

func (nodes Nodes) contains(id int64) bool {
	return nodes.indexOf(id) >= 0
}

func (nodes Nodes) indexOf(id int64) int {
	for i := range nodes {
		if nodes[i].Id == id {
			return i
		}
	}
	return -1
}

func (nodes Nodes) get(id int64) (*Node, bool) {
	n, ok := nodes.FindFirst(func(n *Node) bool {
		return n.Id == id
	})
	return n, ok
}

func (nodes Nodes) findByParentId(id int64) []*Node {
	return nodes.FindAll(func(n *Node) bool {
		return n.ParentId == id
	})
}

func bootstrapNodes() Nodes {
	n := Nodes{
		&Node{
			Id:   100,
			Name: "Model",
			Attrs: Attrs{
				Attr{Key: "description", Value: "root node"},
				Attr{Key: "type", Value: "root"},
			},
			Rels: []*Rel{
				&Rel{
					Predicate: 0,
					Subject:   1,
					Object:    1,
				},
			},
		},
		&Node{
			Id:       110,
			ParentId: 100,
			Name:     "system",
		},
		&Node{
			Id:       111,
			ParentId: 110,
			Name:     "reflect",
		},
		&Node{
			Id:       112,
			ParentId: 110,
			Name:     "config",
		},
		&Node{
			Id:       120,
			ParentId: 100,
			Name:     "user",
		},
		&Node{
			Id:       130,
			ParentId: 100,
			Name:     "lib",
		},
	}

	sort.Stable(NodesByName(n))
	return n
}
