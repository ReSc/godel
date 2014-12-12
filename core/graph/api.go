package graph

import (
	"encoding/xml"
	"strings"
)

func (this AttrMap) Set(key, value string) {
	if attr, ok := this[key]; ok {
		attr.Value = value
	} else {
		attr := NewAttr()
		attr.Name = key
		attr.Value = value
		this.Add(attr)
	}
}

func (g *Graph) Init() *Graph {
	g.Nodes = NewNodeMap()
	g.Edges = NewEdgeMap()
	g.nextEdgeId = 0
	g.nextNodeId = 0
	g.Root = g.NewNode("root")
	g.AddNode(g.Root)
	return g
}

func (g *Graph) NextEdgeId() int64 {
	id := g.nextEdgeId
	g.nextEdgeId += 1
	return id
}

func (g *Graph) NextNodeId() int64 {
	id := g.nextNodeId
	g.nextNodeId += 1
	return id
}

func (g *Graph) NewPredicate(name string) *Node {
	p := g.NewNode(name)
	p.Tags.Add("predicate")
	return p
}

func (g *Graph) NewNode(name string) *Node {
	n := NewNode()
	n.Id = g.NextNodeId()
	n.Attrs.Set("name", name)
	return n
}

func (g *Graph) AddEdge(sub, prd, obj *Node) *Edge {
	if sub == nil || prd == nil || obj == nil {
		panic("no nil nodes allowed")
	}
	e := NewEdge()
	e.Id = g.NextEdgeId()
	e.Sub = sub.Id
	e.Prd = prd.Id
	e.Obj = obj.Id
	g.AddNode(sub)
	g.AddNode(prd)
	g.AddNode(obj)
	sub.OutEdges.Add(e)
	prd.InEdges.Add(e)
	obj.InEdges.Add(e)
	g.Edges.Add(e)
	return e
}

func (g *Graph) DelEdge(e *Edge) {
	g.Edges.Del(e)
	if s, ok := g.Nodes[e.Sub]; ok {
		s.OutEdges.Del(e)
	}
	if p, ok := g.Nodes[e.Prd]; ok {
		p.InEdges.Del(e)
	}
	if o, ok := g.Nodes[e.Obj]; ok {
		o.InEdges.Del(e)
	}
}

func (g *Graph) AddNode(n *Node) bool {
	if g.Nodes.Add(n) && g.nextNodeId <= n.Id {
		g.nextNodeId = n.Id + 1
		return true
	}
	return false
}

func (g *Graph) DelNode(n *Node) bool {
	if g.Nodes.Del(n) {
		edges := n.InEdges
		n.InEdges = NewEdgeList()
		for _, e := range edges {
			g.DelEdge(e)
		}
		edges = n.OutEdges
		n.OutEdges = NewEdgeList()
		for _, e := range edges {
			g.DelEdge(e)
		}
		return true
	}
	return false
}

func (n *Node) Name() string {
	if a, ok := n.Attrs["name"]; ok {
		return a.Value
	}
	return ""
}

func (m NodeMap) MarshalXML(d *xml.Encoder, start xml.StartElement) error {
	keys := m.Keys()
	start.Attr = nil
	start.Name.Local = "Node"

	for _, key := range keys {
		err := d.EncodeElement(m[key], start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m EdgeMap) MarshalXML(d *xml.Encoder, start xml.StartElement) error {
	keys := m.Keys()
	start.Attr = nil
	start.Name.Local = strings.TrimSuffix(start.Name.Local, "Edges")
	for _, key := range keys {
		err := d.EncodeElement(m[key], start)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m EdgeList) MarshalXML(d *xml.Encoder, start xml.StartElement) error {
	start.Attr = nil
	start.Name.Local = strings.TrimSuffix(start.Name.Local, "Edges")
	for key := range m {
		err := d.EncodeElement(m[key], start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m AttrMap) MarshalXML(d *xml.Encoder, start xml.StartElement) error {
	keys := m.Keys()
	start.Attr = nil
	start.Name.Local = "Attr"
	for _, key := range keys {
		err := d.EncodeElement(m[key], start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m TagSet) MarshalXML(d *xml.Encoder, start xml.StartElement) error {
	start.Attr = nil
	start.Name.Local = "Tag"
	for key, _ := range m {
		err := d.EncodeElement(key, start)
		if err != nil {
			return err
		}
	}
	return nil
}
