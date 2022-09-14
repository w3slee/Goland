/*
	&& and
	|| or
	!  not

	Truth Table
	-----------
	
	Expression	Value
	----------	-----
	true && true	true
	true && false	false
	false && true	false
	false && false	false

	true  || true	true
	true  || false	true
	false || true	true
	false || false	false

	!true		false
	!false		true
	*/
package main

import "fmt"

func main(){
	// true
	fmt.Println(true && true)
	// false
	fmt.Println(true && false)
	// true
	fmt.Println(true || true)
	// true
	fmt.Println(true || false)
	// false
	fmt.Println(!true)
}
