package main

import "fmt"

func sq(x *float64) {
    *x = *x * *x
}

func main() {
    x := 1.5
    sq(&x)
    fmt.Println(x)
}
