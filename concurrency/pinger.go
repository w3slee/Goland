package main

import (
    "fmt"
    "time"
)

func pinger(c chan int) {
    for i := 0; ; i++ {
        c <- i
    }
}

func ponger(c chan int) {
    for i := 0; ;i++ {
        c <- i
        i += 2
    }
}

func printer(c chan int){
    for {
        msg := <- c 
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    var c chan int = make(chan int)

    go pinger(c)
    go ponger(c)
    go printer(c)

    var input int
    fmt.Scanln(&input)
}


