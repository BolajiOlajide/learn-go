// Structures don’t have constructors.
// Instead, you create a function that returns an instance of
// the desired type (like a factory)

package main

// NewSaiyan - Factory method to create new Saiyan instance
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}
