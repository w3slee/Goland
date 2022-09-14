package main

import "fmt"

func main(){
	// cant assign to bx since it declared as a constant
	// program will not compile

	const bx string = "AA313-131e1-123"
	fmt.Println(bx)
	bx = "@@!22e1"

}
