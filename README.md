# go-swap

Testing utility for swapping out private field members without the need of dependency inversion of control.

### Example
```go
ex := example{
    nested: nestedFields{
        subPrivateField: "old",
    },
}
swapper := Swapper{}
myPointer := swapper.Init(&ex).
    Find("nested", func(p unsafe.Pointer) interface{} {
        return (*nestedFields)(p)
    }).
    Find("subPrivateFieldOne", func(p unsafe.Pointer) interface{} {
        return (*string)(p)
    }).Pointer()

// Assert type
str := myPointer.(*string)

// Reassign new value to pointer
*str = "new"

// swap.example{nested:swap.nestedFields{subPrivateFieldOne:"old", subPrivateFieldTwo:0, subPrivateFieldThree:0, subPrivateFieldFour:0, subPrivateFieldFive:false, subPrivateFieldSix:0, subPrivateFieldSeven:0}}
// swap.example{nested:swap.nestedFields{subPrivateFieldOne:"new", subPrivateFieldTwo:0, subPrivateFieldThree:0, subPrivateFieldFour:0, subPrivateFieldFive:false, subPrivateFieldSix:0, subPrivateFieldSeven:0}}

```