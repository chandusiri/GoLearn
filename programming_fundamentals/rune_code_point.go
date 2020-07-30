/*
Print every rune code point of the uppercase alphabet. Your output should look like this:
65	U+0041 'A'
66	U+0042 'B'
 … through the rest of the alphabet characters
*/
package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
