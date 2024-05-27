package main

import (
	"fmt"
	"reflect"
)

// Data marshalling is the process of converting data from one format to another.
// In Go, the encoding/json package provides support for marshalling and unmarshalling data to and from JSON.

// The reflect package and it's support for data marshalling come into play for addressing dynamic data types and structures.
// Reflection provides the ability to inspect and manipulate the structure of data at runtime.

func inspectVariable(variable interface{}) {
	t := reflect.TypeOf(variable)
	v := reflect.ValueOf(variable)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
}

func main() {
	myVar := 42
	inspectVariable(myVar)
}