/*
Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
print out the TYPE of the array

*/
package main

import "fmt"

func main() {
	arr := [5]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i, v := range arr {
		fmt.Println(i, v)

	}
	fmt.Printf("%T", arr)
}
