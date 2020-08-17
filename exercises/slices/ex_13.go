/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.


*/
package main

import "fmt"

func main() {
	x := [][]string{{}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(x)
	x[0] = append(x[0], []string{"James", "Bond", "Shaken, not stirred"}...)

	for _, v := range x {
		for _, s := range v {
			fmt.Println(s)
		}
	}
}
