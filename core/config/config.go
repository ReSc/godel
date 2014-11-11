package config

import (
	"encoding/xml"
	"errors"
	"github.com/ReSc/fmt"
	"github.com/ReSc/godel/core/model"
	"io"
	"strconv"
	"strings"
)

type (
	Configurator struct {
		*model.Model
		*model.Editor
	}
)

func (c *Configurator) reset() {
	c.Model = model.New("model")
	c.Editor = model.NewEditor(c.Model)
}

func NewConfigurator() *Configurator {
	return &Configurator{}
}

func (c *Configurator) ParseXML(d *xml.Decoder) (*model.Model, error) {
	c.reset()
	for {
		t, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return c.Model, nil
			}
			return nil, err
		}

		switch t := t.(type) {
		case xml.ProcInst:
		case xml.Directive:
		case xml.StartElement:
			switch t.Name.Local {
			case "config":

				err := c.parseNode(c.Model, d, t)
				if err == nil {
					return c.Model, nil
				}
				return nil, err
			default:
				return nil, &xml.SyntaxError{Msg: "Unexpected token " + t.Name.Local}
			}
		case xml.EndElement:
		case xml.CharData:
		case xml.Comment:
		default:
			return nil, &xml.SyntaxError{Msg: "Unexpected token type" + fmt.String("%T", t)}
		}
	}

	return nil, &xml.SyntaxError{Msg: "nothing to parse"}
}

func (c *Configurator) parseNode(n model.Node, d *xml.Decoder, e xml.StartElement) error {
	print(">", n.FullName(), "\n")
	if err := c.parseAttrs(n, e); err != nil {
		return &xml.SyntaxError{Msg: err.Error()}
	}

	for {
		t, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch t := t.(type) {
		case xml.ProcInst:
		case xml.Directive:
		case xml.StartElement:
			switch t.Name.Local {
			case "node":
				child, err := c.NewChild(n, "new_child__")
				if err != nil {
					panic("error creating new node on parent " + n.FullName() + ": " + err.Error())
				}
				err = c.parseNode(child, d, t)
				if err != nil {
					return err
				}
			default:
				return &xml.SyntaxError{Msg: "unexpected element " + e.Name.Local}
			}
		case xml.EndElement:
			print("<", n.FullName(), "\n")
			return nil
		case xml.CharData:
		case xml.Comment:
		default:
			return &xml.SyntaxError{Msg: "unexpected token type " + fmt.String("%T", t)}
		}
	}
}

func (c *Configurator) parseAttrs(n model.Node, e xml.StartElement) error {
	attr := attrs(e.Attr)
	if id, ok := attr.Int64("id"); ok {
		c.SetId(n, id)
	}

	if name, ok := attr.Value("name"); ok {
		if err := c.SetName(n, name); err != nil {
			return errors.New(err.Error() + " on " + n.FullName())
		}
	} else {
		return &xml.SyntaxError{Msg: "mandatory attribute name not found on " + n.FullName()}
	}

	if tags, ok := attr.Strings("tags", " "); ok {
		if err := c.SetTag(n, tags...); err != nil {
			return errors.New(err.Error() + " on " + n.FullName())
		}
	}

	return nil
}

type (
	attrs []xml.Attr
)

func (a attrs) Each(f func(xml.Attr)) {
	for i := range a {
		f(a[i])
	}
}

func (a attrs) indexOf(name string) int {
	for i := range a {
		if a[i].Name.Local == name {
			return i
		}
	}
	return -1
}

func (a attrs) Value(name string) (string, bool) {
	if i := a.indexOf(name); i >= 0 {
		return a[i].Value, true
	}
	return "", false
}

func (a attrs) Strings(name string, sep string) ([]string, bool) {
	if i := a.indexOf(name); i >= 0 {
		return strings.Split(a[i].Value, sep), true
	}
	return nil, false
}

func (a attrs) Int64(name string) (int64, bool) {
	if i := a.indexOf(name); i >= 0 {
		val, err := strconv.ParseInt(a[i].Value, 10, 64)
		if err != nil {
			return 0, false
		}
		return val, true
	}
	return 0, false
}

func (a attrs) Float64(name string) (float64, bool) {
	if i := a.indexOf(name); i >= 0 {
		val, err := strconv.ParseFloat(a[i].Value, 64)
		if err != nil {
			return 0, false
		}
		return val, true
	}
	return 0, false
}
