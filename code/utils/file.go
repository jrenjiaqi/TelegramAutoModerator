package utils

import (
	"log"
	"os"
)

/*
Append a string and a newline to the end of a file.

parameters:
  - pathname string: the pathname to the file.
  - text string: the text to append (note: a newline char will be added).
*/
func Append_to_file_newline(pathname string, text string) {
	// From: https://pkg.go.dev/os#example_OpenFile_append
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(pathname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("\n")); err != nil { // adds newline after item is appended.
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
