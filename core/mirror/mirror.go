package mirror

import (
	reflect "github.com/ReSc/godel/core/reflect"
)

type Any interface {
}

// DataSource is a struct
type DataSource struct {
	Current      Any
	CurrentIndex int
	DataMember   string
	DataValue    string
	Schema       *Schema
}

// NewDataSource creates a new instance of DataSource
func NewDataSource() *DataSource {
	return &DataSource{
		Schema: NewSchema(),
	}
}

// DockContainer is a struct
type DockContainer struct {
}

// NewDockContainer creates a new instance of DockContainer
func NewDockContainer() *DockContainer {
	return &DockContainer{}
}

// DockPanel is a struct
type DockPanel struct {
}

// NewDockPanel creates a new instance of DockPanel
func NewDockPanel() *DockPanel {
	return &DockPanel{}
}

// Input is a struct
type Input struct {
}

// NewInput creates a new instance of Input
func NewInput() *Input {
	return &Input{}
}

// Label is a struct
type Label struct {
}

// NewLabel creates a new instance of Label
func NewLabel() *Label {
	return &Label{}
}

// Layout is a struct
type Layout struct {
}

// NewLayout creates a new instance of Layout
func NewLayout() *Layout {
	return &Layout{}
}

// Page is a struct
type Page struct {
}

// NewPage creates a new instance of Page
func NewPage() *Page {
	return &Page{}
}

// Schema is a struct
type Schema struct {
	Type *reflect.Type
}

// NewSchema creates a new instance of Schema
func NewSchema() *Schema {
	return &Schema{}
}

// Stack is a struct
type Stack struct {
}

// NewStack creates a new instance of Stack
func NewStack() *Stack {
	return &Stack{}
}

// TabContainer is a struct
type TabContainer struct {
}

// NewTabContainer creates a new instance of TabContainer
func NewTabContainer() *TabContainer {
	return &TabContainer{}
}

// TabPanel is a struct
type TabPanel struct {
}

// NewTabPanel creates a new instance of TabPanel
func NewTabPanel() *TabPanel {
	return &TabPanel{}
}

// Table is a struct
type Table struct {
}

// NewTable creates a new instance of Table
func NewTable() *Table {
	return &Table{}
}
