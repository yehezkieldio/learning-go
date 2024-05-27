package main

import "fmt"

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	// Sprint formats using the default formats for its operands and returns the resulting string.
	return fmt.Sprintf("Error code %d: %s", e.Code, e.Message)
}

func myFunction() error {
	return &MyError{Message: "Something went wrong", Code: 500}
}

func main() {
	err := myFunction()
	if err != nil {
		switch e := err.(type) {
		case *MyError:
			fmt.Println("Custom error:\n", e)
		default:
			fmt.Println("Default error:\n", e)
		}
	}
}
