package repo

import (
	"log"
	"strconv"
	"strings"

	"github.com/jrenjq/MiniChatSentryBot/structs"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

/*
Helper function: builds the URL query string for the GET HTTP request by replacing placeholder with dynamic values.

parameters:
  - delete_base_uri_string string: the static part of the URL query string.
  - chat_id_placeholder string: the placeholder string value for chat_id, to be replaced.
  - chat_id int64: the Telegram chat id.
  - message_id_placeholder string: the placeholder string value for message_id, to be replaced.
  - message_id int: the Telegram message id.

returrns:
  - string: the complete URL query string to delete the message.
*/
func build_delete_base_uri_string(
	delete_base_uri_string string,
	chat_id_placeholder string,
	chat_id int64,
	message_id_placeholder string,
	message_id int,
) string {
	delete_uri_string := strings.Replace(
		delete_base_uri_string,
		chat_id_placeholder,
		strconv.FormatInt(chat_id, 10),
		1,
	)
	delete_uri_string = strings.Replace(
		delete_uri_string,
		message_id_placeholder,
		strconv.Itoa(message_id),
		1,
	)
	return delete_uri_string
}

/*
Deletes messages.

parameters:
  - delete_uri_env_file_path string,
  - messages_to_delete *[]structs.Telegram_message_id,

returrns:
  - int: number of delete GET requests sent.
  - err: any errors that occurred during delete.
*/
func Delete_messages(
	delete_uri_env_file_path string,
	messages_to_delete *[]structs.Telegram_message_id,
	debug_mode bool,
) (int, error) {
	chat_id_placeholder := "<chat_id>"
	message_id_placeholder := "<message_id>"
	delete_base_uri_string := Get_delete_base_uri_from_env_file(delete_uri_env_file_path)
	count := 0
	err := new(error)
	for index, message_to_delete := range *messages_to_delete {
		delete_uri_string := build_delete_base_uri_string(
			delete_base_uri_string,
			chat_id_placeholder,
			message_to_delete.Chat_id,
			message_id_placeholder,
			message_to_delete.Message_id,
		)
		if debug_mode {
			log.Println(delete_uri_string)
		}
		delete_JSON_struct_ptr := new(structs.Delete_response)                   // or &structs.Update_response{}
		*err = utils.Http_GET_JSON(delete_uri_string, 5, delete_JSON_struct_ptr) // send GET request to delete message.
		if debug_mode {
			log.Println(delete_JSON_struct_ptr)
		}
		count = index
	}
	return count, *err
}
