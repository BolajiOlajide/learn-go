package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(8) // tell go to use 8CPUs for the process = to allow for parallelism
	go abcGen()

	fmt.Println("This comes first!")

	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		go fmt.Println(string(l))
	}
}
