# go-swap

Testing utility for swapping out private field members without the need of dependency inversion of control.

### Example
```go
obj := example{
    nested: nestedFields{
        subPrivateField: "old",
    },
}
swapper := Swapper{}
myPointer := swapper.Init(&obj).
    Find("nested", func(p unsafe.Pointer) interface{} {
        return (*NestedFields)(p)
    }).
    Find("subPrivateFieldOne", func(p unsafe.Pointer) interface{} {
        return (*string)(p)
    }).Pointer()

// Assert type
str := myPointer.(*string)

// Reassign new value to pointer
*str = "new"
```