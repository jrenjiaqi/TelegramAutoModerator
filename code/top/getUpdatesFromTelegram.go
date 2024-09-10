package top

import (
	"log"

	"github.com/jrenjq/MiniChatSentryBot/repo"
	"github.com/jrenjq/MiniChatSentryBot/structs"
)

func Get_updates_from_telegram(uri_env_pathname string, isLogged bool) *structs.Update_response {
	// assemble URI string from environment file.
	uri_string := repo.Get_update_uri_from_env_file(uri_env_pathname)
	if isLogged {
		log.Println(uri_string)
	}
	// get bot update JSON via HTTP GET to URI string.
	get_response_JSON := repo.Get_update_JSON_from_URI(uri_string)
	if !get_response_JSON.Ok {
		log.Fatal("Get update failed; something external to this program is not working.")
	}
	if isLogged {
		log.Printf("FROM TELEGRAM BOT API: %+v\n", get_response_JSON)
	}
	return get_response_JSON
}
