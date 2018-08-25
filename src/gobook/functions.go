package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
	value, exists := power("goku")
	if exists == true {
		// handle this error case
		fmt.Printf("Goku exists and his power is %d\n", value)
	}
}

// no return type defined
func log(message string) {
}

// return type specified as integer
func add(a int, b int) int {
	return 9
}

// return type specified as both integer and boolean
func power(name string) (int, bool) {
	return 9, true
}
