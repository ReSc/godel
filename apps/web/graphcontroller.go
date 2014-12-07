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

func (c *GraphController) GetNodeEditor() http.Handler {
	if id, ok := c.IdAsInt64(); ok {
		if n, ok := c.graph.Nodes[id]; ok {
			return c.View(n)
		}
	}
	return c.NotFound()
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
