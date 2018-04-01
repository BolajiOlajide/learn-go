package main

import "fmt"
import "os"

// this is a comment

func main() {
    fmt.Println("Hello world")
    fmt.Println("Just fine 😉")
    fmt.Println("1 + 1 =", 1.0 + 1.0)
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
    fmt.Println("The result is ", 321325 * 424521)
    fmt.Println(len("Bolaji"))
    fmt.Println((true && false) || (false && true) || !(false && false))
    fmt.Println(false && false)
    os.Exit(0)
    fmt.Println("Happy")
}

/*
This is another way of writing commments
In this comment style, you can write a comment spanning multiple lines.
*/