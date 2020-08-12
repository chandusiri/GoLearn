package main

/*
	// Syntax MAP[key TYPE]Value TYPE{
	//	key: Value,
	//		}
	// Map syntax is similar to Slice []TYPE{}. Go is elegant!

*/
import "fmt"

func main() {

	m := map[string]int{
		"james":           32,
		"Miss MoneyPenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["james"])

	// "comma ok idiom"
	val, ok := m["akbar"]
	fmt.Println(val)
	fmt.Println(ok)

	// Add an element to map
	m["baby"] = 5

	// Common idiomatic chunk of code
	if v, ok := m["baby"]; ok {
		fmt.Println("Oh baby is", v)
	}

	// Loop over a MAP using RANGE just like a Slice
	for k, v := range m {
		fmt.Println(k, "is", v)
	}

	// Delete an entry in map
	delete(m, "Miss MoneyPenny")
	fmt.Println(m)

}
