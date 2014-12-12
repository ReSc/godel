package main

import (
	"github.com/ReSc/fmt"
	"github.com/ReSc/godel/core/graph"
	"github.com/ReSc/godel/core/mvc"
	"net/http"
)

type GraphController struct {
	mvc.ControllerBase
	graph *graph.Graph
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
