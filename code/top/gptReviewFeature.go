package top

import (
	"fmt"

	"github.com/jrenjq/MiniChatSentryBot/repo"
	"github.com/jrenjq/MiniChatSentryBot/structs"
	"github.com/jrenjq/MiniChatSentryBot/utils"
)

func Gpt_review_feature(
	telegram_bot_updates *structs.Update_response,
	api_env_pathname string,
	uri_env_pathname string,
	system_prompt_pathname string,
	is_logged bool,
) {
	// get config setting(s) for this feature.
	debug_mode, feature_is_on := repo.Get_gpt_review_config_values_from_env_file(
		".conf",
		"GPT_REVIEW_FEATURE_DEBUG_MODE",
		"GPT_REVIEW_FEATURE_ON",
	)
	// configuration file variable decides whether to run this feature or not.
	if feature_is_on {
		// load Claude API key and system prompt string.
		claude_api_uri, claude_api_key, claude_system_prompt := repo.Get_claude_info(
			api_env_pathname,
			system_prompt_pathname,
			debug_mode,
		)

		// extract messages from the monolith Update object.
		messages_slice_ptr := repo.Get_msgs_from_updates(telegram_bot_updates)
		for _, message := range *messages_slice_ptr {
			fmt.Printf("%+v\n", message.Text)
		}

		// send message to Claude with its API key and system prompt, get a response.
		utils.Send_to_claude_one_message(
			claude_api_uri,
			claude_api_key,
			claude_system_prompt,
			"hello world",
			10,
			debug_mode,
		)
	}
}
