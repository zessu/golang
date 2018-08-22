package main

import "fmt"

// User defines our user
type User struct {
	fname string
	lname string
}

// Greet greets a user
func greet(user User) {
	fmt.Println("hi", user.fname, user.lname, " we are glad to have you")
}

// functionReturns returns something
func functionReturns() string {
	message := "this one returns"
	return message
}

// returnsTwoTypes returns more that one value, allowing us to check for errors in a new way
// compared to older languages where we had to check them in-band
func returnsTwoTypes() (name string, errored bool) {
	errored = true
	name = "johndhoe"
	return
}

// variadicFunc takes a variable number of args
func variadicFunc(name string, varArgs ...string) {
	fmt.Println(name, varArgs[0], varArgs[1], varArgs[2])
}

// takesFunction demostrates a function that accepts another function as an arg
func takesFunction(name string, do func(string)) {
	do(name)
}

func main() {
	var user = User{fname: "john", lname: "dhoe"}
	greet(user)
	fmt.Println(functionReturns())
	// ignore the first value returned, if we don't need to use it, otherwise we'd get errors
	message, error := returnsTwoTypes()
	_, error2 := returnsTwoTypes() // ignore message2, do not need to declare, will otherwise get an error
	if error {
		fmt.Println("There was an error")
	}
	if error2 {
		fmt.Println("There was an error")
	}
	fmt.Println(message)
	variadicFunc("john", "dhoe", "emmanuela", "somebodyelse")
	takesFunction("welovetoparty", func(arg1 string) { fmt.Println(arg1) })
}
