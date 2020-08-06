package main

import "fmt"

// Array is a numbered sequence (zero based indexing) of elements of single TYPE
// The length is part of the array's type. eg var x [5]int
// Use slice instead

func main() {
	// The length is part of the array's type.
	// Means x and y are two different TYPEs
	var x [5]int
	var y [6]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
