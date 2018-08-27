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
	done := make(chan string)
	go func() {
		done <- sayXConcurrently()
	}()
	fmt.Println(<-done)
}
