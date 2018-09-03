package main

import "fmt"

// Reader represents a source we can grab data from
type Reader interface {
	Read() bool // encapsulation, anything of this type will have Read implemented differently though
}

// File represents a File reader type
type File struct {
}

// TCP represents a TCP reader type
type TCP struct {
}

// WebSocket represents a WebSocket reader type
type WebSocket struct {
}

func (f *File) Read() bool {
	fmt.Println("Reading from a file")
	return true
}

func (tcpconn *TCP) Read() bool {
	fmt.Println("Reading from a TCP connection")
	return true
}

func (websock *WebSocket) Read() bool {
	fmt.Println("Reading from a WebSocket")
	return true
}

// ReadFromSource given a source we can read from, reads from the source
// Take different types of sources, is polymorphic
func ReadFromSource(source Reader) {
	source.Read()
}

func main() {
	file := &File{}
	TCPConnection := &TCP{}
	WebSocketConnection := &WebSocket{}
	ReadFromSource(file) // we achieved encapsulation here, message passing as well
	ReadFromSource(TCPConnection)
	ReadFromSource(WebSocketConnection)
}
