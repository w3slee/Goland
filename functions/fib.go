package main

//import "fmt"

func fibb(n int) int {
    fibb(n - 1) + fibb(n - 2)
    return n
}

func main() {
    fibb(22)
}
