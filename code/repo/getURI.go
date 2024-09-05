package repo

import (
	"github.com/jrenjq/MiniChatSentryBot/structs"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

/*
Returns update URI string, constructed from the environment file.

parameters:
  - env_file string: name of environment file that contains needed info.

returns:
  - string: the update URI string built from URI struct.
*/
func Get_update_uri_from_env_file(env_file string) string {
	utils.Load_env_file(env_file)
	uri_struct := structs.URI_get_updates{
		API_URL:                utils.Get_env_value_or_err("API_URL"),
		BOT_TOKEN:              utils.Get_env_value_or_err("BOT_TOKEN"),
		GET_UPDATES_PATH:       utils.Get_env_value_or_err("GET_UPDATES_PATH"),
		ALLOWED_UPDATES_NAME:   utils.Get_env_value_or_err("ALLOWED_UPDATES_NAME"),
		MESSAGE_REACTIONS_NAME: utils.Get_env_value_or_err("MESSAGE_REACTIONS_NAME"),
	}
	return uri_struct.Get_update_uri_string_from_struct()
}

/*
Returns delete URI BASE string (without chat id and message id) constructed from the environment file.

Requires further replacement of <chat_id> and <message_id>. These change, depending on which message to delete.

parameters:
  - env_file string: name of environment file that contains needed info.

returns:
  - string: the BASE delete URI string built from URI struct (still needs replacing of <chat_id> and <message_id>).
*/
func Get_delete_base_uri_from_env_file(env_file string) string {
	utils.Load_env_file(env_file)
	uri_struct := structs.URI_delete_message{
		API_URL:             utils.Get_env_value_or_err("API_URL"),
		BOT_TOKEN:           utils.Get_env_value_or_err("BOT_TOKEN"),
		DELETE_MESSAGE_PATH: utils.Get_env_value_or_err("DELETE_MESSAGE_PATH"),
		CHAT_ID_NAME:        utils.Get_env_value_or_err("CHAT_ID_NAME"),
		MESSAGE_ID_NAME:     utils.Get_env_value_or_err("MESSAGE_ID_NAME"),
	}
	return uri_struct.Get_base_delete_uri_string_from_struct()
}
