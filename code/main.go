package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/jrenjq/MiniChatSentryBot/repo"
	"github.com/jrenjq/MiniChatSentryBot/structs"
)

func main() {
	uri_string := repo.Get_update_uri_from_env_file(".env")
	fmt.Println(uri_string)
	get_response_JSON := repo.Get_update_JSON_from_URI(uri_string)
	fmt.Println(get_response_JSON)
	if !get_response_JSON.Ok {
		log.Fatal("Get update failed; cause is beyond program control.")
	}
	last_24h_updates := get_response_JSON.Result   // Telegram only gives bots last 24 hour's updates.
	thumbs_down_message_uids := []string{}         // message with thumbs down = add one instance of message id in. vice versa.
	for _, update_item := range last_24h_updates { // iterate through updates from last 24 hours.
		// thumbs down action on a message => add one count of message id to thumbs down message id slice.
		if len(update_item.MessageReaction.NewReaction) != 0 && // update item is a message reaction.
			update_item.MessageReaction.NewReaction[0].Emoji == "👎" { // reaction is a thumbs down.

			thumbs_down_message := structs.Telegram_message_id{Chat_id: update_item.MessageReaction.Chat.ID, Message_id: update_item.MessageReaction.MessageID}
			thumbs_down_message_uid := thumbs_down_message.Get_uid_string()
			log.Printf("👎 detected! adding message %s to slice!\n", thumbs_down_message_uid)
			thumbs_down_message_uids = append(thumbs_down_message_uids, thumbs_down_message_uid)
		}
		// thumbs down action was retracted => remove one count of message id to thumbs down message id slice.
		if len(update_item.MessageReaction.OldReaction) != 0 && // update item is a message reaction RETRACTION.
			update_item.MessageReaction.OldReaction[0].Emoji == "👎" { // retracted reaction is a thumbs down.

			thumbs_down_message := structs.Telegram_message_id{Chat_id: update_item.MessageReaction.Chat.ID, Message_id: update_item.MessageReaction.MessageID}
			thumbs_down_message_uid := thumbs_down_message.Get_uid_string()
			log.Printf("👎 RETRACTION detected! removing message id %+v from slice!\n", thumbs_down_message_uid)
			index_of_msg_id_to_delete := slices.Index(thumbs_down_message_uids, thumbs_down_message_uid) // returns first occurrence of uid in slice. -1 if none.
			if index_of_msg_id_to_delete != -1 {                                                         // possible that there is none, due to 24h limit on updates.
				thumbs_down_message_uids = slices.Delete(thumbs_down_message_uids, index_of_msg_id_to_delete, index_of_msg_id_to_delete+1) // delete one count from slice.
			}
		}
	}
	fmt.Printf("thumbs_down_message_uids: %v\n", thumbs_down_message_uids)
}
