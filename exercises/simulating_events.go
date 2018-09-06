package main

import (
	"fmt"
)

// Button represents a HTML button
type Button struct {
	eventListeners map[string][]chan string
}

// MakeButton returns a button
func MakeButton() *Button {
	button := &Button{}
	button.eventListeners = make(map[string][]chan string)
	return button
}

// AddEventListener adds an events listener to the button
func (button *Button) AddEventListener(event string, channel chan string) {
	if _, present := button.eventListeners[event]; present {
		// add communication channel
		button.eventListeners[event] = append(button.eventListeners[event], channel)
	}
	// create
	button.eventListeners[event] = []chan string{channel}
}

// RemoveEventListener removes an events listener from the button
func (button *Button) RemoveEventListener(event string) {
	delete(button.eventListeners, event)
}

// RemoveSpecificChannelFromMap removes a specific channel event
func (button *Button) RemoveSpecificChannelFromMap(event string, channelToRemove chan string) {
	if channels, ok := button.eventListeners[event]; ok {
		for index, value := range channels {
			if value == channelToRemove {
				button.eventListeners[event] = append(channels[:index], channels[index+1:]...)
				break
			}
		}
	}
}

// TriggerEvent siluates a trigger event
func (button *Button) TriggerEvent(event string, message string) {
	if _, present := button.eventListeners[event]; present {
		// since this is present, lets start it off on
		for _, handler := range button.eventListeners[event] {
			go func(handler chan string) {
				handler <- message
			}(handler)
		}
	}
}

func main() {
	button := MakeButton()
	clickChannel := make(chan string, 1)
	hoverChannel := make(chan string, 1)
	// set up event handlers
	go func() {
		for {
			select {
			case click := <-clickChannel:
				fmt.Println(click)
			case hover := <-hoverChannel:
				fmt.Println(hover)
			default:
				// wait
			}
		}
	}()
	button.AddEventListener("click", clickChannel)
	button.AddEventListener("hoverin", hoverChannel)
	button.AddEventListener("hoverout", hoverChannel)
	fmt.Println(button.eventListeners)
	button.TriggerEvent("click", "user clicked")
	button.TriggerEvent("hoverin", "user hovered in")
	button.TriggerEvent("hoverout", "user hovered out")
	fmt.Scanln()
}
