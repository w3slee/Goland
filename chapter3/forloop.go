package main

import "fmt"

func main(){
	i := 1
	for i <= 10 {
		if i%2 == 0{
			fmt.Println("even",i)
		} else if i % 2 != 0 {
			fmt.Println("odd",i)
		}
		i += 1
	}
}
