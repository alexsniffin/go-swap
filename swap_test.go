package swap

import (
	"fmt"
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
	ex := example{
		nested: nestedFields{
			subPrivateFieldOne:   "old",
			subPrivateFieldTwo:   0,
			subPrivateFieldThree: 0,
			subPrivateFieldFour:  0,
			subPrivateFieldFive:  false,
			subPrivateFieldSix:   0,
			subPrivateFieldSeven: 0,
		},
	}

	myPointer := Init(&ex).
		Find("nested", func(p unsafe.Pointer) interface{} {
			return (*nestedFields)(p)
		}).
		Find("subPrivateFieldOne", func(p unsafe.Pointer) interface{} {
			return (*string)(p)
		}).Pointer()

	fmt.Println(fmt.Sprintf("%#v", ex))
	str := myPointer.(*string)
	*str = "new"
	fmt.Println(fmt.Sprintf("%#v", ex))

	if ex.nested.subPrivateFieldOne != "new" {
		t.Errorf("unexpected value for subPrivateFieldOne: %s", ex.nested.subPrivateFieldOne)
	}
}
