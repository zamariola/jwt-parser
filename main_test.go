package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// A sample JWT token for testing
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	// Keep backup of the real stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Create a new pipe to simulate stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write the token to the write-end of the pipe
	w.Write([]byte(token))
	w.Close()

	// Keep backup of the real stdout
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	// Create a new pipe to capture stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Run the main function
	main()

	// Close the write-end of the stdout pipe
	wOut.Close()

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, rOut)

	// Restore stdout
	os.Stdout = oldStdout

	output := buf.String()

	// Check for header
	if !strings.Contains(output, `"alg": "HS256"`) {
		t.Error("Expected to find algorithm in header")
	}

	// Check for payload
	if !strings.Contains(output, `"name": "John Doe"`) {
		t.Error("Expected to find name in payload")
	}
}
