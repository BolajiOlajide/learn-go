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
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(fifth)
	fmt.Println("")

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

	// arithmetic operator
	add := 1 + 2
	fmt.Println(add)

	// subtraction operator
	subtract := 10 - 5
	fmt.Println(subtract)

	// remainder (modulus operator)
	remainder := 7 % 3
	fmt.Println(remainder)

	// division operator
	divide := 10 / 5
	fmt.Println(divide)

	// multiplication operator
	multiply := 4 * 5
	fmt.Println(multiply)

	// incrememnt operator
	fmt.Println("incrememting stuff!")
	inc := 1
	inc++
	fmt.Println(inc)
	inc++
	fmt.Println(inc)
	inc += 5
	fmt.Println(inc)
}

func init() {
	name = "Bolaji"
}
