package main

import (
	"encoding/xml"
	"github.com/ReSc/fmt"
	"github.com/ReSc/godel/core/graph"
	"github.com/ReSc/godel/core/mvc"
	"net/http"
	"os"
	"strings"
)

type GraphController struct {
	mvc.ControllerBase
	graph *graph.Graph
}

func (c *GraphController) PostLoad() http.Handler {
	return c.NotFound()
}

func (c *GraphController) PostSave() http.Handler {
	path := "./data/graph.xml"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0x755)
	if err != nil {
		return c.InternalServerError(err.Error())
	} else {
		defer file.Close()
	}

	e := xml.NewEncoder(file)
	e.Indent("", "\t")
	err = e.Encode(c.graph)
	if err != nil {
		return c.InternalServerError(err.Error())
	}
	e.Flush()
	return c.RedirectToMyAction("index")
}

func (c *GraphController) GetNavigator() http.Handler {
	type Edge struct {
		Id       int64  `json:"id"`
		NodeId   int64  `json:"nodeid"`
		PredId   int64  `json:"predid"`
		NodeName string `json:"nodename"`
		PredName string `json:"predname"`
	}

	type Attr struct {
		Name  string
		Value string
	}

	type Node struct {
		Id    int64    `json:"id"`
		Name  string   `json:"name"`
		Tags  []string `json:"tags"`
		Attrs []*Attr  `json:"attrs"`
		In    []*Edge  `json:"in"`
		Out   []*Edge  `json:"out"`
	}

	if id, ok := c.IdAsInt64(); ok {
		nn := c.graph.Nodes
		if n, ok := nn[id]; ok {
			result := &Node{
				Id:   id,
				Name: n.Name(),
			}

			for tag, _ := range n.Tags {
				result.Tags = append(result.Tags, tag)
			}

			for _, attr := range n.Attrs {
				result.Attrs = append(result.Attrs, &Attr{attr.Name, attr.Value})
			}

			for _, e := range n.InEdges {
				result.In = append(result.In, &Edge{
					Id:       e.Id,
					NodeId:   e.Sub,
					NodeName: nn[e.Sub].Name(),
					PredId:   e.Prd,
					PredName: nn[e.Prd].Name(),
				})
			}

			for _, e := range n.OutEdges {
				result.Out = append(result.Out, &Edge{
					Id:       e.Id,
					NodeId:   e.Obj,
					NodeName: nn[e.Obj].Name(),
					PredId:   e.Prd,
					PredName: nn[e.Prd].Name(),
				})
			}

			return c.View(result)
		}
	}

	return c.NotFound()
}

func (c *GraphController) GetAttributes() http.Handler {
	type Result struct {
		Name string `json:"name"`
		Desc string `json:"description"`
	}
	query := c.Request.URL.Query().Get("q")
	found := make(map[string]bool)
	results := make([]Result, 0, 15)
	for _, n := range c.graph.Nodes {
		for _, attr := range n.Attrs {
			if strings.Contains(attr.Name, query) {
				if _, ok := found[attr.Name]; ok {
					continue
				}
				found[attr.Name] = true
				results = append(results, Result{attr.Name, attr.Name})
				if len(results) == cap(results) {
					break
				}
			}
		}
		if len(results) == cap(results) {
			break
		}
	}
	return c.Json(results)
}

func (c *GraphController) GetNodeCreate() http.Handler {
	return c.View("")
}

func (c *GraphController) PostNodeCreate() http.Handler {
	type Form struct {
		Name string `json:"name"`
	}

	form := new(Form)
	if err := c.FormBody(form); err == nil {
		n := c.graph.NewNode(form.Name)
		c.graph.Nodes.Add(n)
		return c.RedirectTo(c.Controller(), "node-editor", fmt.String("%d", n.Id))
	} else {
		return c.InternalServerError(err.Error())
	}

	return c.BadRequest()
}

func (c *GraphController) GetNodeList() http.Handler {
	if _, ok := c.Var("id"); !ok {
		keys := c.graph.Nodes.Keys()
		nodeNames := make(map[int64]string)
		for _, key := range keys {
			nodeNames[key] = c.graph.Nodes[key].Attrs["name"].Value
		}
		return c.View(nodeNames)
	}

	return c.BadRequest()
}

func (c *GraphController) GetNodeEditor() http.Handler {
	if id, ok := c.IdAsInt64(); ok {
		if n, ok := c.graph.Nodes[id]; ok {
			return c.View(n)
		}
	}

	return c.RedirectToMyAction("node-list")
}

func (c *GraphController) PostNodeEditor() http.Handler {
	if id, ok := c.IdAsInt64(); ok {
		if n, ok := c.graph.Nodes[id]; ok {
			c.Request.ParseForm()
			form := c.Request.Form
			for _, key := range n.Attrs.Keys() {
				if _, ok := form[key]; ok {
					n.Attrs.Set(key, form.Get(key))
				}
			}
			return c.RedirectTo(c.Controller(), c.Action(), fmt.String("%d", id))
		}
	}

	return c.BadRequest()
}

func (c *GraphController) GetNode() http.Handler {
	if id, ok := c.IdAsInt64(); ok {
		if n, ok := c.graph.Nodes[id]; ok {
			return c.Json(n)
		}
	}
	return c.NotFound()
}

func (c *GraphController) GetEdge() http.Handler {
	if id, ok := c.IdAsInt64(); ok {
		if e, ok := c.graph.Edges[id]; ok {
			return c.Json(e)
		}
	}
	return c.NotFound()
}
