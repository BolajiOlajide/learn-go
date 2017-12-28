package main

import "fmt"

func main() {
	foo := 10

	if foo == 1 {
		fmt.Println("bar!")
	}

	if foo > 5 {
		fmt.Println("Greater")
	} else {
		fmt.Println("Lesser!")
	}

	// another way of writing if statements by linking the initializer to the branch
	if foo := 2; foo > 1 {
		fmt.Println("Greater than One!")
	}

	// switch statement
	switch foo := 2; foo {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// refactoring
	base := 3

	switch {
	case base == 1:
		fmt.Println("One")
	case base > 2:
		fmt.Println("Greater than two!")
	}
}
