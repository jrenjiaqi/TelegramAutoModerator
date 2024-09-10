package structs

import "strconv"

// v0.1: extracts message id for each Telegram message.
type Telegram_message_id struct {
	Chat_id    int64
	Message_id int
}

// v0.1: get a unique message identifier via combining chat_id and message_id strings.
func (msg *Telegram_message_id) Get_uid_string() string {
	BASE := 10
	return strconv.FormatInt(msg.Chat_id, BASE) + ":" + strconv.FormatInt(int64(msg.Message_id), BASE)
}
