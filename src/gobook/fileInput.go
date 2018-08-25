package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("a_file_to_read")
	// file, err := os.Open("input.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// read the file
}
