package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
Performs HTTP GET on a specified address

parameters:
  - address (string): scheme, host, port (if any), path (e.g. http://example.com:3030/foo/bar or http://example.com/foo/bar)

returns:
  - body_string (string): string representation of response body
*/
func http_get(address string) string {
	resp, err := http.Get(address)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body_bytedata, _ := io.ReadAll(resp.Body)
	body_string := string(body_bytedata)
	return body_string
}

func main() {
	get_response_string := http_get("https://example.com/")
	fmt.Print(get_response_string)
}
