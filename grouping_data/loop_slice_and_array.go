package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50}
	fmt.Println("length of x is:", len(x))

	// loop over the SLICE using RANGE
	for i, v := range x {
		fmt.Println(i, v)
	}

	// loop SLICE without RANGE
	for i := 0; i < len(x); i++ {
		fmt.Printf("%v\t", x[i])
	}
	fmt.Println()

	// Loop an ARRAY using RANGE
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array y is:", y)
	for i, v := range y {
		fmt.Println(i, v)
	}

}
