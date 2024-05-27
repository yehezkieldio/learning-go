package main

import "fmt"

// Set is a map with string keys and struct{} values
// to represent a set
type Set map[string]struct{}

func (s Set) Add(element string) {
	s[element] = struct{}{}
}

func (s Set) Remove(element string) {
	// delete is a built-in function to remove an element from a map
	delete(s, element)
}

func (s Set) Contains(element string) bool {
	_, found := s[element]
	return found
}

func main() {
	mySet := make(Set)

	mySet.Add("foo")
	mySet.Add("bar")
	mySet.Add("baz")

	// Contains returns true
	// because "foo" is in the set
	if mySet.Contains("foo") {
		println("foo is in the set")
	}

	mySet.Remove("foo")

	fmt.Println(mySet)
}