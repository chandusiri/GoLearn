package main

import (
	"fmt"
)

// composite data STRUCTure, hence the name struct.
// Struct is a composite data type (aka aggregate data types, aka complex data types)
// STRUCT allows us to compose together values of different TYPE

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Siri Chandana",
		last:  "Anumula",
		age:   28,
	}
	p2 := person{
		first: "Roop Kumar",
		last:  "Bhattaram",
		age:   29,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, "is", p1.age)
}
