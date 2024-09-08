package repo

import (
	"log"
	"os"

	"github.com/jrenjq/MiniChatSentryBot/utils"
)

func Get_claude_info(api_env_file string, system_prompt_pathname string, debug_mode bool) (string, string, string) {
	const api_key_last_x_chars int = 8
	utils.Load_env_file(api_env_file) // loading env file is calling function's responsibility.
	claude_api_uri := utils.Get_env_value_or_err("CLAUDE_API_URI")
	claude_api_key := utils.Get_env_value_or_err("CLAUDE_API_KEY")
	system_prompt_bytes, err := os.ReadFile(system_prompt_pathname)
	if err != nil {
		panic(err)
	}
	if debug_mode {
		// log last 8 characters of API key.
		log.Printf("CLAUDE API KEY LAST 8 CHARS: %s\n", claude_api_key[len(claude_api_key)-api_key_last_x_chars:])
		log.Printf("SYSTEM PROMPT: %s\n", string(system_prompt_bytes))
	}
	return claude_api_uri, claude_api_key, string(system_prompt_bytes)
}
