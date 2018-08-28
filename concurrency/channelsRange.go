package main

import "fmt"

type User struct {
	fname string
	lname string
}

// SendUsersToChannel used a test function for using channels with range
func SendUsersToChannel(users []User, channel chan User) {
	for _, user := range users {
		channel <- user
	}
	close(channel)
}

func main() {
	users := make([]User, 2)
	users[0] = User{fname: "greedy", lname: "joe"}
	users[1] = User{fname: "mark", lname: "nganag"}
	channel := make(chan User)
	go SendUsersToChannel(users, channel)
	for s := range channel {
		fmt.Println(s)
	}
}
