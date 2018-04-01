package main

import "fmt"

/*
Collectively the parameters and the return type are known as the function's signature.
*/


func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
    fmt.Println(f2())
    x,y := f()
    fmt.Println(x,y)
    fmt.Println(add(1,2,3))
    arr := []int{20,30,40,50,80} 
    // i think you can't pass arrays to variadic functions
    // just slices
    fmt.Println(add(arr...))
}

// named return type
func f2() (r int) {
    r = 1
    return 
}

// Go also allows to return multiple values
func f() (int, int) {
    return 5, 6
}

// variadic functions
func add(args ...int) int {
    total := 0
    for _, value := range args {
        total += value
    }
    return total
}