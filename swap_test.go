package swap

import (
	"testing"
	"unsafe"
)

type example struct {
	nested nestedFields
}

type nestedFields struct {
	subPrivateFieldOne   string
	subPrivateFieldTwo   int
	subPrivateFieldThree int32
	subPrivateFieldFour  int64
	subPrivateFieldFive  bool
	subPrivateFieldSix   float64
	subPrivateFieldSeven float32
}

func TestSwapper_Integration(t *testing.T) {
	obj := example{
		nested: nestedFields{
			subPrivateFieldOne:   "",
			subPrivateFieldTwo:   0,
			subPrivateFieldThree: 0,
			subPrivateFieldFour:  0,
			subPrivateFieldFive:  false,
			subPrivateFieldSix:   0,
			subPrivateFieldSeven: 0,
		},
	}

	swapper := Swapper{}
	myPointer := swapper.Init(&obj).
		Find("nested", func(p unsafe.Pointer) interface{} {
			return (*nestedFields)(p)
		}).
		Find("subPrivateFieldOne", func(p unsafe.Pointer) interface{} {
			return (*string)(p)
		}).Pointer()

	str := myPointer.(*string)
	*str = "test"

	if obj.nested.subPrivateFieldOne != "test" {
		t.Errorf("unexpected value for subPrivateFieldOne: %s", obj.nested.subPrivateFieldOne)
	}
}
