package main

import (
	"fmt"
)

type changeable interface {
	Change(nameName string)
}

type User struct {
	fname string
	lname string
}

// types implement interfaces
/* in go, any type that has a method accepting of the type and with a name of the function implementing the interface
implements the interface */

// Change modifies a users fname
func (user *User) Change(name string) {
	user.fname = name
}

// item passed must implement a method Change
func changeName(item changeable) {
	item.Change("frog")
}

func main() {
	user := User{fname: "john", lname: "dhoe"}
	changeName(&user)
	fmt.Println(user)
}
