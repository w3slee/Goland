/* 

    This program consists of two goroutines. 
    - The first goroutine is implicit and is the main function itself. 
    - The second goroutine is created when we call go f(0)
*/
package main

import (
    "fmt"
    "time"
    "math/rand"
)
func f(n int){
    for i := 0; i < 10; i++ {
        fmt.Println("n: ",i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

// modified to run 10 go routines
func main() {
    for i := 0; i < 10; i++ {
        go f(i)
    }
    //go f(0)
    var input string
    fmt.Scanln(&input)
}
