package main

import "fmt"

func main() {
	// && - both conditions should be true
	// || - either of the conditions should be true
	// !  - not operator
	fmt.Println(true && true)
	fmt.Println(false && false)
	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println(!false)
}
