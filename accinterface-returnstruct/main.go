package main

import "bytes"

// Accept Interface Return Struct
// Accepting interfaces champions flexibility, ensuring functions aren’t tethered to concrete types. This strategy not only promotes adaptability but also clearly communicates expected behaviors.

// define a Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// a function that accepts a Reader interface and returns a struct
func NewContent(reader Reader) (*ContentResult, error) {
	// create a buffer to store the data
    buf := new(bytes.Buffer)
	// read from the reader and store the data in the buffer
    _, err := buf.ReadFrom(reader)
	// return the ContentResult struct and any error
    return &ContentResult{data: buf.String()}, err
}

type ContentResult struct {
    data string
}

func (cr *ContentResult) String() string {
    return cr.data
}

