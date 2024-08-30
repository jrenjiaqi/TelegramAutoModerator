package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

/*
Helper function: creates a HTTP client with a particular timeout.

parameters:
  - timeout_seconds int: timeout for client's actions.

returns:
  - http.Client: http client with specified timeout.
*/
func get_http_client_timeout(timeout_seconds int) http.Client {
	// Golang Time: stackoverflow.com/a/17573390/24829786
	return http.Client{
		Timeout: time.Duration(timeout_seconds) * time.Second,
	}
}

/*
Uses HTTP GET to a particular URI to obtain a JSON response.

parameters:
  - uri string: URI to send the GET request.
  - timeout_seconds int: how many seconds until client should timeout.
  - target interface{}: pointer to a struct to fit decoded JSON into.

mutates:
  - target interface().

returns:
  - error: any error that occurs.
*/
func Http_GET_JSON(uri string, timeout_seconds int, target interface{}) error {
	/* A. JSON Decoding: stackoverflow.com/a/31129967/24829786
	** B. What is interface{}: https://go.dev/tour/methods/14 */
	client := get_http_client_timeout(5)
	resp, err := client.Get(uri)
	if err != nil {
		return err
	}
	// "docs: It is the caller's responsibility to close Body."
	defer resp.Body.Close()
	/* "A: "The ideal way is not to use ioutil.ReadAll, but rather use a decoder on the reader directly.
	   Here's a nice function that gets a url and decodes its response onto a target structure." */
	return json.NewDecoder(resp.Body).Decode(target)
}
