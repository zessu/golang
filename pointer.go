package main

import (
	"fmt"
)

func main() {
	name := "wambua makenzi"
	name1 := &name
	fmt.Println(name, *name1)
}
