package main

import (
	"fmt"
)

func main() {
	x := [5]int{}
	fmt.Println(x)
	x[0] = 7
	fmt.Println(x)
	// x[100] = 12 arrays have a fixed size
	// append(x, 3) we also cannot do this
}
