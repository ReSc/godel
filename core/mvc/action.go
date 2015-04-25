package mvc

import (
	"net/http"
	"reflect"

	"github.com/ReSc/fmt"
)

type action struct {
	name       string
	httpMethod string
	method     reflect.Value
}

func invoke(c Controller, a *action) http.Handler {
	receiver := reflect.ValueOf(c)
	args := []reflect.Value{receiver}
	results := a.method.Call(args)
	result := results[0].Interface()

	if handler, ok := result.(http.Handler); ok {
		if handler != nil {
			return handler
		} else {
			return &statusHandler{
				code:    http.StatusNoContent,
				message: http.StatusText(http.StatusNoContent),
			}
		}
	} else {
		return &statusHandler{
			code:    http.StatusInternalServerError,
			message: http.StatusText(http.StatusInternalServerError),
		}
	}
}

func verifyMethodSignature(c interface{}, a *action) bool {
	controllerValue := reflect.ValueOf(c)
	ctrl := controllerValue.Interface()
	m := a.method
	mType := a.method.Type()
	if mType.NumIn() != 1 {
		fmt.Printline("argument mismatch for action %s, %+v ", mType.Name(), mType)
		return false
	}

	if mType.NumOut() != 1 {
		fmt.Printline("too many return values for action %s, %+v , %+v ,%+v", mType.Name(), controllerValue, ctrl, m)
		return false
	}

	if !mType.Out(0).AssignableTo(handlerInterface) {
		fmt.Printline("wrong type of return value for action %s, %+v , %+v ,%+v", m.Type().Name(), controllerValue, ctrl, m)
		return false
	}

	return true
}

func newActionMap() *actionMap {
	return &actionMap{}
}

type actionMap struct {
	actions []*action
}

func (m *actionMap) Add(c interface{}, httpMethod string, name string, method reflect.Value) bool {
	action := &action{
		name:       name,
		method:     method,
		httpMethod: httpMethod,
	}
	if verifyMethodSignature(c, action) {
		m.actions = append(m.actions, action)
		return true
	} else {
		return false
	}
}

func (m *actionMap) Get(httpMethod string, name string) (*action, bool) {
	for i := range m.actions {
		a := m.actions[i]
		if a.httpMethod == httpMethod &&
			a.name == name {
			return a, true
		}
	}
	return nil, false
}
