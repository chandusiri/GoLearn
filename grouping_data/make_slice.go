package main

import "fmt"

// MAKE - MAKE the SLICE with an estimated CAPacity
// Reduces runtime effort to re create a new SLICE everytime you append element
func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// x[10] = 3456 -> observe error
	x = append(x, 3443, 3444)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// When you append more than CAP, it recreates the SLICE
	// Observe the CAP of newly created SLICE. Its double than the previous CAP
	x = append(x, 3445)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
