package main

import "fmt"

// maps in go are like dictionaries in python or objects in javascript or hashmaps in java
// they are key-value pairs
// the formula for declaring a map is:
// var mapName map[keyType]valueType


func main() {
	// chicken is a map with string keys and int values
	// make in go is used to initialize a map, slice or channel which are dynamic data structures
	// the formula for make is:
	// make(map[keyType]valueType)
	chicken := make(map[string]int)

	chicken["age"] = 2
	chicken["weight"] = 5
	chicken["height"] = 10

	fmt.Println(chicken)

}