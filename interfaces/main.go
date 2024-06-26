package main

import "fmt"

// Interfaces express a contract of behavior without specifying how the behavior is implemented.
// The implementation of shared methods allows a single function to communicate with different types, achieving polymorphism.

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Robot struct {
	Model string
}

// The Speak method is implemented for the Robot type.
func (r Robot) Speak() string {
	return "Beep Boop!"
}

// This function works with any type that implements the Speaker interface.
func introduceSpeaker(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{ Name: "Buddy" }
	robot := Robot{ Model: "T-800" }

	introduceSpeaker(dog)
	introduceSpeaker(robot)
}

