package reflect2

import (
	xml "encoding/xml"
	sort "sort"
)

type Accepter interface {
	Accept(v Visitor)
}

// Attr is a struct
type Attr struct {
	Name  string `xml:"name,attr" `
	Value string `xml:"value,attr" `
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

// Field is a struct
type Field struct {
	Attrs    AttrMap `xml:"attr" `
	DataType string  `xml:"data-type,attr" `
	Name     string  `xml:"name,attr" `
	Type     *Type   `json:"-" xml:"-" `
	Value    string  `xml:"value,attr" `
}

// NewField creates a new instance of Field
func NewField() *Field {
	return &Field{
		Attrs: NewAttrMap(),
	}
}

func (this *Field) Accept(v Visitor) {
	this.doAccept(v)
}

// FieldMap is a map[string]*Field
type FieldMap map[string]*Field

// NewFieldMap creates a new instance of FieldMap
func NewFieldMap() FieldMap {
	return make(FieldMap)
}

func (m FieldMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewField()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Name key exists in the FieldMap
func (this FieldMap) Contains(item *Field) bool {
	_, ok := this[item.Name]
	return ok
}

// ContainsKey returns true if the key exists in the FieldMap
func (this FieldMap) ContainsKey(key string) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the FieldMap
func (this FieldMap) Len() int {
	return len(this)
}

// Add adds the item using item.Name as the key
func (this FieldMap) Add(item *Field) bool {
	if !this.Contains(item) {
		this[item.Name] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this FieldMap) DelKey(key string) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Name
func (this FieldMap) Del(item *Field) bool {
	if this.ContainsKey(item.Name) {
		delete(this, item.Name)
		return true
	}
	return false
}

// Keys returns a new []string of all keys
func (this FieldMap) Keys() []string {
	keys := make([]string, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns a new sorted []string of all keys
func (this FieldMap) SortedKeys() []string {
	keys := this.Keys()
	sort.Strings(keys)
	return keys
}

// Each calls f for each map entry
func (this FieldMap) Each(f func(string, *Field)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Field containing all values
func (this FieldMap) Values() []*Field {
	values := make([]*Field, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

// Import is a struct
type Import struct {
	Alias   string   `xml:"alias,attr" `
	Package *Package `json:"-" xml:"-" `
	Path    string   `xml:"path,attr" `
}

// NewImport creates a new instance of Import
func NewImport() *Import {
	return &Import{}
}

// ImportMap is a map[string]*Import
type ImportMap map[string]*Import

// NewImportMap creates a new instance of ImportMap
func NewImportMap() ImportMap {
	return make(ImportMap)
}

func (m ImportMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewImport()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Path key exists in the ImportMap
func (this ImportMap) Contains(item *Import) bool {
	_, ok := this[item.Path]
	return ok
}

// ContainsKey returns true if the key exists in the ImportMap
func (this ImportMap) ContainsKey(key string) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the ImportMap
func (this ImportMap) Len() int {
	return len(this)
}

// Add adds the item using item.Path as the key
func (this ImportMap) Add(item *Import) bool {
	if !this.Contains(item) {
		this[item.Path] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this ImportMap) DelKey(key string) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Path
func (this ImportMap) Del(item *Import) bool {
	if this.ContainsKey(item.Path) {
		delete(this, item.Path)
		return true
	}
	return false
}

// Keys returns a new []string of all keys
func (this ImportMap) Keys() []string {
	keys := make([]string, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns a new sorted []string of all keys
func (this ImportMap) SortedKeys() []string {
	keys := this.Keys()
	sort.Strings(keys)
	return keys
}

// Each calls f for each map entry
func (this ImportMap) Each(f func(string, *Import)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Import containing all values
func (this ImportMap) Values() []*Import {
	values := make([]*Import, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

// MetaType is a struct
type MetaType struct {
	ElementName     string
	ElementType     string
	ElementTypeName string
	IsContainer     bool
	KeyName         string
	KeyType         string
	Name            string
	Type            string
}

// NewMetaType creates a new instance of MetaType
func NewMetaType() *MetaType {
	return &MetaType{}
}

// Method is a struct
type Method struct {
	Name    string    `xml:"name,attr" `
	Params  *ParamMap `xml:"param" `
	Returns *ParamMap `xml:"return" `
}

// NewMethod creates a new instance of Method
func NewMethod() *Method {
	return &Method{
		Params: NewParamMap(), Returns: NewParamMap(),
	}
}

func (this *Method) Accept(v Visitor) {
	this.doAccept(v)
}

// MethodMap is a map[string]*Method
type MethodMap map[string]*Method

// NewMethodMap creates a new instance of MethodMap
func NewMethodMap() MethodMap {
	return make(MethodMap)
}

func (m MethodMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewMethod()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Name key exists in the MethodMap
func (this MethodMap) Contains(item *Method) bool {
	_, ok := this[item.Name]
	return ok
}

// ContainsKey returns true if the key exists in the MethodMap
func (this MethodMap) ContainsKey(key string) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the MethodMap
func (this MethodMap) Len() int {
	return len(this)
}

// Add adds the item using item.Name as the key
func (this MethodMap) Add(item *Method) bool {
	if !this.Contains(item) {
		this[item.Name] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this MethodMap) DelKey(key string) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Name
func (this MethodMap) Del(item *Method) bool {
	if this.ContainsKey(item.Name) {
		delete(this, item.Name)
		return true
	}
	return false
}

// Keys returns a new []string of all keys
func (this MethodMap) Keys() []string {
	keys := make([]string, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns a new sorted []string of all keys
func (this MethodMap) SortedKeys() []string {
	keys := this.Keys()
	sort.Strings(keys)
	return keys
}

// Each calls f for each map entry
func (this MethodMap) Each(f func(string, *Method)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Method containing all values
func (this MethodMap) Values() []*Method {
	values := make([]*Method, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

// Model is a struct
type Model struct {
	Name     string     `xml:"name,attr" `
	Packages PackageMap `xml:"package" `
	Path     string     `xml:"path,attr" `
}

// NewModel creates a new instance of Model
func NewModel() *Model {
	return &Model{
		Packages: NewPackageMap(),
	}
}

func (this *Model) Accept(v Visitor) {
	this.doAccept(v)
}

// Package is a struct
type Package struct {
	Imports ImportMap `xml:"import" `
	Model   *Model    `json:"-" xml:"-" `
	Name    string    `xml:"name,attr" `
	Types   TypeMap   `xml:"type" `
}

// NewPackage creates a new instance of Package
func NewPackage() *Package {
	return &Package{
		Imports: NewImportMap(), Types: NewTypeMap(),
	}
}

func (this *Package) Accept(v Visitor) {
	this.doAccept(v)
}

// PackageMap is a map[string]*Package
type PackageMap map[string]*Package

// NewPackageMap creates a new instance of PackageMap
func NewPackageMap() PackageMap {
	return make(PackageMap)
}

func (m PackageMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewPackage()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Name key exists in the PackageMap
func (this PackageMap) Contains(item *Package) bool {
	_, ok := this[item.Name]
	return ok
}

// ContainsKey returns true if the key exists in the PackageMap
func (this PackageMap) ContainsKey(key string) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the PackageMap
func (this PackageMap) Len() int {
	return len(this)
}

// Add adds the item using item.Name as the key
func (this PackageMap) Add(item *Package) bool {
	if !this.Contains(item) {
		this[item.Name] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this PackageMap) DelKey(key string) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Name
func (this PackageMap) Del(item *Package) bool {
	if this.ContainsKey(item.Name) {
		delete(this, item.Name)
		return true
	}
	return false
}

// Keys returns a new []string of all keys
func (this PackageMap) Keys() []string {
	keys := make([]string, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns a new sorted []string of all keys
func (this PackageMap) SortedKeys() []string {
	keys := this.Keys()
	sort.Strings(keys)
	return keys
}

// Each calls f for each map entry
func (this PackageMap) Each(f func(string, *Package)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Package containing all values
func (this PackageMap) Values() []*Package {
	values := make([]*Package, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

// Param is a struct
type Param struct {
	DataType string `xml:"data-type,attr" `
	Name     string `xml:"name,attr" `
}

// NewParam creates a new instance of Param
func NewParam() *Param {
	return &Param{}
}

// ParamMap is an ordered map of *Param keyed by *Param.Name
// key lookups are an O(N) operation, so don't use it for large maps
type ParamMap struct {
	list []*Param
}

// NewParamMap creates a new sorted map
func NewParamMap() *ParamMap {
	return &ParamMap{make([]*Param, 0, 8)}
}

func (this *ParamMap) Add(item *Param) bool {
	return this.InsertAt(this.Len(), item)
}

func (this *ParamMap) InsertAt(index int, item *Param) bool {
	if this.ContainsKey(item.Name) {
		return false
	}
	if 0 > index || index > this.Len() {
		return false
	}
	if index < this.Len() {
		var defaultValue *Param
		this.list = append(this.list, defaultValue)
		copy(this.list[index+1:], this.list[index:])
		this.list[index] = item
	} else {
		this.list = append(this.list, item)
	}
	return true
}

func (this *ParamMap) Len() int {
	return len(this.list)
}

func (this *ParamMap) Contains(item *Param) bool {
	return this.ContainsKey(item.Name)
}

func (this *ParamMap) ContainsKey(key string) bool {
	return this.IndexOfKey(key) >= 0
}

func (this *ParamMap) ContainsValue(item *Param) bool {
	return this.IndexOfValue(item) >= 0
}

func (this *ParamMap) IndexOf(item *Param) int {
	return this.IndexOfKey(item.Name)
}

func (this *ParamMap) IndexOfKey(key string) int {
	for index, item := range this.list {
		if item.Name == key {
			return index
		}
	}
	return -1
}

func (this *ParamMap) IndexOfValue(item *Param) int {
	for index, item_ := range this.list {
		if item == item_ {
			return index
		}
	}
	return -1
}

func (this *ParamMap) Each(f func(string, *Param)) {
	for _, item := range this.list {
		f(item.Name, item)
	}
}

func (this *ParamMap) Keys() []string {
	keys := make([]string, 0, this.Len())
	for _, item := range this.list {
		keys = append(keys, item.Name)
	}
	return keys
}

// SortedKeys returns a new sorted []string of all keys
func (this ParamMap) SortedKeys() []string {
	keys := this.Keys()
	sort.Strings(keys)
	return keys
}

func (this *ParamMap) Values() []*Param {
	values := make([]*Param, this.Len())
	copy(values, this.list)
	return values
}

func (this *ParamMap) Del(item *Param) bool {
	return this.DelKey(item.Name)
}

func (this *ParamMap) DelValue(item *Param) bool {
	index := this.IndexOfValue(item)
	return this.DelAt(index)
}

func (this *ParamMap) DelKey(key string) bool {
	index := this.IndexOfKey(key)
	return this.DelAt(index)
}

func (this *ParamMap) DelAt(index int) bool {
	if 0 <= index && index < this.Len() {
		var defaultValue *Param
		copy(this.list[index:], this.list[:index+1])
		this.list[this.Len()-1] = defaultValue
		this.list = this.list[:this.Len()-1]
		return true
	}
	return false
}

func (this *ParamMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewParam()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	this.Add(item)
	return nil
}

// Type is a struct
type Type struct {
	Fields   FieldMap  `xml:"field" `
	Meta     *MetaType `json:"-" xml:"-" `
	MetaType string    `xml:"meta-type,attr" `
	Methods  MethodMap `xml:"method" `
	Name     string    `xml:"name,attr" `
	Package  *Package  `json:"-" xml:"-" `
}

// NewType creates a new instance of Type
func NewType() *Type {
	return &Type{
		Fields: NewFieldMap(), Methods: NewMethodMap(),
	}
}

func (this *Type) Accept(v Visitor) {
	this.doAccept(v)
}

// TypeMap is a map[string]*Type
type TypeMap map[string]*Type

// NewTypeMap creates a new instance of TypeMap
func NewTypeMap() TypeMap {
	return make(TypeMap)
}

func (m TypeMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	item := NewType()
	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}
	m.Add(item)
	return nil
}

// Contains returns true if item.Name key exists in the TypeMap
func (this TypeMap) Contains(item *Type) bool {
	_, ok := this[item.Name]
	return ok
}

// ContainsKey returns true if the key exists in the TypeMap
func (this TypeMap) ContainsKey(key string) bool {
	_, ok := this[key]
	return ok
}

// Len returns the number of entries in the TypeMap
func (this TypeMap) Len() int {
	return len(this)
}

// Add adds the item using item.Name as the key
func (this TypeMap) Add(item *Type) bool {
	if !this.Contains(item) {
		this[item.Name] = item
		return true
	}
	return false
}

// DelKey removes the item by key
func (this TypeMap) DelKey(key string) bool {
	if this.ContainsKey(key) {
		delete(this, key)
		return true
	}
	return false
}

// Del removes the item by item.Name
func (this TypeMap) Del(item *Type) bool {
	if this.ContainsKey(item.Name) {
		delete(this, item.Name)
		return true
	}
	return false
}

// Keys returns a new []string of all keys
func (this TypeMap) Keys() []string {
	keys := make([]string, 0, this.Len())
	for key, _ := range this {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns a new sorted []string of all keys
func (this TypeMap) SortedKeys() []string {
	keys := this.Keys()
	sort.Strings(keys)
	return keys
}

// Each calls f for each map entry
func (this TypeMap) Each(f func(string, *Type)) {
	for key, val := range this {
		f(key, val)
	}
}

// Values retuns a new []*Type containing all values
func (this TypeMap) Values() []*Type {
	values := make([]*Type, 0, this.Len())
	for _, value := range this {
		values = append(values, value)
	}
	return values
}

type Visitor interface {
	VisitField(f *Field)
	VisitMethod(m *Method)
	VisitModel(m *Model)
	VisitPackage(p *Package)
	VisitType(t *Type)
}
