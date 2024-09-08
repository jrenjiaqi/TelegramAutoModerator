package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func Send_to_claude_one_message(
	claude_api_uri string,
	claude_api_key string,
	system_prompt string,
	text_to_send string,
	response_timeout_seconds int,
	debug_mode bool,
) (string, error) {
	var err error // variable to hold potential error.

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
	json.NewEncoder(json_byte_buffer).Encode(claude_payload)
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
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return "response", err
}
