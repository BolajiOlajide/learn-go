package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	// fmt.Println(scores)
	sort.Ints(scores)
	worst := make([]int, 5)
	copy(worst[2:4], scores[:10])
	fmt.Println(worst)
	worst2 := make([]int, 5)
	copy(worst2, scores[:5])
	fmt.Println(worst2)
	just := scores[:10]
	fmt.Println(just)
}
