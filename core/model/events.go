package model

type (
	Event struct {
		Source Node
		Type   EventType
		Name   string
		Value  interface{}
	}
)

type (
	UnsubscribeFunc func()

	Observable interface {
		Subscribe(o Observer) UnsubscribeFunc
	}

	subject struct {
		source    *node
		observers []Observer
	}
)

func (subject *subject) ctor(source *node) *subject {
	subject.source = source
	return subject
}

func (subject *subject) Subscribe(o Observer) UnsubscribeFunc {
	subject.observers = append(subject.observers, o)
	return func() { subject.unsubscribe(o) }
}

func (subject *subject) unsubscribe(o Observer) {
	for i, oldo := range subject.observers {
		if oldo == o {
			oo := subject.observers
			l := len(oo)
			copy(oo[i:], oo[i+1:])
			oo[l-1] = nil
			subject.observers = oo[:l-1]
			return
		}
	}
}

func (subject *subject) raise(e *Event) {
	if e == nil {
		panic("e is nil")
	}

	for _, o := range subject.observers {
		o.OnNext(e)
	}

	if subject.source != nil && subject.source.parent != nil {
		subject.source.parent.raise(e)
	}
}

func (subject *subject) raiseSignal(name string) {
	subject.raise(&Event{Source: subject.source, Type: Signal, Name: name, Value: name})
}

func (subject *subject) raiseChildSet(child Node) {
	subject.raise(&Event{Source: subject.source, Type: ChildSet, Name: child.Name(), Value: child})
}

func (subject *subject) raiseChildDel(child Node) {
	subject.raise(&Event{Source: subject.source, Type: ChildDel, Name: child.Name(), Value: child})
}

func (subject *subject) raiseTagSet(tag string) {
	subject.raise(&Event{Source: subject.source, Type: TagSet, Name: tag, Value: tag})
}

func (subject *subject) raiseTagDel(tag string) {
	subject.raise(&Event{Source: subject.source, Type: TagDel, Name: tag, Value: tag})
}

func (subject *subject) raiseAttrSet(attr string, value interface{}) {
	subject.raise(&Event{Source: subject.source, Type: AttrSet, Name: attr, Value: value})
}

func (subject *subject) raiseAttrDel(attr string, oldValue interface{}) {
	subject.raise(&Event{Source: subject.source, Type: AttrDel, Name: attr, Value: oldValue})
}

type (
	OnNextFunc  func(e *Event)
	OnErrorFunc func(err error)
	OnDoneFunc  func()

	Observer interface {
		OnNext(e *Event)
		OnError(err error)
		OnDone()
	}

	observer struct {
		onNext  OnNextFunc
		onError OnErrorFunc
		onDone  OnDoneFunc
	}
)

func NewObserverNext(onNext OnNextFunc) Observer {
	return NewObserver(onNext, nil, nil)
}

func NewObserverDone(onDone OnDoneFunc) Observer {
	onError := func(err error) {
		if onDone != nil {
			onDone()
		}
	}
	return NewObserver(nil, onError, onDone)
}

func NewObserverNextDone(onNext OnNextFunc, onDone OnDoneFunc) Observer {
	onError := func(err error) {
		if onDone != nil {
			onDone()
		}
	}
	return NewObserver(onNext, onError, onDone)
}

func NewObserver(onNext OnNextFunc, onError OnErrorFunc, onDone OnDoneFunc) Observer {
	return &observer{
		onNext:  onNext,
		onError: onError,
		onDone:  onDone,
	}
}

func (observer *observer) OnNext(e *Event) {
	f := observer.onNext
	if f != nil {
		f(e)
	}
}

func (observer *observer) OnError(err error) {
	f := observer.onError
	observer.clear()
	if f != nil {
		f(err)
	}
}

func (observer *observer) OnDone() {
	f := observer.onDone
	observer.clear()
	if f != nil {
		f()
	}
}

func (observer *observer) clear() {
	observer.onNext = nil
	observer.onError = nil
	observer.onDone = nil
}
