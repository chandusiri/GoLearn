package main

import (
	"fmt"
)

// Go uses UTF8 coding scheme.
// In UTF8, a character can correspond from one byte to 4 bytes (32 bits).
// Those bytes correspond to hexa decimal numeral system in UTF8
// Hex is base 16 numbering from 1 to 10 and a-f/A-F
// 1 byte = uint8 (unsigned(only positive) 8 bit int)
// rune = int32 (32 bits = 4 bytes)
// so each rune is a code point in UTF8

// String is sequence of bytes or slice of bytes
// String and slice of bytes - %s, %q, %x, %X

// resource to refer Strings bytes runes and characters in Go is
// https://blog.golang.org/strings
func main() {
	s := "Hello playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Printf("%s\n", s) //uninterpreted string
	fmt.Printf("%q\n", s) //double quotes - interpreted string
	fmt.Printf("%x\n", s) //16 bit, lower case
	fmt.Printf("%X\n", s) //16 bit, upper case

	fmt.Printf("---%x\n", "世")

	// As string is slice of bytes, convert TYPE string to slice of bytes
	// This prints decimal value for each letter of s
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// Print UTF8 format for each character of s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	// For better understanding, see printable characters in ascii wiki page
	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

}
