/*  
    A function with one variadic parameter that finds the greatest number in
    the list of numbers    
*/ 
package main

import (
    "fmt"
    "sort"
)


func greatest(numbers ...int) []int {
    sort.Ints(numbers)
    return numbers
    }

func main() {
    bp := greatest(420,911,411,911)
    fmt.Println(bp[0])
}
