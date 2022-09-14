// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    elements := make(map[string]string)
    elements["He"] = "Helium"
    elements["H"] = "Hydrogen"
    elements["Li"] = "Lithium"
    elements["O"] = "Oxygen"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["B"] = "Boron"
    elements["N"] = "Neon"
    elements["F"] = "Fluorine"
    elements["Be"] = "Beryllium"
    // Technically a map returns zero value for the value type (which for strings is the empty string)
    // Althought we could check for the zero value in a condition.
    // Go provides a better way
    var chemical string 
    fmt.Println("Chemical Query Repository\nPlease insert symbol to search")
    fmt.Scanf("%s", &chemical)
    name, ok := elements[chemical]

    if ok != true {
        fmt.Println("element not found")
    } else{
        fmt.Printf("Element(%s) -> %s found\n",chemical, name)
    }
}
