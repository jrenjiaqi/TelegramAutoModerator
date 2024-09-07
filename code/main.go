package main

import (
	"log"

	"github.com/jrenjq/MiniChatSentryBot/repo"
)

func main() {
	uri_env_filename := ".env" // environment filename with uri information.

	thumbs_down_count_to_del_msg, debug_mode := repo.Get_config_values_from_env_file(
		".conf",
		"THUMBS_DOWN_COUNT_TO_DELETE_MSG",
		"DEBUG_MODE",
	)

	// assemble URI string from environment file.
	uri_string := repo.Get_update_uri_from_env_file(uri_env_filename)
	log.Println(uri_string)

	// get bot update JSON via HTTP GET to URI string.
	get_response_JSON := repo.Get_update_JSON_from_URI(uri_string)
	if !get_response_JSON.Ok {
		log.Fatal("Get update failed; something external to this program is not working.")
	}
	log.Println(get_response_JSON)

	// determine messages to delete based on net thumbs down reactions count.
	messages_to_delete := repo.Get_messages_to_delete_from_JSON(
		get_response_JSON,
		thumbs_down_count_to_del_msg,
		debug_mode,
	)
	log.Println(messages_to_delete)

	// delete said messages via Telegram Bot API.
	count, err := repo.Delete_messages(uri_env_filename, &messages_to_delete, debug_mode)
	if err != nil {
		panic(err)
	}
	log.Printf("deleted %d messages.\n", count)
}
