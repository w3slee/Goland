/* 
    A function which takes an integer and halves it and returns true
    if it was even or false if it was odd

*/
package main

import "fmt"

// param : number int
func half(number int)(int){
    fmt.Println("xxxx")
    if (number / 2 % 2) == 0 {
        fmt.Println(number)
    }
}
func main() {
    half(20)
}
