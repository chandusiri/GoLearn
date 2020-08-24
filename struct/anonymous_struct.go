package main

import (
	"fmt"
)

func main() {
	// Instead of giving name to struct Type as Person, use it in oneliner if needed.
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p)
}
