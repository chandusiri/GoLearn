package main

import (
	"fmt"
)

// struct can be embedded in another struct

type person struct {
	first string
	last  string
	age   int
}

// Inside struct are fields
// person is an Anonymous field or Embedded field.
// "Unqualified TYPE name" person is the field name here

type secretAgent struct {
	person
	license bool
	last    string
}

func main() {

	// field values are initialized this way
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   30,
		},
		license: true,
		last:    "gov",
	}
	fmt.Println(sa1)
	// inner TYPE person gets promoted to outer TYPE SECRE
	// So you can use sa1.first instead of sa1.person.first if there is no other TYPE first in sa1
	fmt.Println(sa1.first, "has license to kill")
	fmt.Println(sa1.person.first, "is same as", sa1.first)
	fmt.Println(sa1.last)
	fmt.Println(sa1.person.last)
}
