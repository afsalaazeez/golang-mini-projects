package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
	"io"
)

func TesthandleAddBook(t *testing.T) {

	testCases := []struct {
		name            string
		id              string
		title           string
		author          string
		price           string
		image_url       string
		addNewBook      bool
		expectedOutput  string
		expectedSuccess bool
	}{
		{
			name:            "Test Add Book with all valid parameters",
			id:              "1",
			title:           "Test Book",
			author:          "Test Author",
			price:           "10",
			image_url:       "http://test.com",
			addNewBook:      true,
			expectedOutput:  "Book added successfully\n",
			expectedSuccess: true,
		},
		{
			name:            "Test Add Book with missing parameters",
			id:              "",
			title:           "Test Book",
			author:          "",
			price:           "10",
			image_url:       "http://test.com",
			addNewBook:      true,
			expectedOutput:  "Please provide book id, title, author,price\n",
			expectedSuccess: false,
		},
		{
			name:            "Test Update existing book",
			id:              "1",
			title:           "Updated Test Book",
			author:          "Updated Test Author",
			price:           "20",
			image_url:       "http://testupdated.com",
			addNewBook:      false,
			expectedOutput:  "Book added successfully\n",
			expectedSuccess: true,
		},
		{
			name:            "Test Update non-existing book",
			id:              "100",
			title:           "Updated Test Book",
			author:          "Updated Test Author",
			price:           "20",
			image_url:       "http://testupdated.com",
			addNewBook:      false,
			expectedOutput:  "Book not found\n",
			expectedSuccess: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			os.Args = []string{"cmd", "add", "-id", tc.id, "-title", tc.title, "-author", tc.author, "-price", tc.price, "-image_url", tc.image_url}
			addCmd := flag.NewFlagSet("add", flag.ExitOnError)
			id := addCmd.String("id", "", "Book id")
			title := addCmd.String("title", "", "Book title")
			author := addCmd.String("author", "", "Book author")
			price := addCmd.String("price", "", "Book price")
			image_url := addCmd.String("image_url", "", "Book image url")

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			handleAddBook(addCmd, id, title, author, price, image_url, tc.addNewBook)

			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				outC <- buf.String()
			}()
			w.Close()
			os.Stdout = old
			out := <-outC

			if out != tc.expectedOutput {
				t.Errorf("Expected output to be %v but got %v", tc.expectedOutput, out)
			}

			if tc.expectedSuccess {
				books := getBooks()
				for _, book := range books {
					if book.Id == tc.id && book.Title == tc.title && book.Author == tc.author && book.Price == tc.price && book.Imageurl == tc.image_url {
						return
					}
				}
				t.Errorf("Expected book to be added/updated but it was not")
			}
		})
	}
}
