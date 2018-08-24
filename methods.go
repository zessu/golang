package main

import (
	"fmt"
)

type User struct {
	fname string
	lname string
}

// value receiver, associate by passing struct after func
func (user User) renameWithoutPointer(newName string) {
	// we passed a copy of the object not the original
	user.fname = newName
}

// pointer receiver, associate by passing struct after func
func (user *User) renameWithPointer(newName string) {
	// we are passing a pointer here, so we can modify our code
	user.fname = newName
}

func main() {
	user1 := User{fname: "john", lname: "dhoe"}
	user1.renameWithoutPointer("wambua")
	fmt.Println(user1)
	user1.renameWithPointer("wambua")
	fmt.Println(user1)
}
