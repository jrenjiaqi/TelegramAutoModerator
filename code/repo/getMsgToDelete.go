package repo

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/jrenjq/MiniChatSentryBot/structs"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

// Helper function: unpacks the uid string into a message id struct.
func get_message_id_struct_from_uid(uid string) structs.Telegram_message_id {
	split_string := strings.Split(uid, ":")
	chat_id, err := strconv.ParseInt(split_string[0], 10, 64)
	if err != nil {
		panic(err)
	}
	message_id, err := strconv.Atoi(split_string[1])
	if err != nil {
		panic(err)
	}
	return structs.Telegram_message_id{Chat_id: chat_id, Message_id: message_id}
}

func Get_messages_to_delete_from_JSON(
	get_response_JSON *structs.Update_response,
	thumbs_down_count_to_del_msg int,
	is_debug_mode bool,
) []structs.Telegram_message_id {
	last_24h_updates := get_response_JSON.Result   // Telegram only gives bots last 24 hour's updates.
	thumbs_down_message_uids := []string{}         // message with thumbs dowgit puln = add one instance of message id in. vice versa.
	for _, update_item := range last_24h_updates { // iterate through updates from last 24 hours.
		// thumbs down action on a message => add one count of message id to thumbs down message id slice.
		if len(update_item.MessageReaction.NewReaction) != 0 && // update item is a message reaction.
			update_item.MessageReaction.NewReaction[0].Emoji == "👎" { // reaction is a thumbs down.

			thumbs_down_message := structs.Telegram_message_id{Chat_id: update_item.MessageReaction.Chat.ID, Message_id: update_item.MessageReaction.MessageId}
			thumbs_down_message_uid := thumbs_down_message.Get_uid_string()
			if is_debug_mode {
				log.Printf("👎 detected! adding message %s to slice!\n", thumbs_down_message_uid)
			}
			thumbs_down_message_uids = append(thumbs_down_message_uids, thumbs_down_message_uid)
		}
		// thumbs down action was retracted => remove one count of message id to thumbs down message id slice.
		if len(update_item.MessageReaction.OldReaction) != 0 && // update item is a message reaction RETRACTION.
			update_item.MessageReaction.OldReaction[0].Emoji == "👎" { // retracted reaction is a thumbs down.

			thumbs_down_message := structs.Telegram_message_id{Chat_id: update_item.MessageReaction.Chat.ID, Message_id: update_item.MessageReaction.MessageId}
			thumbs_down_message_uid := thumbs_down_message.Get_uid_string()
			if is_debug_mode {
				log.Printf("👎 RETRACTION detected! removing message id %+v from slice!\n", thumbs_down_message_uid)
			}
			index_of_msg_id_to_delete := slices.Index(thumbs_down_message_uids, thumbs_down_message_uid) // returns first occurrence of uid in slice. -1 if none.
			if index_of_msg_id_to_delete != -1 {                                                         // possible that there is none, due to 24h limit on updates.
				thumbs_down_message_uids = slices.Delete(thumbs_down_message_uids, index_of_msg_id_to_delete, index_of_msg_id_to_delete+1) // delete one count from slice.
			}
		}
	}
	if is_debug_mode {
		log.Printf("thumbs_down_message_uids: %v\n", thumbs_down_message_uids) // e.g. [-1002237832629:12 -1002237832629:15 -1002237832629:15]
	}
	thumbs_down_message_uid_count_map := utils.Count_string_slice_occurrences(thumbs_down_message_uids) // create a count for how many thumbs down reactions there were.
	messages_to_delete := []structs.Telegram_message_id{}
	for uid, count := range thumbs_down_message_uid_count_map {
		if count >= thumbs_down_count_to_del_msg {
			message_to_delete := get_message_id_struct_from_uid(uid)
			if is_debug_mode {
				log.Printf("chat_id: %d, message_id: %d\n", message_to_delete.Chat_id, message_to_delete.Message_id)
			}
			messages_to_delete = append(messages_to_delete, message_to_delete)
		}
	}
	return messages_to_delete // returns a slice of struct on messages eligible for deletion.
}
