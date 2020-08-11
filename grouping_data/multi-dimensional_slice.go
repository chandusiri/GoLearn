package main

import "fmt"

func main() {
	emp1 := []string{"Chandans", "HPE", "Go"}
	emp2 := []string{"Gnan", "HPE", "Infrastructure"}
	employees := [][]string{emp1, emp2}
	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(employees)
	fmt.Println(len(emp1))
	fmt.Println(len(employees)) //2 dimensional

	// APPEND SLICE to a multi dimensional SLICE
	digits := [][]int{{0, 2, 4, 6, 8}, {1, 3, 5}}
	digits = append(digits, []int{10, 20})
	digits[2] = append(digits[2], 30, 40, 50)
	fmt.Println(digits)
}
