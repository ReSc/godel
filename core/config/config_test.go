package config

import (
	"encoding/xml"
	"github.com/ReSc/godel/core/model"
	//	"io"
	"strings"
	"testing"
)

func TestImportBootstrap(t *testing.T) {
	r := strings.NewReader(Bootstrap)
	r.Seek(0, 0)
	d := xml.NewDecoder(r)
	m, err := NewConfigurator().ParseXML(d)
	if err != nil {
		t.Fatal(err)
	}
	dump(t, m)
}

func TestImport(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
<config id="0" name="test" tags="godel-config-version-1">
	<node id="1" name="child" tags="cxcbv">
		<node id="2" name="grandchild" tags="xbcvb">
			<node id="3" name="jan" tags="t1" />
			<node id="4" name="piet" tags="t1 t2" />
		</node>
	</node> 
</config>`
	r := strings.NewReader(data)
	r.Seek(0, 0)
	d := xml.NewDecoder(r)
	m, err := NewConfigurator().ParseXML(d)
	if err != nil {
		t.Fatal(err)
	}

	dump(t, m)
	t.Log("\n")
	if m.Name() != "test" {
		t.Fatal(m.Name(), "test")
	}
}

func dump(t *testing.T, m *model.Model) {

	indent := func(i int) string {
		return strings.Repeat("  ", i)
	}

	var printer func(model.Node, int)

	printer = func(n model.Node, i int) {
		t.Logf("%v %v", indent(i), n.Name())
		n.Tags().Each(func(tg string) {
			t.Logf("%v tag: %v", indent(i+1), tg)
		})

		n.Children().Each(func(n model.Node) {
			printer(n, i+1)
		})
	}

	printer(m, 0)
}
