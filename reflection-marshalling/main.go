package main

import (
	"encoding/json"
	"fmt"
	"log"
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

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	myVar := 42
	inspectVariable(myVar)

	p := Person{Name: "Eliziel", Age: 18}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("error marshalling to JSON: %s", err)
	}

	fmt.Println(string(jsonData))

	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		log.Fatalf("error unmarshalling from JSON: %s", err)
	}

	fmt.Println(p2)
}