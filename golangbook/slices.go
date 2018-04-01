package main

import "fmt"

func main() {
    /*
    A slice is a segment of an array. 
    Like arrays slices are indexable and have a length. 
    Unlike arrays this length is allowed to change.
    */

    var x = make([]float64, 5, 10)

    /*
    the first argument specifies the length of the slice
    the second specifies the capacity of the slice
    */

    // If you intend to create a slice yourself you should use the 
    // inbuilt make function 👆👆👆
    fmt.Println(x)
    fmt.Println(len(x))
    fmt.Println(cap(x))

    // another way to create a slice is to create it from an 
    // existing array
    arr := [10]int{0,1,2,3,4,5,6,7,8,9}
    y := arr[:]
    fmt.Println(y)
    y[0] = 100
    fmt.Println(arr)

    // The append method creates a new slice from an exsiting slice
    // and adds more items to it

    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)

    // The copy command is used to copy the content of a slice into another
    slice3 := []int{1,2,3}
    slice4 := make([]int, 2)
    copy(slice4, slice3)
    fmt.Println(slice3, slice4)
    // it takes the copy up to the it's capacity and discard any leftover 👆
}