package main

import "fmt"

/*
	A program that prints out all the numbers evenly divisible by 3
	between 1 and 100
*/
func main(){
	for i := 1; i <= 100; i++ {
		if i % 3 == 0{
			fmt.Println(i, "is divisible by 3")	
		}
	}
}
