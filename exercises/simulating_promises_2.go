package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Promise represents a Promise from the promise concurrency model
type Promise struct {
	successChannel chan interface{}
	errorChannel   chan error
}

// newPromise returns a pointer to a promise
func newPromise() *Promise {
	p := &Promise{}
	p.successChannel = make(chan interface{}, 1)
	p.errorChannel = make(chan error, 1)
	return p
}

// Then takes a promise and runs it on different paths depending on how it executes
func (chained *Promise) Then(success func(message interface{}) error, handleError func(err error), done chan interface{}) *Promise {
	p := newPromise()
	// this will listen to the promise it was passed, successChannel and errorChannel and pass it to the relevant method
	go func() {
	prom:
		for {
			select {
			case message := <-chained.successChannel:
				err := success(message)
				if err != nil {
					p.errorChannel <- err // pass error to next promise
					break prom
				} else {
					p.successChannel <- message // chain the data onto the next channel where we want to operate it
					break prom
				}
			case err := <-chained.errorChannel:
				handleError(err)
				break prom
			default:
			}
		}
		close(done)
	}()
	return p
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

// AsyncActivity
func fetchAPIData() *Promise {
	p := newPromise()
	url := "https://www.google.com"
	go MakeRequest(url, p.successChannel, p.errorChannel)
	return p
}

func successHandler(message interface{}) error {
	fmt.Println(message)
	return nil
}

func errorHandler(err error) {
	fmt.Println(err)
}

func main() {
	CloseProg := make(chan interface{})
	fetchAPIData().Then(successHandler, errorHandler, CloseProg)
	<-CloseProg
}
