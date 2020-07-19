package main

import "fmt"

func main() {
	// := is short declaration operator which declares and assigns a variable for the first time
	x := 10
	fmt.Println("initial value of x is", x)
	// For reassignment, you need just = operator
	x = 100
	y := 10 + 100
	fmt.Println(x, y)
}
