package main

import (
	"fmt"
)

func printWithLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func whileLoop() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func main() {
	printWithLoop()
	whileLoop()
}
