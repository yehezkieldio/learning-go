package main

import (
	"fmt"
	"multiple-files/util"
)

func main() {
	// Add is not accessible here
	// fmt.Println(Add(1, 2))

	// Instead, we need to import the package
	// and use the package name as a prefix
	fmt.Println(util.Add(1, 2))
}