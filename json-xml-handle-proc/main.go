package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
	Pages  int    `json:"pages" xml:"pages"`
}

type Library struct {
	Books []Book `json:"books" xml:"book"`
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

func ExportBooksToXML(books []Book) (string, error) {
	library := Library{Books: books}
	xmlData, err := xml.MarshalIndent(library, "", " ")
	if err != nil {
		return "", err
	}

	return xml.Header + string(xmlData), nil
}

func WriteBooksToXMLFile(filename string, books []Book) error {
	library := Library{Books: books}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", " ")
	return encoder.Encode(library)
}

func ImportBooksFromXML(xmlData string) ([]Book, error) {
	var library Library
	err := xml.Unmarshal([]byte(xmlData), &library)
	if err != nil {
		return nil, err
	}

	return library.Books, nil
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

	xmlOutput, err := ExportBooksToXML(books)
	if err != nil {
		fmt.Println("error exporting books", err)
		return
	}
	fmt.Println("XML Output", xmlOutput)

	filename = "read-write-files/books.xml"
	if err := WriteBooksToXMLFile(filename, books); err != nil {
		fmt.Println("Error writing XML file:", err)
		return
	}

	importedBooks, err := ImportBooksFromXML(xmlOutput)
	if err != nil {
		fmt.Println("Error importing books from XML:", err)
		return
	}
	fmt.Println("Imported books:", importedBooks)
}
