/*
	- An array is a numbered sequence of elements of a single type with
	  a fixed length.

	  var x [5]int 
	  x is an example of an array which is composed of 5 ints
*/
package main

import "fmt"

func main(){
	// 5 Elements -> index starts at 0
	var x [5]int
	// set the 5th element in the array to 100
	x[4] = 100
	// print the output
	fmt.Println(x)
	// 0, 0, 0, 0, 100
}
