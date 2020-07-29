package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hello Go playground")
	}
	if !true {
		fmt.Println("Wont wana Print")
	}
	if !false {
		fmt.Println("Wana Print yay")
	}
	// Scope of x is limited to only if block
	if x := 1; x < 2 {
		fmt.Println("Initialization in if works!")
	}
	// fmt.Println(x)
	// if else if and else
	a := 42
	if a < 40 {
		fmt.Println("a was less than 42")
	} else if a > 42 {
		fmt.Println("a was greater than 42")
	} else {
		fmt.Println("a was 42")
	}

}
