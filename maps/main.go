package main

import "fmt"

// maps in go are like dictionaries in python or objects in javascript or hashmaps in java
// they are key-value pairs
// the formula for declaring a map is:
// var mapName map[keyType]valueType


func main() {
	var chicken map[string]int
	// chicken is a map with string keys and int values
	chicken = map[string]int{}

	chicken["age"] = 2
	chicken["weight"] = 5
	chicken["height"] = 10

	fmt.Println(chicken)

}