package main

import "fmt"

func main() {
    // map of key string to int value
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x["key"])

    // map of key type int to int value
    y := make(map[int]int)
    y[1] = 1000
    y[5] = 5000
    for i := 0; i <= len(y); i++ {
        fmt.Println(y.keys)
    }
    fmt.Println(y[1])
}
