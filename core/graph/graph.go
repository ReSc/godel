package graph

import (
	xml "encoding/xml"
	sort "sort"
)

// Attr is a struct
type Attr struct {
	Name  string
	Value string
}

// NewAttr creates a new instance of Attr
func NewAttr() *Attr {
	return &Attr{}
}

// AttrMap is a map[string]*Attr
type AttrMap map[string]*Attr

// NewAttrMap creates a new instance of AttrMap
func NewAttrMap() AttrMap {
	return make(AttrMap)
}

func (m AttrMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewAttr()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Name key exists in the AttrMap
func (this AttrMap) Contains(item *Attr) bool {
	_, ok := this[item.Name]
	return ok
}

// ContainsKey returns true if the key exists in the AttrMap
func (this AttrMap) ContainsKey(key string) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the AttrMap
func (this AttrMap) Len() int {
	return len(this)
}

// Add adds the item using item.Name as the key
func (this AttrMap) Add(item *Attr) bool {
	if !this.Contains(item) {
		this[item.Name] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this AttrMap) DelKey(key string) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Name
func (this AttrMap) Del(item *Attr) bool {
	if this.ContainsKey(item.Name) {
		delete(this, item.Name)
		return true
	}
	return false
}

// Keys returns a new []string of all keys
func (this AttrMap) Keys() []string {
	keys := make([]string, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns a new sorted []string of all keys
func (this AttrMap) SortedKeys() []string {
	keys := this.Keys()
	sort.Strings(keys)
	return keys
}

// Each calls f for each map entry
func (this AttrMap) Each(f func(string, *Attr)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Attr containing all values
func (this AttrMap) Values() []*Attr {
	values := make([]*Attr, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

// Edge is a struct
type Edge struct {
	Id  int64
	Obj int64
	Prd int64
	Sub int64
}

// NewEdge creates a new instance of Edge
func NewEdge() *Edge {
	return &Edge{}
}

// EdgeList is a []*Edge
type EdgeList []*Edge

// NewEdgeList creates a new instance of EdgeList
func NewEdgeList() EdgeList {
	return make(EdgeList, 0, 4)
}

// NewEdgeListSized creates a new instance of EdgeList
func NewEdgeListSized(size, capacity int) EdgeList {
	return make(EdgeList, size, capacity)
}

func (this *EdgeList) Len() int {
	return len(*this)
}

func (this *EdgeList) Contains(value *Edge) bool {
	return this.IndexOf(value) >= 0
}

func (this *EdgeList) Del(value *Edge) bool {
	if i := this.IndexOf(value); i >= 0 {
		l := *this
		copy(l[i:], l[i+1:])
		var defaultValue *Edge
		l[len(l)-1] = defaultValue
		l = l[:len(l)-1]
		*this = l
		return true
	}
	return false
}

func (this *EdgeList) Add(value *Edge) bool {
	l := *this
	l = append(l, value)
	*this = l
	return true
}

func (this *EdgeList) IndexOf(value *Edge) int {
	items := *this
	for index, item := range items {
		if item == value {
			return index
		}
	}
	return -1
}

// Each iterates over all items
func (this *EdgeList) Each(f func(int, *Edge)) {
	items := *this
	for index, item := range items {
		f(index, item)
	}
}

// EdgeMap is a map[int64]*Edge
type EdgeMap map[int64]*Edge

// NewEdgeMap creates a new instance of EdgeMap
func NewEdgeMap() EdgeMap {
	return make(EdgeMap)
}

func (m EdgeMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewEdge()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Id key exists in the EdgeMap
func (this EdgeMap) Contains(item *Edge) bool {
	_, ok := this[item.Id]
	return ok
}

// ContainsKey returns true if the key exists in the EdgeMap
func (this EdgeMap) ContainsKey(key int64) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the EdgeMap
func (this EdgeMap) Len() int {
	return len(this)
}

// Add adds the item using item.Id as the key
func (this EdgeMap) Add(item *Edge) bool {
	if !this.Contains(item) {
		this[item.Id] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this EdgeMap) DelKey(key int64) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Id
func (this EdgeMap) Del(item *Edge) bool {
	if this.ContainsKey(item.Id) {
		delete(this, item.Id)
		return true
	}
	return false
}

// Keys returns a new []int64 of all keys
func (this EdgeMap) Keys() []int64 {
	keys := make([]int64, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// Each calls f for each map entry
func (this EdgeMap) Each(f func(int64, *Edge)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Edge containing all values
func (this EdgeMap) Values() []*Edge {
	values := make([]*Edge, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

// Graph is a struct
type Graph struct {
	Edges      EdgeMap
	Id         int64
	Name       string
	Nodes      NodeMap
	Root       *Node
	nextEdgeId int64
	nextNodeId int64
}

// NewGraph creates a new instance of Graph
func NewGraph() *Graph {
	return &Graph{
		Edges: NewEdgeMap(), Nodes: NewNodeMap(),
	}
}

// Node is a struct
type Node struct {
	Attrs    AttrMap
	Id       int64
	InEdges  EdgeList
	OutEdges EdgeList
	Tags     TagSet
}

// NewNode creates a new instance of Node
func NewNode() *Node {
	return &Node{
		Attrs: NewAttrMap(), InEdges: NewEdgeList(), OutEdges: NewEdgeList(), Tags: NewTagSet(),
	}
}

// NodeMap is a map[int64]*Node
type NodeMap map[int64]*Node

// NewNodeMap creates a new instance of NodeMap
func NewNodeMap() NodeMap {
	return make(NodeMap)
}

func (m NodeMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewNode()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Id key exists in the NodeMap
func (this NodeMap) Contains(item *Node) bool {
	_, ok := this[item.Id]
	return ok
}

// ContainsKey returns true if the key exists in the NodeMap
func (this NodeMap) ContainsKey(key int64) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the NodeMap
func (this NodeMap) Len() int {
	return len(this)
}

// Add adds the item using item.Id as the key
func (this NodeMap) Add(item *Node) bool {
	if !this.Contains(item) {
		this[item.Id] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this NodeMap) DelKey(key int64) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Id
func (this NodeMap) Del(item *Node) bool {
	if this.ContainsKey(item.Id) {
		delete(this, item.Id)
		return true
	}
	return false
}

// Keys returns a new []int64 of all keys
func (this NodeMap) Keys() []int64 {
	keys := make([]int64, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// Each calls f for each map entry
func (this NodeMap) Each(f func(int64, *Node)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Node containing all values
func (this NodeMap) Values() []*Node {
	values := make([]*Node, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

// TagSet is a map[string]bool
type TagSet map[string]bool

// NewTagSet creates a new instance of TagSet
func NewTagSet() TagSet {
	return make(TagSet)
}

func (this TagSet) Contains(key string) bool {
	_, ok := this[key]
	return ok
}

func (this TagSet) Len() int {
	return len(this)
}

func (this TagSet) Add(key string) bool {
	if !this.Contains(key) {
		this[key] = true
		return true
	}
	return false
}

func (this TagSet) Del(key string) bool {
	if this.Contains(key) {
		delete(this, key)
		return true
	}
	return false
}

// Each iterates over all items
func (this TagSet) Each(f func(string)) {
	for key, _ := range this {
		f(key)
	}
}
