package repo

import (
	"fmt"

	"github.com/jrenjq/MiniChatSentryBot/structs"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

func get_uri_string_from_struct(uri_struct_ptr *structs.URI) string {
	return fmt.Sprintf(
		"%s%s/%s?%s=%s",
		uri_struct_ptr.API_URL,
		uri_struct_ptr.BOT_TOKEN,
		uri_struct_ptr.GET_UPDATES_PATH,
		uri_struct_ptr.ALLOWED_UPDATES_NAME,
		uri_struct_ptr.MESSAGE_REACTIONS_NAME,
	)
}

func Get_update_uri_from_env_file(env_file string) string {
	utils.Load_env_file(".env")
	uri_struct := structs.URI{
		API_URL:                utils.Get_env_value_or_err("API_URL"),
		BOT_TOKEN:              utils.Get_env_value_or_err("BOT_TOKEN"),
		GET_UPDATES_PATH:       utils.Get_env_value_or_err("GET_UPDATES_PATH"),
		ALLOWED_UPDATES_NAME:   utils.Get_env_value_or_err("ALLOWED_UPDATES_NAME"),
		MESSAGE_REACTIONS_NAME: utils.Get_env_value_or_err("MESSAGE_REACTIONS_NAME"),
	}
	return get_uri_string_from_struct(&uri_struct)
}
