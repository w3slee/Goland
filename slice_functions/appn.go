package main

import (
    "fmt"
)

func main() {
    // slice1 with an array of 3 int elements 1,2,3
    slice1 := []int{1,2,3}
    // append 4,5 to slice1 
    slice2 := append(slice1, 4, 5)
    // print results
    fmt.Println("original value :",slice1,"\nupdated slice :", slice2)
}
