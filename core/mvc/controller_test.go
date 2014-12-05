package mvc

import (
	. "testing"
)

func Test1(t *T) {
	d := NewDispatcher()
	d.Register(func() Controller { return &ModelController{} })

	for k, _ := range d.factories {
		t.Log("registered ", k)
	}
}
