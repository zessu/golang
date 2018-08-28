package main

import (
	"fmt"
)

func sayXConcurrently() string {
	return "This string was concurrently returned"
}

func sayXBlock() string {
	return "Say X normal"
}

func main() {
	fmt.Println(sayXBlock())
	done := make(chan string) // second param for number of items to fit in channel
	go func() {
		done <- sayXConcurrently()
	}()
	fmt.Println(<-done)
}
