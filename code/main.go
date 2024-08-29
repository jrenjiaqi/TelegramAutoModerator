package main

import (
	"fmt"

	"github.com/jrenjq/MiniChatSentryBot/utils"
)

func main() {
	var address string = "https://example.com/"
	get_response_string := utils.Http_get(&address)
	fmt.Print(get_response_string)
}
