package main

import (
	//	"github.com/ReSc/fmt"
	"github.com/ReSc/godel/core/rest"
	"net/http"
	"strconv"
	"strings"
)

// node resource

type nodeResource struct {
	rest.Resource

	nodes Nodes
}

func newNodeResource() *nodeResource {
	return &nodeResource{
		nodes: bootstrapNodes(),
	}
}

// HTTP handlers
//
// Put(w http.ResponseWriter, r *http.Request)
// Patch(w http.ResponseWriter, r *http.Request)

func (res *nodeResource) Post(w http.ResponseWriter, r *http.Request) {

	print("=> POST ", r.RequestURI, "\n")
	vars := res.Vars(r)

	if _, ok := vars["id"]; ok {
		res.Error(w, "POST not allowed", http.StatusMethodNotAllowed)
		return
	}

	node := new(Node)

	if err := res.DecodeBody(w, r, node); err != nil {
		print(err.Error())
		return
	}

	node.Id = res.getMaxId() + 1
	res.nodes.append(node)
	nodePath, _ := res.Path(rest.Params{"id": node.Id})
	res.SeeOther(w, nodePath)
}

func (res *nodeResource) Delete(w http.ResponseWriter, r *http.Request) {

	print("=> DELETE ", r.RequestURI, "\n")
	vars := res.Vars(r)

	if id, ok := vars["id"]; ok {

		if idval, err := strconv.ParseInt(id, 10, 64); err != nil {
			res.Error(w, stripSource(err), http.StatusBadRequest)
		} else {
			if n, ok := res.nodes.get(idval); ok {
				res.nodes.del(idval)
				res.Return(w, r, n)
			} else {
				res.Error(w, "Not found", http.StatusNotFound)
			}
		}

	} else if nodeIds, ok := r.URL.Query()["id"]; ok {

		if ids, err := parseInt64(nodeIds...); err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			children := res.nodes.FindAll(func(n *Node) bool {
				return containsInt64(ids, n.Id)
			})
			for _, n := range children {
				res.nodes.del(n.Id)
			}
			res.Return(w, r, children)
		}

	} else {

		res.Error(w, "Please supply id or parentId query parameters", http.StatusBadRequest)

	}
}

func (res *nodeResource) Get(w http.ResponseWriter, r *http.Request) {

	print("=> GET ", r.RequestURI, "\n")
	vars := res.Vars(r)

	if id, ok := vars["id"]; ok {

		if idval, err := strconv.ParseInt(id, 10, 64); err != nil {
			res.Error(w, stripSource(err), http.StatusBadRequest)
		} else {
			if n, ok := res.nodes.get(idval); ok {
				res.Return(w, r, n)
			} else {
				res.Error(w, "Not found", http.StatusNotFound)
			}
		}

	} else if parentId := r.URL.Query().Get("parentId"); parentId != "" {

		if parentIds, err := parseInt64(parentId); err != nil {
			res.Error(w, stripSource(err), http.StatusBadRequest)
		} else {
			children := res.nodes.findByParentId(parentIds[0])
			res.Return(w, r, children)
		}

	} else if nodeIds, ok := r.URL.Query()["id"]; ok {

		if ids, err := parseInt64(nodeIds...); err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			children := res.nodes.FindAll(func(n *Node) bool {
				return containsInt64(ids, n.Id)
			})
			res.Return(w, r, children)
		}

	} else {

		res.Error(w, "Please supply id or parentId query parameters", http.StatusBadRequest)

	}
}

func (res *nodeResource) getMaxId() int64 {
	maxId := int64(0)
	res.nodes.Each(func(n *Node) {
		if n.Id > maxId {
			maxId = n.Id
		}
	})
	return maxId
}

func stripSource(e error) string {
	msg := e.Error()
	i := strings.Index(msg, ":")
	if i > 0 {
		return msg[i+len(":"):]
	}
	return msg
}

func containsInt64(list []int64, val int64) bool {
	for i := range list {
		if list[i] == val {
			return true
		}
	}
	return false
}

func parseInt64(vals ...string) ([]int64, error) {
	var ints []int64
	for _, val := range vals {
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			ints = append(ints, i)
		} else {
			return nil, err
		}
	}
	return ints, nil
}
