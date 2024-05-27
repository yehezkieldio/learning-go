package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

var bookDetailsPattern = regexp.MustCompile(`Title: (.+), Author: (.+), Pages: (\d+)`)

func ParseBooksFromFile(filename string) ([]Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var books []Book
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matchers := bookDetailsPattern.FindStringSubmatch(scanner.Text())
		if len(matchers) != 4 {
			continue
		}

		title := matchers[1]
		author := matchers[2]
		// convert the string to an integer
		pages, err := strconv.Atoi(matchers[3])
		if err != nil {
			fmt.Println("Failed to convert the string to an integer")
			continue
		}

		books = append(books, Book{title, author, pages})
	}

	return books, scanner.Err()
}

func main() {
	filename := "regex-data-parse/books.txt"
	books, err := ParseBooksFromFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Pages: %d\n", book.Title, book.Author, book.Pages)
	}
}