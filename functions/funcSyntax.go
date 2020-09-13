package main

import (
	"fmt"
)

// func (r receiver) functionName(parameters) (returns) {...}
// Note: Difference between terms "parameters" and "arguments"
//		1. We define a func with parameters (if any)
//		2. We call a func with arguments (if any)
func main() {
	foo()
	bar("customer")
	wooSays := woo("James", "Bond")
	fmt.Println(wooSays)
	micSays, truth := mic("Ian", "Fleming")
	fmt.Println(micSays)
	fmt.Println(truth)
}

// function with no parameters and no return value
func foo() {
	fmt.Println("Hello foo!")
}

// function with parameter and no return
// everything in Go is PASS BY VALUE
func bar(cus string) {
	fmt.Println("Hello", cus)
}

// func with multi parameters and a return value
func woo(fn, ln string) string {
	// print to a string
	return fmt.Sprint("Hello ", fn, " ", ln)
}

//func with multi parameters and multiple return values
func mic(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, " says Greeting puts a smile on the face")
	b := true
	return a, b
}
