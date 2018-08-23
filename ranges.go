package main

import (
	"fmt"
)

type User struct {
	fname string
	lname string
}

func printAllWithRange(userList []User) {
	for _, user := range userList {
		fmt.Println(user)
	}
}

func main() {
	users := []User{
		{fname: "wambua", lname: "makenzi"},
		{fname: "samantha", lname: "andea"},
		{fname: "grace", lname: "havemercy"},
	}
	printAllWithRange(users)
}
