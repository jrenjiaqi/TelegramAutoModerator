package main

import (
	"fmt"

	"github.com/jrenjq/MiniChatSentryBot/repo"
)

func main() {
	uri_string := repo.Get_update_uri_from_env_file(".env")
	fmt.Println(uri_string)
	get_response_string := repo.Get_update_JSON_from_URI(uri_string)
	fmt.Println(get_response_string)
}
