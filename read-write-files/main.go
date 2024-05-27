package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func SaveBooks(filename string, books []Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, book := range books {
		jsonData, err := json.Marshal(book)
		if err != nil {
			return err
		}

		_, err = writer.WriteString(string(jsonData) + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	books := []Book{
		{"The Art of Computer Programming", "Donald Knuth", 3168},
		{"The Go Programming Language", "Alan A. A. Donovan & Brian W. Kernighan", 380},
		{"Design Patterns: Elements of Reusable Object-Oriented Software", "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides", 395},
	}
	filename := "read-write-files/books.json"

	if err := SaveBooks(filename, books); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
