package model

import (
	"testing"
)

func TestFullNameOnModel(t *testing.T) {
	m := New("model")
	result := m.FullName()
	if result != "/model" {
		t.Error("expected '/model' not '" + result + "'")
	}
}

func TestFullNameOnNode(t *testing.T) {
	m := New("model")
	e := NewEditor(m)
	child, _ := e.NewChild(m, "child")

	if result := child.FullName(); result != "/model/child" {
		t.Errorf("expected non-nil result, got " + result)
	}
}

func TestSubscribeOnModel(t *testing.T) {
	m := New("model")
	var e *Event
	unsub := m.Subscribe(NewObserverNext(func(ev *Event) {
		e = ev
		// t.Logf("%v{ Node: '%v', Name: '%v', Value: '%+v' }", ev.Type, ev.Source.Name(), ev.Name, ev.Value)
	}))
	defer unsub()

	editor := NewEditor(m)
	editor.NewChild(m, "child")

	if e == nil || e.Type != ChildSet {
		t.Error("expected non-nil event")
	}
}
