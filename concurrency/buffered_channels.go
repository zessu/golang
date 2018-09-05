package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "When is the last time you took a shower"
	wordSlice := strings.Split(phrase, " ")
	channel := make(chan string, len(wordSlice))
	for _, word := range wordSlice {
		channel <- word // this is done async
	}
	close(channel) // doesnt delete the channel, disables sending to channel, but we can still read
	for range wordSlice {
		fmt.Println(<-channel)
	}
}
