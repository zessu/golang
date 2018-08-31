package main

import "fmt"

// Snake represents a cow
type Snake struct {
}

// SaySomething prints whetever we pass to it
func (s *Snake) SaySomething(message string) {
	fmt.Println(message)
}

// CreateSnake creates snake
func CreateSnake() *Snake {
	snake := &Snake{}
	return snake
}

// CreateChannel asdasd
func CreateChannel(channel chan string, done chan bool, obj *Snake) {
	for message := range channel { // waiting for line 34, close for loop when like 34 is run
		obj.SaySomething(message)
	}
	done <- true
}

func main() { //nthread
	channel := make(chan string)
	done := make(chan bool)
	snake := CreateSnake()
	go CreateChannel(channel, done, snake)
	channel <- "saysomethingelse"
	close(channel)
	<-done // waiting for line 25
}
