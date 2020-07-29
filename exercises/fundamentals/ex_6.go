/*
Write a program that prints a number in decimal, binary, and hex
*/
package main

import "fmt"

func main() {
	number := 100
	fmt.Printf("%d\n", number)
	fmt.Printf("%b\n", number)
	fmt.Printf("%#x\n", number)
}
