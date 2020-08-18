package main

import "fmt"

func main() {
	var y = 14
	// Prints without newline
	fmt.Print(y)
	// Prints with newline
	fmt.Println("This is value of y")
	// Format printing. Formats the value of y
	// Print VALUE of y
	fmt.Printf("value of y is %v\n", y)
	// Print TYPE of y
	fmt.Printf("%T\n", y)
	// print y in # - go representation of hexa decimal format and X - caps in hex A-F
	fmt.Printf("%#X\n", y)
	// Print y in go representation lower case hex and binary and normal lower case hex
	// Multiple variables printing in same line
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// sprint is print to a string
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	// Fprint methods are to print to a file
}
