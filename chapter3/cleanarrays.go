package main

import "fmt"

func main(){
	var x [5]float64
	x[0] = 98
	x[1] = 90
	x[2] = 33
	x[3] = 75
	x[4] = 67

	var total float64 = 0
	// we don't use i so the compiler will throw an erro
	// for this to work change i to _
	// _ tells the compiler we don't need i, the iterator variable
	for _, value := range x {
		total += value
	}
	fmt.Println("Average: ", total/float64(len(x)))
}
