package main

import "fmt"

func main() {
    j := 1
    for j <= 10 {
        fmt.Println(j)
        j++
    }

    fmt.Println("The second type of for loop")

    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}