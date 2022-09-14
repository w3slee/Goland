/*
    defer is mostly used when some resources need to be freed in some way
*/
package main

import "fmt"

func first() {
    fmt.Println("first")
}
func second() {
    fmt.Println("second")
}
func main() {
    defer first()
    second()
}




