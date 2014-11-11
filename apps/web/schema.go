package main

import (
	_ "github.com/lib/pq"
	"io"
)

type (
	Graph interface {
		Node(id int64) (*Node, bool)
		NodeAt(path string) (*Node, bool)

		Children(parentId int64) []*Node
		ChildrenAt(parentPath string) []*Node
	}

	GraphLoader interface {
		Load(r io.Reader) (Graph, error)
	}

	Saver interface {
		Save(w io.Writer) error
	}

	SchemaLoader interface {
		Load() (*Schema, error)
	}

	schemaLoader struct {
		g Graph
	}

	Schema struct {
		Node    *Node
		Name    string
		Records []*Record
	}

	Record struct {
		Node   *Node
		Name   string
		Fields []*Field
	}

	Field struct {
		Node         *Node
		Name         string
		Type         string
		DefaultValue string
	}
)

func NewSchemaLoader(g Graph) SchemaLoader {
	return &schemaLoader{g: g}
}

func (l *schemaLoader) Load() (*Schema, error) {
	schema := new(Schema)
	for _, child := range l.g.ChildrenAt("/schema") {
		if err := l.loadRecordDefinition(schema, child); err != nil {
			return nil, err
		}
	}
	return schema, nil
}

func (l *schemaLoader) isRecordDefinition(recordDefinition *Node) bool {
	return false
}

func (l *schemaLoader) loadRecordDefinition(schema *Schema, recordDefinition *Node) error {
	if !l.isRecordDefinition(recordDefinition) {
		return nil
	}
	record := &Record{
		Node: recordDefinition,
		Name: recordDefinition.Name,
	}

	for _, child := range l.g.Children(recordDefinition.Id) {
		if err := l.loadFieldDefinition(record, child); err != nil {
			return err
		}
	}
	schema.AddRecord(record)
	return nil
}

func (schema *Schema) AddRecord(record *Record) {
	schema.Records = append(schema.Records, record)
}

func (l *schemaLoader) isFieldDefinition(fieldDefinition *Node) bool {
	return false
}

func (l *schemaLoader) loadFieldDefinition(record *Record, fieldDefinition *Node) error {
	if !l.isFieldDefinition(fieldDefinition) {
		return nil
	}

	field := &Field{
		Node: fieldDefinition,
		Name: fieldDefinition.Name,
	}

	record.AddField(field)
	return nil
}

func (record *Record) AddField(field *Field) {
	record.Fields = append(record.Fields, field)
}
