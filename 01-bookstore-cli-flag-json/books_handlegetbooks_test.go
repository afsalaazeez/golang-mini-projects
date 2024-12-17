package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func TesthandleGetBooks(t *testing.T) {

	testCases := []struct {
		name     string
		all      bool
		id       string
		expected string
	}{
		{
			name:     "Test handleGetBooks with --all flag",
			all:      true,
			id:       "",
			expected: "",
		},
		{
			name:     "Test handleGetBooks with --id flag",
			all:      false,
			id:       "1",
			expected: "",
		},
		{
			name:     "Test handleGetBooks with no flags",
			all:      false,
			id:       "",
			expected: "subcommand --all or --id needed",
		},
		{
			name:     "Test handleGetBooks with --id flag and no matching book",
			all:      false,
			id:       "1000",
			expected: "Book not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			getCmd := flag.NewFlagSet("test", flag.ExitOnError)
			all := getCmd.Bool("all", tc.all, "Get all books")
			id := getCmd.String("id", tc.id, "Get book by id")
			handleGetBooks(getCmd, all, id)

			os.Stdout = oldStdout
			w.Close()

			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			if output != tc.expected {
				t.Errorf("got %v, want %v", output, tc.expected)
			}
		})
	}
}
