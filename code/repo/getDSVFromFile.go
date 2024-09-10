package repo

import (
	"log"
	"os"
	"strings"
)

/*
Returns a slice of string from the contents of a DSV file.

parameters:
  - filepath string: file path to the DSV file.
  - delimiter string: the delimiter used in this DSV file.

returns:
  - []string: slice of strings split from the DSV file.
*/
func Get_string_slice_from_file(filepath string, delimiter string) []string {
	file_bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Panic(err)
	}
	file_string := string(file_bytes)
	return strings.Split(file_string, delimiter)
}
