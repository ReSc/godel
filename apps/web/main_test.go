package main

import (
	"encoding/xml"
	"testing"
)

func TestXmlEncoding(t *testing.T) {
	n := bootstrapNodes()
	s, err := xml.Marshal(n)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(string(s))
	}
}
