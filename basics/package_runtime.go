package main

import (
	"fmt"
	"runtime"
)

// runtime package can control GOs runtime system. It contains some operations to do that.
// It also has some constants. Lets print few of them.
// go to godoc.org/runtime to see more details and its index

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
