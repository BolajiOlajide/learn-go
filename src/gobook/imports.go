package main

import (
	"fmt"
	"os"
)

func main() {
	i := []string{"foo", "bar", "buz"}
	fmt.Println(len(i))
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}
