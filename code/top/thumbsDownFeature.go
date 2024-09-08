package top

import (
	"log"

	"github.com/jrenjq/MiniChatSentryBot/repo"
	"github.com/jrenjq/MiniChatSentryBot/structs"
)

func Thumbs_down_feature(
	telegram_bot_updates *structs.Update_response,
	uri_env_pathname string,
	config_pathname string,
	isLogged bool,
) {
	// get configuration values from config file.
	thumbs_down_count_to_del_msg, debug_mode, feature_is_on :=
		repo.Get_thumbs_down_config_values_from_env_file(
			config_pathname,
			"THUMBS_DOWN_COUNT_TO_DELETE_MSG",
			"THUMBS_DOWN_FEATURE_DEBUG_MODE",
			"THUMBS_DOWN_FEATURE_ON",
		)
	// configuration file variable decides whether to run this feature or not.
	if feature_is_on {
		// determine messages to delete based on net thumbs down reactions count.
		messages_to_delete := repo.Get_messages_to_delete_from_JSON(
			telegram_bot_updates,
			thumbs_down_count_to_del_msg,
			debug_mode,
		)
		if isLogged {
			log.Printf("THUMBS DOWN MESSAGE(S) TO DELETE: %+v\n", messages_to_delete)
		}
		// delete said messages via Telegram Bot API.
		count, err := repo.Delete_messages(uri_env_pathname, &messages_to_delete, debug_mode)
		if err != nil {
			panic(err)
		}
		if isLogged {
			log.Printf("DELETED %d THUMBS DOWN MESSAGE(S).\n", count)
		}
	}
}
