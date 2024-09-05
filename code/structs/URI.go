package structs

import "fmt"

// ensure minimally these fields are in the environment file.
type URI_get_updates struct {
	API_URL                string
	BOT_TOKEN              string
	GET_UPDATES_PATH       string
	ALLOWED_UPDATES_NAME   string
	MESSAGE_REACTIONS_NAME string
}

// struct method: creates a URI string from a URI struct.
func (uri_struct *URI_get_updates) Get_update_uri_string_from_struct() string {
	return fmt.Sprintf(
		"%s%s/%s?%s=%s",
		uri_struct.API_URL,
		uri_struct.BOT_TOKEN,
		uri_struct.GET_UPDATES_PATH,
		uri_struct.ALLOWED_UPDATES_NAME,
		uri_struct.MESSAGE_REACTIONS_NAME,
	)
}

// ensure minimally these fields are in the environment file.
type URI_delete_message struct {
	API_URL             string
	BOT_TOKEN           string
	DELETE_MESSAGE_PATH string
	CHAT_ID_NAME        string
	MESSAGE_ID_NAME     string
}

// struct method: creates a URI string from a URI struct.
func (uri_struct *URI_delete_message) Get_base_delete_uri_string_from_struct() string {
	return fmt.Sprintf(
		"%s%s/%s?%s=<chat_id>&%s=<message_id>",
		uri_struct.API_URL,
		uri_struct.BOT_TOKEN,
		uri_struct.DELETE_MESSAGE_PATH,
		uri_struct.CHAT_ID_NAME,
		uri_struct.MESSAGE_ID_NAME,
	)
}
