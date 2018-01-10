package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// result := [][]uint8{}
	result := make([][]uint8, dy)
	var j int

	for i := range result {
		for j = 0; j < dx; j++ {
			result[i] = append(result[i], uint8((j+i)/2))
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
