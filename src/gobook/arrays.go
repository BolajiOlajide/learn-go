package main

import "fmt"

func main() {
	// creating an array
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	// creating a slice
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	slice2 := make([]int, 10) // creates a slice with a capacity of 10 and length of 10 initial items
	fmt.Println(slice2)
	slice3 := make([]int, 3, 10) // creates a slice with a cpacoty of 10 but a length of 3
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	scores := make([]int, 2, 10)
	scores[0] = 9033
	fmt.Println(scores)

	// to expand a slice use the append command
	scores = append(scores, 5)
	fmt.Println(scores)
}
