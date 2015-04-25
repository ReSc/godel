package main

import (
	"github.com/ReSc/godel/core/mvc"
	"github.com/ReSc/godel/core/reflect"
	"net/http"
)

type ModelController struct {
	mvc.ControllerBase
	model *reflect.Model
}

func (c *ModelController) ensureModel() {
	if c.model == nil {
		panic("model not available")
	}
}

func (c *ModelController) GetIndex2() http.Handler {
	c.ensureModel()
	names := c.model.Packages.SortedKeys()
	return c.Json(names)
}

func (c *ModelController) GetPackage() http.Handler {
	c.ensureModel()
	if id, ok := c.Id(); ok {
		if pkg, ok := c.model.Packages[id]; ok {
			return c.Json(pkg)
		} else {
			return c.NotFound()
		}
	} else {
		names := c.model.Packages.SortedKeys()
		return c.Json(names)
	}
}
