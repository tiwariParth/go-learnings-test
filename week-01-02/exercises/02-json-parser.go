package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct to represent a Book
// TODO: Create a struct with appropriate fields

// Define a struct to represent a Library
// TODO: Create a struct with appropriate fields

// parseBooksJSON parses JSON data into a slice of Book structs
func parseBooksJSON(jsonData []byte) ([]Book, error) {
	// TODO: Implement this function to parse the JSON into a slice of Book structs
	return nil, nil
}

// printBooksByGenre prints all books of a specific genre
func printBooksByGenre(books []Book, genre string) {
	// TODO: Implement this function to print only books of the specified genre
}

// addBookToLibrary adds a book to the library's collection
func addBookToLibrary(library *Library, book Book) {
	// TODO: Implement this function to add a book to the library
}

func main() {
	// Sample JSON data representing books
	jsonData := `[
		{
			"title": "The Go Programming Language",
			"author": "Alan A. A. Donovan & Brian W. Kernighan",
			"year": 2015,
			"genre": "Programming",
			"isbn": "978-0134190440"
		},
		{
			"title": "Moby Dick",
			"author": "Herman Melville",
			"year": 1851,
			"genre": "Fiction",
			"isbn": "978-1503280786"
		},
		{
			"title": "Go in Action",
			"author": "William Kennedy",
			"year": 2015,
			"genre": "Programming",
			"isbn": "978-1617291784"
		}
	]`

	// TODO: Parse the JSON data into a slice of Book structs
	
	// TODO: Create a library and add the books to it
	
	// TODO: Print all programming books
}
