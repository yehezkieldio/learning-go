package main

import "fmt"

// In Go, methods that need to modify the state of a struct must receive a pointer to the struct.
// This is because Go is a pass-by-value language, meaning that when you pass a variable to a function, Go creates a copy of that variable within the function.
// If you want to modify the original variable, you need to pass a pointer to it.
// Any changes made to the variable inside the function do not affect the original variable.

type Person struct {
	Name string
	Age int
}

func birthday(p *Person) {
	p.Age += 1
}

func main() {
	person := Person{ Name: "Eliziel", Age: 18 }
	fmt.Println(person.Name)

	birthday(&person)
	fmt.Println(person.Age)
}

