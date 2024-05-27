package main

import (
	"fmt"
	"os"
)

// A closure is a function that captures the surrounding state of the environment in which it was created.
// In Go, closures are implemented as anonymous functions.
// Closures are useful for capturing variables that are used in a function that is called later.
// In the example below, the sequenceGenerator function returns an anonymous function that captures the variable i.
// The returned function can then be called to generate a sequence of numbers.
func sequenceGenerator() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}

// Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
// defer is often used where e.g. ensure and finally would be used in other languages.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// Defer is commonly used to close files, release resources, and unlock locks.
// In the example below, the readFile function opens a file and defers the file.Close() call to ensure that the file is closed when the function returns.
// The file.Close() call is executed after the readFile function returns.
// This ensures that the file is closed even if an error occurs while reading the file.
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print("Failed to open the file: ", err)
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}

func main() {
	nextNumber := sequenceGenerator()
	readFile("go.mod")

	println(nextNumber())
	println(nextNumber())
	println(nextNumber())
}