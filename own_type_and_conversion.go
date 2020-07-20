package main

import "fmt"

func main() {
	var a int
	// Creating your own TYPE with an underlying TYPE as int
	// Go lang sees these two TYPEs as different
	type donut int
	a = 42
	var b donut
	var c float32
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// try assinging a = b
	fmt.Printf("%T is different from %T\n", b, a)
	// Conversion from donut to underlying int
	// in Go, its called conversion, not casting
	a = int(b)
	fmt.Println(a)
	// Conversion from int to float
	c = float32(a)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
