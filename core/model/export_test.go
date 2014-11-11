package model

import (
	"encoding/xml"
	"strings"
	"testing"
)

type (
	TestTagList struct {
		XMLName xml.Name `xml:"set"`
		Tags    TestSet  `xml:"tags,attr"`
	}
	TestSet []string
)

func (s TestSet) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: strings.Join(s, " "),
	}, nil
}

func TestXmlListAttribute(t *testing.T) {
	sut := &TestTagList{
		Tags: TestSet{"one", "two", "three"},
	}
	data, err := xml.Marshal(sut)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(data))
	}

}
