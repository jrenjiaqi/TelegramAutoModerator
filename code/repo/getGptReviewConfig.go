package repo

import (
	"log"
	"strconv"

	"github.com/jrenjq/MiniChatSentryBot/utils"
)

/*
Returns configuration values from the environment file.

parameters:
  - env_file string: name of environment file that contains needed info.
  - debug_mode_var_name string: name of variable in said environment file for debug mode.
  - gpt_review_feature_on_var_name string: name of variable in said environment file for feature on.

returns:
  - bool: debug mode
  - bool: feature on
*/
func Get_gpt_review_config_values_from_env_file(
	env_file string,
	debug_mode_var_name string,
	gpt_review_feature_on_var_name string,
) (bool, bool) {
	utils.Load_env_file(env_file)
	debug_mode, err := strconv.ParseBool(utils.Get_env_value_or_err(debug_mode_var_name))
	if err != nil {
		log.Panic(err)
	}
	feature_on, err := strconv.ParseBool(utils.Get_env_value_or_err(gpt_review_feature_on_var_name))
	if err != nil {
		log.Panic(err)
	}
	return debug_mode, feature_on
}
