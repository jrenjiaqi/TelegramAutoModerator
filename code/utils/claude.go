package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ClaudePayload struct {
	Model     string          `json:"model"`
	MaxTokens int             `json:"max_tokens"`
	Messages  []ClaudeMessage `json:"messages"`
	System    string          `json:"system"`
}

type ClaudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

/*
Send to Claude one message, and receive a response via mutation of parameter.

parameters:
  - response_struct_ptr interface{}: struct to be mutated into Claude's response.
  - claude_api_uri string: Claude API endpoint.
  - claude_api_key string: Claude API key.
  - system_prompt string: system prompt for Claude to generate formatted response.
  - text_to_send string: text to send to Claude.
  - response_timeout_seconds int: number of seconds until the request timeouts.
  - debug_mode bool: true for debug information and false otherwise.

mutates:
  - response_struct_ptr interface{}.
*/
func Send_to_claude_one_message(
	response_struct_ptr interface{},
	claude_api_uri string,
	claude_api_key string,
	system_prompt string,
	text_to_send string,
	response_timeout_seconds int,
	debug_mode bool,
) {
	var err error
	claude_message := ClaudeMessage{Role: "user", Content: text_to_send}
	claude_payload := ClaudePayload{
		Model:     "claude-3-5-sonnet-20240620",
		MaxTokens: 1024,
		Messages:  []ClaudeMessage{claude_message},
		System:    system_prompt,
	}
	if debug_mode {
		log.Printf("CLAUDE PAYLOAD AS STRUCT: %+v\n", claude_payload)
	}
	// reference from: https://stackoverflow.com/a/24455606
	json_byte_buffer := new(bytes.Buffer)
	err = json.NewEncoder(json_byte_buffer).Encode(claude_payload)
	if err != nil { // something went wrong with encoding the payload.
		log.Panic(err)
	}
	if debug_mode {
		log.Printf("CLAUDE PAYLOAD AS ENCODED JSON: %s\n", json_byte_buffer)
	}
	req, _ := http.NewRequest("POST", claude_api_uri, json_byte_buffer)
	req.Header.Set("x-api-key", claude_api_key)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Duration(response_timeout_seconds) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response_struct_ptr)
	if err != nil {
		log.Panic(err)
	}
}
