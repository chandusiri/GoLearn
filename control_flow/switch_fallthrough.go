package main

import "fmt"

func main() {
	// Instead of multiple if else statements
	switch {
	case (2 == 3):
		fmt.Println(" 2 is always not equal to 3")
	case (2 == 2):
		fmt.Println(" 2 is always equal to 2")
		fallthrough
	case (2 > 3):
		fmt.Println("NO matter condition fails, prints bcoz of fallthrough")
		fallthrough
	case (2 < 3):
		fmt.Println("fallthrough is funky")
		fallthrough
	default:
		fmt.Println("Don't use fallthrough")
	}
}
