package main

import (
	"fmt"
)

type overarching struct {
	name string
	age  string
}

func (o *overarching) saySomething() {
	fmt.Println("we got here by composition")
}

func (o *overarching) rename(name string) {
	o.name = name
}

type human struct {
	overarching
}

func main() {
	h := &human{}
	h.rename("john dhoe")
	fmt.Println(h.name)
	h.saySomething()
}
