package model

type (
	Ident interface {
		Id() int64
		Name() string
	}
	ident struct {
		id   int64
		name string
	}
)

func (ident *ident) ctor(id int64, name string) *ident {
	ident.id = id
	ident.name = name
	return ident
}

func (ident *ident) Id() int64 {
	return ident.id
}

func (ident *ident) Name() string {
	return ident.name
}
