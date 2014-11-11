package main

import (
	"encoding/xml"
	"sort"
)

type (
	Attr struct {
		XMLName xml.Name `json:"-" xml:"attr"`
		Key     string   `json:"key" xml:"name,attr"`
		Value   string   `json:"val" xml:",chardata"`
	}

	Attrs []Attr
)

func (attrs Attrs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch len(attrs) {
	case 0:
		return nil
	case 1:
		return e.Encode(attrs[0])
	default:
		sortedAttrs := make(Attrs, len(attrs))
		copy(sortedAttrs, attrs)
		sort.Stable(AttrsByKey(sortedAttrs))
		for _, attr := range sortedAttrs {
			if attr.Value == "" {
				continue
			}
			if err := e.Encode(attr); err != nil {
				return err
			}
		}
		return nil
	}
}

type AttrsByKey Attrs

func (attrs AttrsByKey) Len() int           { return len(attrs) }
func (attrs AttrsByKey) Less(i, j int) bool { return attrs[i].Key < attrs[j].Key }
func (attrs AttrsByKey) Swap(i, j int)      { attrs[j], attrs[i] = attrs[i], attrs[j] }
