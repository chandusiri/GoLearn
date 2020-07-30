package main

import "fmt"

func main() {
	// Instead of multiple if else statements
	foo := "bar"
	switch foo {
	case "bar", "baar", "baaaaaaaar":
		fmt.Println("Go get Red Wine!")
	// case "bar", "box" -> duplicate case "bar" in switch
	case "box":
		fmt.Println("Go get case")
	case "beer":
		fmt.Println("not available")
	default:
		fmt.Println("stock empty")
	}
}
