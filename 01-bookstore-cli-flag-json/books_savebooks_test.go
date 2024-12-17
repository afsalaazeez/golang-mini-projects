package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"
)

func TestsaveBooks(t *testing.T) {

	var tests = []struct {
		name    string
		books   []Book
		wantErr error
	}{
		{"Save a valid list of books", []Book{{Id: "1", Title: "Test Book", Author: "Test Author", Price: "10", Imageurl: "http://test.url"}}, nil},
		{"Save an empty list of books", []Book{}, nil},
		{"Save a list of books with invalid data", []Book{{Id: "", Title: "Test Book", Author: "Test Author", Price: "10", Imageurl: "http://test.url"}}, errors.New("Invalid data")},
		{"Save a list of books when the file cannot be written", []Book{{Id: "1", Title: "Test Book", Author: "Test Author", Price: "10", Imageurl: "http://test.url"}}, errors.New("File cannot be written")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := saveBooks(tt.books)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("saveBooks() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				fileBytes, readErr := ioutil.ReadFile("./books.json")
				if readErr != nil {
					t.Errorf("Failed to read the output file: %v", readErr)
				}

				var booksFromFile []Book
				jsonErr := json.Unmarshal(fileBytes, &booksFromFile)
				if jsonErr != nil {
					t.Errorf("Failed to unmarshal the output file content: %v", jsonErr)
				}

				if len(booksFromFile) != len(tt.books) {
					t.Errorf("Mismatch in number of books. got: %v, want: %v", len(booksFromFile), len(tt.books))
				}

				for i, book := range booksFromFile {
					if book != tt.books[i] {
						t.Errorf("Mismatch in book data. got: %v, want: %v", book, tt.books[i])
					}
				}
			}
		})
	}
}
