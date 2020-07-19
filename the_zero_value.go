package main

import "fmt"

func main() {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable
	var y string
	var z int
	// var a
	fmt.Println(y, "This prints nothing")
	fmt.Println("empty string", y, "!observe this line clearly")
	fmt.Printf("y is of type %T\n", y)
	y = "Professor"
	fmt.Println("y is", y)
	fmt.Println(z, "is default value of type int")
	fmt.Printf("z is of type %T", z)
}
