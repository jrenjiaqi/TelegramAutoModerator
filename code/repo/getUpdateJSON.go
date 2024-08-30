package repo

import (
	"github.com/jrenjq/MiniChatSentryBot/structs"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

/*
Gets bot updates from Telegram API server.

parameters:
  - uri string: URI to send the GET request.

returns:
  - *structs.Update_response: pointer to the struct which represents the JSON response.
*/
func Get_update_JSON_from_URI(uri string) *structs.Update_response {
	update_JSON_struct_ptr := new(structs.Update_response) // or &structs.Update_response{}
	utils.Http_GET_JSON(uri, 5, update_JSON_struct_ptr)
	return update_JSON_struct_ptr
}
