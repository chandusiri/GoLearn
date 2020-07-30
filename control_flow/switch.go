package main

import "fmt"

func main() {
	// If switch has no expressions, it evaluates to true.
	// So you can use it as multiple if else statements
	switch {
	case (2 == 3):
		fmt.Println(" 2 is always not equal to 3")
	case (2 == 2):
		fmt.Println(" 2 is always equal to 2")
		// As condition is true and no fallthrough, switch exits here.
	case (2 < 3):
		fmt.Println("Though also 2<3, wont print bcoz there is no fallthrough")
		// fallthrough here wont work because compiler will not come to execute this case at all
	default:
		fmt.Println("If none of above cases are true, then this will get printed")
	}
}
