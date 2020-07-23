package main

import "fmt"

const a = 4

const (
	c string  = `Hey! You there?`
	d float32 = 5.4
)

func main() {
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	// Reassign VARIABLE holding CONSTANT
	// a = 5 + 5
	a := 5 + 5
	fmt.Println(a)
}
