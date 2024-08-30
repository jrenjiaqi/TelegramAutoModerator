package main

import (
	"fmt"

	"github.com/jrenjq/MiniChatSentryBot/repo"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

func main() {
	uri_string := repo.Get_update_uri_from_env_file(".env")
	fmt.Println(uri_string)
	get_response_string := utils.Http_get(&uri_string)
	fmt.Println(get_response_string)
}
