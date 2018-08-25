// Why do we want a pointer to the value,
// rather than the actual value?
// It comes down to the way Go passes arguments
// to a function: as copies.

package main

import "fmt"

func main() {
	goku := &Saiyan{"Goku", 9000} // the & operator gets the address of our value (it’s called the address of operator).
	Super(goku)
	fmt.Println(goku)
	gohan := &Saiyan{"Gohan", 1000}
	gohan.Super()
	fmt.Println(gohan)
}

// Saiyan - structure for creating saiyans
type Saiyan struct {
	Name  string
	Power int
}

// Super - method for incremented saiyan's power
func Super(s *Saiyan) {
	s.Power += 10000
	// s.Name = "Gohan"
	// s = &Saiyan{"Gohan", 1000}
}

// Super - associated struct method
func (s *Saiyan) Super() {
	s.Power += 10000
}
