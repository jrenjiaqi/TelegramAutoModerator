package main

import (
	"github.com/jrenjq/MiniChatSentryBot/top"
)

func main() {
	// it is the calling function's responsibility to ensure environment file is loaded (e.g. by loading it redundantly).
	config_pathname := "./.conf"                                  // general configuration file.
	uri_env_pathname := "./.env"                                  // environment pathname with uri information.
	api_env_pathname := "./.env"                                  // environment pathname with Claude API key.
	system_prompt_pathname := "./claude/prompts/systemPrompt.txt" // system prompt is SENSITIVE and should NOT be made public.
	processed_messages_log := "./processedMessages.log"           // holds a log of already-processed messages from a previous run.

	// get updates (e.g. messages, reactions) about the bot, from the Telegram Bot API.
	// note: should call only once per execution. Telegram Bot API will limit consecutive calls with truncated information.
	telegram_bot_updates := top.Get_updates_from_telegram(
		uri_env_pathname,
		false,
	)

	// v0.1: thumbs down feature: deletes message(s) that exceed(s) configured thumbs down count.
	// (NOTE: DEFUNCT due to undocumented Telegram API changes. No more reactions information from query response.)
	top.Thumbs_down_feature(
		telegram_bot_updates,
		uri_env_pathname,
		config_pathname,
		true,
	)

	// // v0.2: use Claude LLM to review message for scam and inappropriateness.
	top.Gpt_review_feature(
		telegram_bot_updates,
		api_env_pathname,
		uri_env_pathname,
		processed_messages_log,
		system_prompt_pathname,
		true,
	)
}
