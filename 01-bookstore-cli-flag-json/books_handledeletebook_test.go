package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Mocking getBooks function
var mockGetBooks = func() []Book {
	return []Book{
		{
			Id:       "1",
			Title:    "Test Book 1",
			Author:   "Test Author 1",
			Price:    "100",
			Imageurl: "http://testurl1.com",
		},
		{
			Id:       "2",
			Title:    "Test Book 2",
			Author:   "Test Author 2",
			Price:    "200",
			Imageurl: "http://testurl2.com",
		},
	}
}

// Mocking saveBooks function
var mockSaveBooks = func(books []Book) error {
	return nil
}

// Mocking checkError function
var mockCheckError = func(err error) {
	if err != nil {
		fmt.Println("Error Happened ", err)
		os.Exit(1)
	}
}

func TesthandleDeleteBook(t *testing.T) {
	// We need a way to inject our mock functions into the original function.
	// This can be done by defining an interface and then implementing it in our test file.
	// This interface will contain all the methods that we want to mock.
	// Then we can pass this interface as a parameter to our function.

	// Here is how you can do it:

	type BookHandler interface {
		GetBooks() []Book
		SaveBooks([]Book) error
		CheckError(error)
	}

	type MockBookHandler struct{}

	func (mbh MockBookHandler) GetBooks() []Book {
		return mockGetBooks()
	}

	func (mbh MockBookHandler) SaveBooks(books []Book) error {
		return mockSaveBooks(books)
	}

	func (mbh MockBookHandler) CheckError(err error) {
		mockCheckError(err)
	}

	// Now modify your handleDeleteBook function to accept this interface as a parameter.
	// Then use this interface to call the methods instead of directly calling the functions.

	var buf bytes.Buffer
	out = &buf

	tests := []struct {
		name       string
		id         string
		wantOutput string
	}{
		{
			name:       "Valid book ID provided for deletion",
			id:         "1",
			wantOutput: "Book deleted successfully",
		},
		{
			name:       "Invalid book ID provided for deletion",
			id:         "3",
			wantOutput: "Book not found",
		},
		{
			name:       "Empty book ID provided for deletion",
			id:         "",
			wantOutput: "Please provide book --id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			deleteCmd := flag.NewFlagSet("test", flag.ContinueOnError)
			id := deleteCmd.String("id", "", "Book ID to delete")
			*id = tt.id
			handleDeleteBook(deleteCmd, id)
			gotOutput := strings.TrimSpace(buf.String())
			if tt.wantOutput != gotOutput {
				t.Errorf("Expected output %q, but got %q", tt.wantOutput, gotOutput)
			}
		})
	}
}
