package repo

import (
	"strconv"

	"github.com/jrenjq/MiniChatSentryBot/utils"
)

/*
Returns configuration values from the environment file.

parameters:
  - env_file string: name of environment file that contains needed info.
  - thumbs_down_count_var_name string: name of variable in said environment file for thumbs down count.
  - debug_mode_var_name string: name of variable in said environment file for debug mode.

returns:
  - int: thumbs down threshold for message to be deleted
  - bool: debug mode
*/
func Get_config_values_from_env_file(
	env_file string,
	thumbs_down_count_var_name string,
	debug_mode_var_name string,
) (int, bool) {
	utils.Load_env_file(env_file)
	thumbs_down_count, err := strconv.Atoi(utils.Get_env_value_or_err(thumbs_down_count_var_name))
	if err != nil {
		panic(err)
	}
	debug_mode, err := strconv.ParseBool(utils.Get_env_value_or_err(debug_mode_var_name))
	if err != nil {
		panic(err)
	}
	return thumbs_down_count, debug_mode
}
