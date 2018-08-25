// A goroutine is similar to a thread, but it is scheduled by Go,
// not the OS. Code that runs in a goroutine can run
// concurrently with other code.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}
