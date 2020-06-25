package swap

import (
	"reflect"
	"unsafe"
)

type Swapper struct {
	current interface{}
}

type Swap interface {
	Find(field string, castFn func(p unsafe.Pointer) interface{}) Swap
	Pointer() unsafe.Pointer
}

func (s *Swapper) Init(obj interface{}) *Swapper {
	s.current = obj
	return s
}

func (s *Swapper) Find(field string, castFn func(p unsafe.Pointer) interface{}) *Swapper {
	unsafePointer := getUnsafePointer(field, s.current)
	s.current = castFn(unsafePointer)
	return s
}

func getUnsafePointer(field string, obj interface{}) unsafe.Pointer {
	objVal := reflect.ValueOf(obj)
	fieldVal := reflect.Indirect(objVal).FieldByName(field)
	pointer := fieldVal.UnsafeAddr()
	return unsafe.Pointer(pointer)
}

func (s *Swapper) Pointer() interface{} {
	return s.current
}
