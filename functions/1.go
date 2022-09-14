package main

import "fmt"

// takes a slice of numbers and adds them together
func sum(numslice []int) int {
    total := 0
    for _, v := range numslice {
        total += v
    }
    return total
}

func main() {
    bg := []int{4, 3, 2}
    fmt.Println(sum(bg))
}
