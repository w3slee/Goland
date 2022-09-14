package main

import (
    "fmt"
)

func main() {
    slice1 := []int{1,2,3}
    // create slice that hold up to 2 elements max
    slice2 := make([]int, 2)
    // copy slice1 into slice2 according to the size specified
    copy(slice1, slice2)
    // since slice1 has 3 elements in it and slice2 has a storage space for only 2 elements
    // only two elements will be copied
    fmt.Println("slice1[]:len ",slice1, len(slice1),"\nslice2[]:len ", slice2, len(slice2))

}
