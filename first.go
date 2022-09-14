// package declaration - every go program must start with thisAA
// packages are go's way of organizing and reusing code.
// types of go programs : executables and libraries
// 
package main
// the import keyword is how we include code from other packages
// fmt package implements formatting for input and output
import "fmt"

// this is a comment
// function declaration

func main(){
	fmt.Println("Go is running ...")
	// a space is also considered as a character
	var name string  = "mch 1nen"
	fmt.Println("Hello, " + name)
	fmt.Println("Length of name :", len(name))
}
