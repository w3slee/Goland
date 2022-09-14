package main

import (
    "fmt"
    "sort"
)

func main() {
    x := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }
    sort.Ints(x) // sort smallest to largest
    fmt.Println("sorted array\n",x)
    fmt.Println("The smallest value in the array is : ", x[0])
}
   /* type ByAge []Person

        func (a ByAge) Len() int           { return len(a) }
        func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
        func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } */

