package main

import "fmt"

// User defines a custom user object
// capital U makes it global, small u private to this file
type User struct {
	fname string
	lname string
}

func main() {
	var message = User{fname: "john", lname: "dhoe"}
	fmt.Println(message)
}
