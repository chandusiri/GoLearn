package main

import "fmt"

// go is a STATIC programming language, not dynamic
// a VARIABLE is DECLARED to hold a VALUE of certain TYPE
// There are primitive data types and composite data types.
// check them in golang.or -> language specifications
var y = 50

// DECLARE the VARIABLE is of IDENTIFIER "z" of TYPE string.
// This is interpreted string
var z string = "Shaken, not stirred"

// raw string
var r = `James said,
"Shaken,

not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
	fmt.Println(z)
	fmt.Printf("z is of type %T\n", z)
	// uncomment below snippet and observe error
	// z = 43
	// fmt.Println(z)
	fmt.Println(r)
	fmt.Printf("r is of type %T\n", r)
}
