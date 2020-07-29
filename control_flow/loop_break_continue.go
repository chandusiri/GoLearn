package main

import "fmt"

func main() {
	x := 0
	for {
		x++

		// Break the loop at some point (here when x becomes > 10)
		if x > 10 {
			break
		}

		// Skip printing odd numbers
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}
}
