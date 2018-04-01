package main

import "fmt"

func main() {
    for i := 0; i < 7; i++ {
        switch i {
            case 0: fmt.Println(i, "Zero")
            case 1: fmt.Println(i, "One")
            case 2: fmt.Println(i, "Two")
            case 3: fmt.Println(i, "Three")
            case 4: fmt.Println(i, "Four")
            case 5: fmt.Println(i,"Five")
            default: fmt.Println(i, "Unknown number")
        }
    }
}