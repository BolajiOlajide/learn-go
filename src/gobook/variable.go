package main

import "fmt"

func main() {
	const power int = 9000
	var power2 int
	power2 = 900
	power3 := getPower()
	fmt.Printf("It's over \npower = %d\npower2 = %d\npower3 = %d\n", power, power2, power3)

	name, power4 := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power4)
}

func getPower() int {
	return 9001
}
