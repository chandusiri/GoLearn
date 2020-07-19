package main

import "fmt"

// Programming Terminology: Declare and assign = initialization
// var keyword is to declare any variable outside of function body
// Scope of this variable would be of package level
var y = 43

// Assigns "The zero value" of TYPE int to "z" by default
// The zero value is false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// Short declaration variable for initialization
	// scope is statements from declaration till end of function body
	// Best practice is to use variables with lesser scope
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("again", y)
}
