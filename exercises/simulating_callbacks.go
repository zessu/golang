package main

import (
	"fmt"
)

func simulate(age int, callback chan bool) {
	if age > 12 {
		callback <- true
	} else {
		callback <- false
	}
}

func main() {
	callback := make(chan bool, 1)
	age := 1122
	simulate(age, callback)
	go func() {
		for {
			select {
			case isGreater := <-callback:
				if isGreater {
					fmt.Println("is greater than 12")
				} else {
					fmt.Println("is less than 12")
				}
			default:
			}
		}
	}()
	fmt.Scanln()
}
