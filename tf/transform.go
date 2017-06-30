package tf

import (
	"github.com/v2pro/plz/acc"
)

func Transform(fromAccessor acc.Accessor, toAccessor acc.Accessor) (acc.Accessor, error) {
	kind := toAccessor.Kind()
	if fromAccessor.Kind() == acc.Struct && kind == acc.Struct {
		return fromAccessor, nil
	}
	if fromAccessor.Kind() != acc.Struct && kind == acc.Struct {
		kind = acc.Map
	}
	return &transformedAccessor{
		kind:         kind,
		toAccessor:   toAccessor,
		fromAccessor: fromAccessor,
	}, nil
}

type transformedAccessor struct {
	acc.NoopAccessor
	kind         acc.Kind
	toAccessor   acc.Accessor
	fromAccessor acc.Accessor
}

func (accessor *transformedAccessor) Kind() acc.Kind {
	return accessor.kind
}

func (accessor *transformedAccessor) GoString() string {
	return accessor.fromAccessor.GoString()
}

func (accessor *transformedAccessor) Key() acc.Accessor {
	return accessor.fromAccessor.Key()
}

func (accessor *transformedAccessor) Elem() acc.Accessor {
	return accessor.fromAccessor.Elem()
}

func (accessor *transformedAccessor) NumField() int {
	return accessor.toAccessor.NumField()
}

func (accessor *transformedAccessor) Field(index int) acc.StructField {
	field := accessor.toAccessor.Field(index)
	field.Accessor = accessor.Elem()
	return field
}

func (accessor *transformedAccessor) IterateMap(obj interface{}, cb func(key interface{}, elem interface{}) bool) {
	accessor.fromAccessor.IterateMap(obj, cb)
}

func (accessor *transformedAccessor) IterateArray(obj interface{}, cb func(elem interface{}) bool) {
	accessor.fromAccessor.IterateArray(obj, cb)
}

func (accessor *transformedAccessor) Int(obj interface{}) int {
	return accessor.fromAccessor.Int(obj)
}
