package main

import "fmt"

func main() {
    var y [5]int

    y[2] =20
    fmt.Println(y)

     var x [5]float64
     x[0] = 98
     x[1] = 93
     x[2] = 77
     x[3] = 82
     x[4] = 83 
     /*
       A shorter way to do this is to define the array this way
       x := [5]float64{ 98, 93, 77, 82, 83 } and it can also be broken into
       several lines as much as is needed:

       x := [5]float64{
            98,
            93,
            77,
            82,
            83,
        }
     */
     var total float64 = 0
     for i := 0; i < len(x); i++ {
           total += x[i]
     }
     fmt.Println(total / float64(len(x)))

     total = 0
     for _, value := range x {
        total += value
     }
     fmt.Println(total / float64(len(x)))
}