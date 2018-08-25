// Maps grow dynamically. However, we can supply a second
// argument to make to set an initial size:
// lookup := make(map[string]int, 100)
package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	lookup["gohan"] = 7001

	fmt.Println(lookup)

	power, exists := lookup["goku"]
	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)

	total := len(lookup) // returns the amount of keys in the map
	fmt.Printf("There are %d keys in the lookup map.\n", total)

	gohan := &Saiyan{Name: "Gohan"}
	var friends map[string]*Saiyan
	friends = make(map[string]*Saiyan)
	friends["gohan"] = gohan

	goku := &Saiyan{
		Name:    "goku",
		Friends: friends,
	}
	fmt.Println(goku)

	// another way of declaring a map
	lookups := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}
	fmt.Println(lookups)

	// We can iterate over a map using a for loop combined with the range keyword
	for key, value := range lookups {
		fmt.Printf("%s's power is at level %d\n", key, value)
	}
}

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}
