package main

import (
	"fmt"
)

// VAR, TYPE and CONST are keywords.
// See similarity of var and type declaration

// var identifier underlyingtype
var x int

// type identifier underlyingtype
type person struct {
	first string
	last  string
}

// Aliasing a TYPE, we can do that builtin types like int.
// Its not a good practice to alias types.
type foo int

var y int

const bar int = 42
const beer = 43

func main() {
	fmt.Println("Hello, playground")
	// Here compiler does implicit conversion of y to the anonymous type beer
	y = beer
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))
	// Try doing y = bar and notice
	// You cannot assign a named type to a different named type
}
