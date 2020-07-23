package main

import "fmt"

// see below kb -> mb -> gb before understanding this const section
const (
	_ = iota             // 0 is ignored and not used anywhere
	k = 1 << (10 * iota) // shift digit 1, (ten times iota) 10 * 1 to get value 1024 in decimal
	m = 1 << (10 * iota) // shift digit 1, 10 * 2 times
	g = 1 << (10 * iota) // shift digit 1, 10 * 3 times
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b", x, x)
	fmt.Println()
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	// kb -> mb -> gb
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t\t%b\n", kb, kb) // 1 followed by 10 zeros (1 * 10 zeros) in binary format leads to decimal 1024
	fmt.Printf("%d\t\t\t%b\n", mb, mb) // 1 followed by 20 zeros (2 * 10 zeros)
	fmt.Printf("%d\t\t%b\n", gb, gb)   // 1 followed by 30 zeros (3 * 10 zeros)

	// bit shifting using iota
	fmt.Println(k)
	fmt.Printf("%d\t\t\t%b\n", k, k)
	fmt.Printf("%d\t\t\t%b\n", m, m)
	fmt.Printf("%d\t\t%b\n", g, g)
}
