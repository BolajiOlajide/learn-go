package main

import "fmt"

func main() {
    arr := make([]int, 3, 9)
    fmt.Println(arr)

    x := [6]string{"a","b","c","d","e","f"}
    res := x[2:5]
    fmt.Println(res)

    plenty := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }

    smallest := plenty[0]
    for _, value := range plenty {
        if (value < smallest) {
            smallest = value
        }
    }
    fmt.Println(smallest)
}