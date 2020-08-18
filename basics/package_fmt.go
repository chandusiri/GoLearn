package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("This is to print a line and capture its return values n and err")
	// n -> number of bytes writen
	fmt.Println(n)
	//err -> any error if present.
	fmt.Println(err)
	//To avoid catching any var lets say err
	n, _ = fmt.Println("_ Tells go compiler to ignore the variable thats not being used in program")
	fmt.Println(n)
	// try catching both n, e and not print or use e and observe ouptut - e declared and not used

}
