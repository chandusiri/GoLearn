package main

import "fmt"

func main() {
	fmt.Println("Entry goes with this")
	foo()
	fmt.Println("something more")

	// using loop
	fmt.Println("Lets print even numbers from 0 to 10")
	for i := 0; i < 10; i++ {
		// using if
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("Function foo")
}

func bar() {
	fmt.Println("Exit the bar")
}

// control flow
// 1. sequence
// 2. loop
// 3. conditional
