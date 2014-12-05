package main

import (
	"github.com/ReSc/godel/core/graph"
	"github.com/ReSc/godel/core/mvc"
	"net/http"
)

type GraphController struct {
	mvc.ControllerBase
	graph *graph.Graph
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
