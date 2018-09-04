package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4) // allow go to use 4 processors if available

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("what is this")
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("what is that")
		}
	}()

	dur, _ := time.ParseDuration("1ms")
	time.Sleep(dur)
}
