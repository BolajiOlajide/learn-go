package main

import "fmt"

func main() {
	sayHello()
	var name string

	fmt.Println("Please enter your name")
	fmt.Scanln(&name)

	sayMyName(name)

	// sending parameter by reference
	slogan := "Send me money!"
	manipulateString(&slogan)

	fmt.Println(slogan)

	sayWords("Hello", "Go", "from", "the", "other", "side")

	numTerms, sum := add(1, 2)
	fmt.Println("Added:", numTerms, "terms for a total of", sum)

	a, b := add2(5, 15, 25, 35, 50)
	fmt.Println("Added:", a, "terms for a total of", b)
}

func sayHello() {
	fmt.Println("Hello!")
}

func sayMyName(name string) {
	fmt.Println("Hello, My name is", name)
}

func manipulateString(slogan *string) {
	fmt.Println(*slogan)
	*slogan = "Send me money now!"

	// this is the way we can manipulate a parameter in Go!
}

// variadic function
func sayWords(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}

// named return value
func add2(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}

	numTerms = len(terms)
	return
}
