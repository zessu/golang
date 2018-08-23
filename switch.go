package main

import "fmt"

type user struct {
	fname string
	lname string
}

func getMessage(name string) (prefix string) {
	switch name {
	case "andrew":
		prefix = "sir"
	case "wambua":
		prefix = "madam"
		// we can create a list here
	case "john", "random":
		prefix = "bulldozer"
		// fallthrough -> we can add this to 'fall through' and set the value to mamamama
	default:
		prefix = "mamamama"
	}
	return
}

func weirdSwitch(name string) (prefix string) {
	// we can leave out declaration of var to switch and include it inside the tests statement
	switch {
	case name == "andrew":
		prefix = "sir"
	case name == "wambua":
		prefix = "madam"
		// we can create a list here
	case name == "john" || name == "random":
		prefix = "bulldozer"
		// fallthrough -> we can add this to 'fall through' and set the value to mamamama
	default:
		prefix = "mamamama"
	}
	return
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("an integer")
	case string:
		fmt.Println("a string")
	case user:
		fmt.Println("a user object")
	default:
		fmt.Println("unsupported object type")
	}
}

func main() {
	x := user{fname: "random", lname: "dhoe"}
	fmt.Println(getMessage(x.fname))
	fmt.Println(weirdSwitch(x.fname))
	switchOnType(1)
	switchOnType("a string")
	switchOnType(user{fname: "x", lname: "y"})
	switchOnType(1.222)
}
