package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCheckError(t *testing.T) {

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var testCases = []struct {
		name        string
		err         error
		expectPanic bool
		expectedMsg string
	}{
		{
			name:        "Error is nil",
			err:         nil,
			expectPanic: false,
			expectedMsg: "",
		},
		{
			name:        "Error is not nil",
			err:         errors.New("test error"),
			expectPanic: true,
			expectedMsg: "Error Happened test error\n",
		},
		{
			name:        "Error with specific message",
			err:         errors.New("specific test error"),
			expectPanic: true,
			expectedMsg: "Error Happened specific test error\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if r == nil && tc.expectPanic {
					t.Errorf("%s: expected panic but did not panic", tc.name)
				} else if r != nil && !tc.expectPanic {
					t.Errorf("%s: did not expect panic but panicked", tc.name)
				}
			}()

			checkError(tc.err)

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = oldStdout

			if string(out) != tc.expectedMsg {
				t.Errorf("%s: expected %s but got %s", tc.name, tc.expectedMsg, out)
			}
		})
	}
}

