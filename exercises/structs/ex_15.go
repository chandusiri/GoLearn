package main

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/
import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	flavors   []string
}

func main() {
	p1 := person{
		firstName: "Rk",
		lastName:  "B",
		flavors: []string{
			"Mango",
			"Strawberry",
		},
	}

	p2 := person{
		firstName: "SC",
		lastName:  "A",
		flavors: []string{
			"Sitaphal",
			"Blackberry",
		},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	for i, v := range p1.flavors {
		fmt.Println(i, v)
	}

	for _, v := range p2.flavors {
		fmt.Println(v)
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(m)

	// Creating own type (here x) for builtin types is not recommended.
	type x []string
	var y x
	y = []string{"a", "s", "d", "f"}
	fmt.Println(y)
	fmt.Printf("%T\n", []string(y))

}
