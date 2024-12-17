package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestgetBooks(t *testing.T) {

	testCases := []struct {
		name          string
		mockFileName  string
		mockFileData  string
		expectedBooks []Book
		expectError   bool
	}{
		{
			name:         "Successful Retrieval of Books",
			mockFileName: "books.json",
			mockFileData: `[{"id":"1","title":"Book1","author":"Author1","year":"2001"},{"id":"2","title":"Book2","author":"Author2","year":"2002"}]`,
			expectedBooks: []Book{
				{"1", "Book1", "Author1", "2001"},
				{"2", "Book2", "Author2", "2002"},
			},
			expectError: false,
		},
		{
			name:         "File Not Found Error",
			mockFileName: "non_existent.json",
			expectError:  true,
		},
		{
			name:         "Invalid JSON Data",
			mockFileName: "books.json",
			mockFileData: `invalid`,
			expectError:  true,
		},
		{
			name:          "Empty JSON File",
			mockFileName:  "books.json",
			mockFileData:  ``,
			expectedBooks: []Book{},
			expectError:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.mockFileName != "non_existent.json" {
				err := ioutil.WriteFile(testCase.mockFileName, []byte(testCase.mockFileData), 0644)
				if err != nil {
					t.Fatalf("Failed to create mock file: %v", err)
				}
				defer os.Remove(testCase.mockFileName)
			}

			books := getBooks()

			if !testCase.expectError {
				if len(books) != len(testCase.expectedBooks) {
					t.Errorf("Expected %d books, but got %d", len(testCase.expectedBooks), len(books))
				}
				for i, book := range books {
					if book != testCase.expectedBooks[i] {
						t.Errorf("Expected book %v, but got %v", testCase.expectedBooks[i], book)
					}
				}
			} else {

			}
		})
	}
}
