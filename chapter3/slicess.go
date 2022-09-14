/*
	- A slice is a segment of an array
	- Slices are indexable and have a length
	- The only diff btw a slice and an array is the missing length between
	  the brackets
*/

// example of a slice
// in this case x has been created with a length of 0
var x[]float64
// if you want to create a slice you should use the built-in make function
x := make([]float64,5)
// this creates a slice that is associated with an underlying float64 array
// of length 5
// slices are always associated with some array
// slices can be smaller but never longer than the array
// the make function allows a 3rd parameter
x := make([]float64, 5, 10)
// 10 represents the capacity of the underlying array which the slice points to

/*
	- Another way to create slices is to use the [low : high] expression
	- low is the index of where to start the slice
	- high is the index of where to end it ( but not including he index itself)
	*/
arr := [5]float64{1,2,3,4,5}
x := arr[0:5]
