package main

import "fmt"

func main() {
    // we call this function after panic() has been called
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    // runs before func
    panic("Panic")
}
