package main

type (
	DataSource interface {
		Id() int64
		Name() string
		Len() int
		Fields() Fields
		Range(start, count int) DataSourceIterator
	}

	DataSourceIterator interface {
		Iterator
		Value(f Field) Value
	}

	Fields interface {
		Next() bool
		Field() Field
	}

	Field interface {
		Id() int64
		Name() string
	}

	Value interface {
		Id() int64
		Name() string
		String() string
		Interface() interface{}
	}

	TableRenderer interface {
		Begin(ds DataSource, start, count int)
		BeginRow(index int)
		Cell(field Field, value Value)
		EndRow()
		End()
	}
)

func Render(r TableRenderer, ds DataSource, start, count int) {
	row := ds.Range(start, count)

	r.Begin(ds, start, count)
	for index := 0; row.Next(); index++ {
		r.BeginRow(index)
		for f := ds.Fields(); f.Next(); {
			field := f.Field()
			value := row.Value(field)
			r.Cell(field, value)
		}
		r.EndRow()
	}
	r.End()
}
