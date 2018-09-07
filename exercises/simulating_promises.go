package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Promise represents a promise, from the promises concurrency model
type Promise struct {
	successChannel chan interface{}
	errorMessage   chan error
}

// SimulateAsyncActivity simulate an async activity, returns a promise
func SimulateAsyncActivity(p *Promise) {
	// do some async stuff here and call relevant functions
	url := "https://www.google.com"
	go MakeRequest(url, p.successChannel, p.errorMessage)
}

// MakeRequest makes a http request
func MakeRequest(url string, successChannel chan interface{}, errorChannel chan error) {
	resp, error := http.Get(url)
	if error != nil {
		errorChannel <- error
		close(errorChannel)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		successChannel <- string(body)
		close(successChannel)
	}
}

// CreatePromise create a promise object
func createPromise(done chan interface{}) *Promise {
	p := &Promise{}
	p.successChannel = make(chan interface{}, 1)
	p.errorMessage = make(chan error, 1)

	go func() {
	promise:
		for {
			select {
			case message, valid := <-p.successChannel:
				if !valid {
					break promise
				}
				fmt.Println(message)
			case err, valid := <-p.errorMessage:
				if !valid {
					break promise
				}
				fmt.Println(err)
			default:
			}
		}
		close(done)
	}()
	return p
}

func main() {
	reqCompletedChan := make(chan interface{})
	promise := createPromise(reqCompletedChan)
	SimulateAsyncActivity(promise)
	<-reqCompletedChan
}
