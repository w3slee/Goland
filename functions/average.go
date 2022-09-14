package main

import "fmt"

func average(xs []float64) float64 {
    total := 0.0
    // for i in value range of xs 
    for _, v := range xs {
        // total = current total + current value
        total += v
    }
    // return statement causes the function to immediately stop and return
    // the value after it to the function that called this one
    return total / float64(len(xs))
    // causes a runtime error
    // panic("Not implemented")
}

func main() {
    pb := []float64{32, 43, 64, 31, 21, 98, 1, 79}
    fmt.Println(average(pb))
}
