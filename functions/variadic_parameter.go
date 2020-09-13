package main

import (
	"fmt"
)

// You can create a func which takes an unlimited number of arguments.
// Variadic parameter uses LEXICAL ELEMENT operator "...T"
// There, T specifies any Type

func main() {
	foo()
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func foo(x ...int) {
	fmt.Println(x)
	// TYPE of x is a SLICE of int
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("At index", i, "and corresponding value", v, ", cumulative sum is", sum)
	}

}
