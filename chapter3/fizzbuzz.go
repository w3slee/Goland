/*
 	A program that prints the numbers from 1 to 100
	But for multiples of 3 print Fizz
	    for multiples of 5 print Buzz
*/
package main

import "fmt"

func main(){
	for i :=1; i <= 100; i++{
		if i % 3 == 0{
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
