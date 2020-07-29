/*
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

package main

import "fmt"

const (
	//a		observe error
	b     = 42
	c int = 55
)

func main() {
	a := "Hello"
	fmt.Println(a, b, c)
}
