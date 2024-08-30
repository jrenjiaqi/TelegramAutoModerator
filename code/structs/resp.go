package structs

// JSON -> Go Struct: https://transform.tools/json-to-go

type Update_response struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message,omitempty"`
		MessageReaction struct {
			Chat struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
			MessageID int `json:"message_id"`
			User      struct {
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"user"`
			Date        int           `json:"date"`
			OldReaction []interface{} `json:"old_reaction"`
			NewReaction []struct {
				Type  string `json:"type"`
				Emoji string `json:"emoji"`
			} `json:"new_reaction"`
		} `json:"message_reaction,omitempty"`
	} `json:"result"`
}
