package main

import (
	"fmt"
	"io"
)

func main() {
	var input int
	fmt.Print("Please enter a digit: ")
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("no more input!")
	} else {
		fmt.Printf("The user just entered %d\n", input)
	}
}
