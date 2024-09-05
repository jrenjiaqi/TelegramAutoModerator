package repo

import (
	"log"
	"strconv"
	"strings"

	"github.com/jrenjq/MiniChatSentryBot/structs"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

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

func Delete_messages(
	delete_uri_env_file_path string,
	messages_to_delete *[]structs.Telegram_message_id,
) (int, error) {
	chat_id_placeholder := "<chat_id>"
	message_id_placeholder := "<message_id>"
	delete_base_uri_string := Get_delete_base_uri_from_env_file(delete_uri_env_file_path)
	for _, message_to_delete := range *messages_to_delete {
		delete_uri_string := build_delete_base_uri_string(
			delete_base_uri_string,
			chat_id_placeholder,
			message_to_delete.Chat_id,
			message_id_placeholder,
			message_to_delete.Message_id,
		)
		log.Println(delete_uri_string)
		delete_JSON_struct_ptr := new(structs.Delete_response) // or &structs.Update_response{}
		utils.Http_GET_JSON(delete_uri_string, 5, delete_JSON_struct_ptr)
		log.Println(delete_JSON_struct_ptr)
	}
	return 0, nil
}
