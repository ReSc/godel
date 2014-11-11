package model

type (
	Attribute interface {
		Name() string
		Value() interface{}
	}

	Attributes interface {
		Each(func(Attribute))
		Get(name string) (attr Attribute, ok bool)
	}

	attributes map[string]*attribute

	attribute struct {
		name  string
		value interface{}
	}
)

func (a *attribute) Name() string       { return a.name }
func (a *attribute) Value() interface{} { return a.value }

func (attrs attributes) Each(f func(Attribute)) {
	for _, v := range attrs {
		f(v)
	}
}

func (attrs attributes) Get(name string) (Attribute, bool) {
	attr, ok := attrs[name]
	return attr, ok
}

func (attrs attributes) add(name string, value interface{}) (*attribute, bool) {
	if _, ok := attrs[name]; ok {
		return nil, false
	}
	attr := &attribute{name: name, value: value}
	attrs[name] = attr
	return attr, true
}
