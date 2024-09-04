package main

import (
	"log"

	"github.com/jrenjq/MiniChatSentryBot/repo"
)

func main() {
	uri_string := repo.Get_update_uri_from_env_file(".env")
	log.Println(uri_string)
	get_response_JSON := repo.Get_update_JSON_from_URI(uri_string)
	log.Println(get_response_JSON)
	if !get_response_JSON.Ok {
		log.Fatal("Get update failed; something external to this program is not working.")
	}
	messages_to_delete := repo.Get_messages_to_delete_from_JSON(get_response_JSON, 2, true)
	log.Println(messages_to_delete)
}
