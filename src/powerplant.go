package main

import (
	"fmt"
)

func main() {
	var plantCapacities []float64
	plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	// slice not an array
	activePlants := []int{0, 1}

	gridLoad := 75.

	fmt.Println("1) Generate Power plant report")
	fmt.Println("2) Generate Power Grid report")
	fmt.Println("Please choose an option: ")

	var option string

	fmt.Scanln(&option)
	// the ampersand here is used to link the address of the option
	// entered by the user to the variable. If the ampersand wasnt there
	// only a copy would be passed and one wouldn't get access to what the
	// user entered

	println(option, "Your option!")

	switch option {
	case "1":
		for idx, cap := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
		}
	case "2":
		capacity := 0.
		for _, plantID := range activePlants {
			capacity += plantCapacities[plantID]
		}

		fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
		fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
		fmt.Printf("%-20s%.1f%%\n", "Utilization: ", (gridLoad/capacity)*100)
	default:
		fmt.Println("Invalid value entered!")
	}

}
