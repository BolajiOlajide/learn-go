package main

import (
	"errors"
	"fmt"
)

func main() {
	de := process(2)
	fmt.Println(de)
	fmt.Println("YEah!")
}

func process(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}
	return nil
}
