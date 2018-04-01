package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println(output)

    fmt.Print("Enter degree in celsius to convert: ")
    fmt.Scanf("%f", &input)
    convertedDegree := (input - 32) * 5/9
    fmt.Println(convertedDegree)
}