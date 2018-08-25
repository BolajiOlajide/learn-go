package main

import "fmt"

func main() {
	var aanu Saiyan
	goku := &Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	bolaji := &Saiyan{Name: "Bolaji"} // valid
	aanu = Saiyan{"Aanu", 3000, goku} // this should only be done for structures with few fields
	fmt.Println(goku)
	fmt.Println(bolaji)
	fmt.Println(aanu)
	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: bolaji,
		},
	}
	fmt.Println(gohan)
}

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}
