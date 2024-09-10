package repo

import (
	"log"
	"os"
	"strings"
)

func Get_string_slice_from_file(filepath string, delimiter string) []string {
	file_bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Panic(err)
	}
	file_string := string(file_bytes)
	return strings.Split(file_string, delimiter)
}
