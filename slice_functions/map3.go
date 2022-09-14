package main

import (
    "fmt"
)

func main() {
    // create a map that contains chemical element names and state
    // create map of elements

    /*
        map[string]map[string]string
        A map of strings to maps of strings to strings
        The outer map is used as a lookup table based on the element's symbol
    */  
    elements := map[string]map[string]string{
        "H" : map[string]string{
            "name": "Hydrogen",
            "state": "Gas",
        },
        "He": map[string]string{
            "name": "Helium",
            "state": "Gas",
        },
    }
    var name string
    fmt.Scanf("%s", &name)
    if el, ok := elements[name]; ok {
        fmt.Println(el["name"], el["state"])
    } else {
        fmt.Printf("%s not found\n", el["name"])
    }
}
