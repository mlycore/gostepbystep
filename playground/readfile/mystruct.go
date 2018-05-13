package main

type MyStruct struct {
	N int
}
var n MyStruct
//n := MyStruct{ 1 }
// get
immutable := reflect.ValueOf(n)
val := immutable.FieldByName("N").Int()
fmt.Printf("N=%d\n", val) // prints 1

// set
mutable := reflect.ValueOf(&n).Elem()
mutable.FieldByName("N").SetInt(7)
fmt.Printf("N=%d\n", n.N) // prints 7
