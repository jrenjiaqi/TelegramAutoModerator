package top

import (
	"github.com/jrenjq/MiniChatSentryBot/repo"
	"github.com/jrenjq/MiniChatSentryBot/structs"
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
		claude_api_key, claude_system_prompt := repo.Get_claude_info(
			api_env_pathname,
			system_prompt_pathname,
			debug_mode,
		)

		// send message to Claude with its API key and system prompt, get a response.
		_ = repo.Send_claude_query_and_get_response(
			claude_api_key,
			claude_system_prompt,
			debug_mode,
		)
	}
}
