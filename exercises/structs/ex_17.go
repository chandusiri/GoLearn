package main

/*
Create and use an anonymous struct.
*/

import (
	"fmt"
)

func main() {
	a := struct {
		fruits  []string
		healthy bool
	}{
		fruits:  []string{"Guava", "Banana"},
		healthy: true,
	}
	fmt.Println(a)
}
