package main

import "fmt"

func main() {
	foo := myStruct{}
	foo.myField = "bar" // or get rid of this line and do this above
	// foo := myStruct{"bar"} or
	// foo := new(myStruct) , foo.myField = "bar" or
	// foo := &myStruct{"bar "}

	fmt.Println(foo.myField)

	cake := newMyStruct()
	cake.myMap["bar"] = "baz"

	fmt.Println(cake)
}

type myStruct struct {
	myField string
}

type myStruct2 struct {
	myMap map[string]string // remember maps have to be initialized before they csn be used
}

// create a constructor function to handle initialization of maps , syntax says it must start with `new`
func newMyStruct() *myStruct2 {
	result := myStruct2{}
	result.myMap = map[string]string{}

	return &result
}
