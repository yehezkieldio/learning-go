package main

import "fmt"

// Generics allow you to write functions, data structures, and algorithms that can work with any type.

// The Filter function takes a slice of any type and a criteria function that returns a boolean.
// The slice is a collection of elements of the same type.
// The criteria function is used to filter the elements of the slice.
func Filter[T any](slice []T, criteria func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if criteria(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	ints := []int{1, 2, 3, 4, 5}

	// Filter the slice of integers to get only the even numbers.
	// The criteria function checks if the number is even by using the modulo operator.
	// We get even numbers by modding the number by 2 and checking if the result is 0.
	// Because modding by 2 gives us the remainder when dividing by 2, if the remainder is 0, the number is even.
	even := Filter(ints, func(n int) bool {
		return n % 2 == 0
	})
	fmt.Println(even)

	// Fiter the slice of strings to get only the strings that start with the letter 'a'.
	// The criteria function checks if the first character of the string is 'a'.
	// We get the first character of the string by indexing the string with [0].
	// If the first character is 'a', we return true; otherwise, we return false.
	strings := []string{"apple", "banana", "cherry", "date", "elderberry", "apricot"}
	withA := Filter(strings, func(s string) bool {
		return s[0] == 'a'
	})
	fmt.Println(withA)
}