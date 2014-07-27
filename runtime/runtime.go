package runtime

import (
	"fmt"
)

type Object interface {
}

type Error interface {
}

type Runtime interface {
	Self() Object
	StringLiteral(string, string) Object
	StringDup(Object) Object
	AllowPrivate()
	True() Object
	SendSite(Object, string) func(...Object) Object
}

type runtime struct {
	self          Object
	_true         Object
	_false        Object
	_nil          Object
	allow_private bool
}

func (rt *runtime) Self() Object {
	return rt.self
}

func (rt *runtime) AllowPrivate() {
	rt.allow_private = true
}

func (rt *runtime) True() Object {
	return nil
}

func (rt *runtime) SendSite(obj Object, sym string) func(...Object) Object {
	return func(args ...Object) Object {
		fmt.Println(args)
		return nil
	}
}

func NewRuntime() Runtime {
	rt := &runtime{}
	return rt
}
