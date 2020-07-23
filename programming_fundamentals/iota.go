package main

import "fmt"

// Within a constant declaration, the predeclared identifier iota
// represents successive untyped integer constants.
// Its value is the index of the respective ConstSpec in that constant declaration, starting at zero.
const (
	a = iota
	b = iota
	c = iota
	x = 10
	y
	z
)

// iota resets to 0 whenever the reserved word const appears in the source
// It can be used to construct a set of related constants
// Dont need write iota again and again if increment is needed
const (
	d = iota
	e
	f
	g
)

func main() {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
