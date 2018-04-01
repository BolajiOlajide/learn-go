package main

import "fmt"

func main() {
    /*
    A map is an unordered collection of key-value pairs.

    Also known as an associative array, a hash table or a dictionary, 
    maps are used to look up a value by its as- sociated key.

    var x map[string]int
    */

    // var x map[string]int
    // x["key"] = 10
    // fmt.Println(x)

    // 👆 This can't work because maps have to be initialized before use
    x := make(map[string]int)
    x["key"] = 10
    x["next"] = 50
    delete(x, "key") // to delete a value in a map by using it's key
    fmt.Println(x)

    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    fmt.Println(elements)
    name, ok := elements["Li"]
    fmt.Println(name, ok)
    name1, ok1 := elements["Uun"]
    fmt.Println(name1, ok1)
}