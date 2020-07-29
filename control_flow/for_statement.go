package main

import "fmt"

func main() {
	a := 1
	// There is no "while" in Go
	// For statements with a single conditions
	// Like a C lang while loop
	for a < 6 {
		fmt.Println(a)
		a++
	}
	fmt.Println("Done.")

	// eternal loop.
	// Like a C lang for(;;) loop
	x := 1
	for {
		if x > 5 {
			break
		}
		fmt.Printf("%v\t", x)
		x++
	}

}
