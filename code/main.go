package main

import (
	"fmt"
	"log"

	"github.com/jrenjq/MiniChatSentryBot/utils"
)

func main() {
	utils.Load_env_file() // .env file must be in same dir as program entrypoint.
	if value, exists := utils.Get_env_value("API_URL"); !exists {
		log.Print("No such environment variable (key name).")
	} else {
		fmt.Printf("API_URL: %s\n", value)
	}
	var address string = "https://example.com/"
	get_response_string := utils.Http_get(&address)
	fmt.Print(get_response_string)
}
