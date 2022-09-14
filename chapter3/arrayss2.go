
package main

import "fmt"

func main(){
	var x [5]float64
	x[0] = 42
	x[1] = 22
	x[2] = 65
	x[3] = 90
	x[4] = 55
	// total variable
	var total float64 = 0
	// while i < 5 , add values at each index upto 5 to total
	for i:=0; i < 5; i++ {
		total += x[i]
	}
	// average
	fmt.Println("Total : ",total)
	// better to use length of array
	fmt.Println("Average : ",total / float64(len(x)))
}
