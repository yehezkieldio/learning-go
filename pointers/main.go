package main

import "fmt"

func main() {
	var a int = 58
	var p *int = &a

	fmt.Println("Address of a: ", p)
	fmt.Println("Value of a: ", *p)
}