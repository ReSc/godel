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

func (c *GraphController) PostNodeEditor() http.Handler {
	fmt.Printline("POST")
	type args struct {
		Name string `json:"name"`
	}
	a := new(args)
	if err := c.JsonBody(a); err != nil {
		fmt.Printline("%v", err)
		return c.InternalServerError(err.Error())
	}
	if id, ok := c.IdAsInt64(); ok {
		if n, ok := c.graph.Nodes[id]; ok {
			n.Attrs.Set("name", a.Name)
			fmt.Printline("REDIRECTING!")
			return c.RedirectToMyAction("node-editor")
		}
	}

	fmt.Printline("BADREQUEST!")
	return c.BadRequest()
}

func (c *GraphController) GetNodeEditor() http.Handler {
	if id, ok := c.IdAsInt64(); ok {
		if n, ok := c.graph.Nodes[id]; ok {
			return c.View(n)
		}
	}
	return c.NotFound()
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
