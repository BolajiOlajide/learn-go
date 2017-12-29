package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Refactored for loop!")

	// a non-efficient way (kinda) of writing a for loop
	i := 0
	for {
		i++
		println(i)

		if i > 5 {
			break
		}
	}

	// looping through an array with the range keyword

	s := []string{"foo", "bar", "buz"}

	for idx, v := range s {
		fmt.Println(idx, v)
	}

	// looping through a map (Object!)

	m := make(map[string]string)

	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "buz"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
