package main

import "fmt"

func main() {
	// for initialization; condition; post{}
	for i := 0; i < 4; i++ {
		fmt.Println("The outer loop has", i)
		for j := 36; j > 33; j-- {
			fmt.Printf("\t\t\tThe inner loop has both decimal %v  and ascii %#U\n", j*i, j*i)
		}
	}
}
