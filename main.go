package main

import "fmt"

var name string

const (
	first = iota
	second
	third
)

const (
	fourth = 1 << (10 * iota)
	fifth
)

func main() {
	fmt.Println("Hello " + name)

	myArray := [3]int{} // an array with a fixed size
	myArray[0] = 42
	myArray[1] = 20
	myArray[2] = 18

	fmt.Println(myArray)

	myArray2 := [...]int{58, 90, 12, 1, 43, 98} // an array with no fixed size
	fmt.Println(myArray2)
	fmt.Println(len(myArray2)) // Get the size of an array

	// Slices are a better fit than arrays (known as Collections) in Go because multiple operations can be performed on them
	// but sometimes it can get really messy and complicated
	mySlice := []float32{14., 15., 16.}
	mySlice = append(mySlice, 17.2)
	fmt.Println(mySlice)

	// maps are equivalent of dictionaries (or Object)
	myMap := make(map[string]string)

	myMap["name"] = "Bolaji"
	myMap["location"] = "Lagos"

	fmt.Println(myMap)
}

func init() {
	name = "Bolaji"
}
