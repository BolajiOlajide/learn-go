package main

import (
	"fmt"
)

func main() {
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}

		numTerms = len(terms)
		return
	}

	numTerms, sum := addFunc(12, 12)
	fmt.Println("Added:", numTerms, "terms for a total of", sum)
}
