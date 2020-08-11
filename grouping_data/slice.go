package main

import "fmt"

/*
	// SLICE allows you to group together VALUES of the same TYPE
	// COMPOSITE LITERAL used to construct SLICE
	// Syntax: [] type{values}
*/

func main() {

	x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	//x[6] = 10 - panic: runtime error: index out of range [6] with length 6

	// SLICE a SLICE
	fmt.Println(x[2:5])
	x = x[:]
	fmt.Println("x is", x)

	// APPEND to a SLICE
	// func append(slice []T, elements ...T) []T
	x = append(x, 111, 222, 333)
	fmt.Println("x after append is:", x)

	y := append(x[6:], 1024, 2048)
	fmt.Println("y is", y)
	fmt.Println("len of y is:", len(y))
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", x)

	z := []int{55, 555, 5555}
	// This again proves Go is all about TYPE
	// z = append(z, y) cannot use y (type []int) as type int in append
	// z = append(z, y...) ... specifies append whole bunch slice y to z
	z = append(z, y[0:3]...)
	fmt.Println("z is:", z)

	// Delete from a SLICE
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	// Initial CAPACITY of SLICE = LENGTH of SLICE
	n := []int{1, 2, 3, 5, 6}
	fmt.Println(cap(n))
	// After APPEND, CAP is double the initial LEN
	n = append(n, 7, 8)
	fmt.Println(cap(n))
}
