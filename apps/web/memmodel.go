package main

import (
	"encoding/json"
	"github.com/ReSc/godel/core/reflect"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type MemoryModelHandler struct {
	mutex sync.Mutex
	model *reflect.Model
}

func NewMemoryModelHandler(path string) *MemoryModelHandler {
	h := &MemoryModelHandler{}
	model, err := reflect.LoadModelFile(path)
	if err != nil {
		panic(err.Error())
	}
	h.model = model
	return h
}

func (h *MemoryModelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	r.Header.Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	if name, ok := vars["name"]; ok {
		if pkg, ok := h.model.Packages[name]; ok {
			json.NewEncoder(w).Encode(pkg)
		} else {
			http.NotFound(w, r)
		}
	} else {
		names := h.model.Packages.SortedKeys()
		json.NewEncoder(w).Encode(names)
	}
}
