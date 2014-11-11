package main

import (
	"encoding/xml"
	"sort"
)

type Rels []*Rel

type Rel struct {
	XMLName    xml.Name `json:"-" xml:"rel"`
	Subject    int64    `json:"sub" xml:"sub,attr"`
	Predicate  int64    `json:"pre" xml:"pre,attr,omitempty"`
	Object     int64    `json:"obj" xml:"obj,attr"`
	Provenance int64    `json:"prv" xml:"prv,attr"`
}

func (rels Rels) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch len(rels) {
	case 0:
		return nil
	case 1:
		return e.Encode(rels[0])
	default:
		sortedRels := make(Rels, len(rels))
		copy(sortedRels, rels)
		sort.Stable(RelsById(sortedRels))
		for _, rel := range sortedRels {
			if err := e.Encode(rel); err != nil {
				return err
			}
		}
		return nil
	}
}

type RelsById Rels

func (rels RelsById) Len() int      { return len(rels) }
func (rels RelsById) Swap(i, j int) { rels[j], rels[i] = rels[i], rels[j] }

func (rels RelsById) Less(i, j int) bool {
	if rels[i].Subject < rels[j].Subject {
		return true
	}
	if rels[i].Object < rels[j].Object {
		return true
	}
	if rels[i].Predicate < rels[j].Predicate {
		return true
	}
	return false
}
