package main

import (
	"fmt"
)

func main() {
	// make(chan string); if our channel is not buffered, this will block till something can receive it, causing a deadlock
	ch := make(chan string, 1)
	ch <- "hello world" // if our channel wasn't buffered, we would block here
	fmt.Println(<-ch)
}
