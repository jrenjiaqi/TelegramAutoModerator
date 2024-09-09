package structs

type Data struct {
	UpdateId        int             `json:"update_id"`
	Message         MessageObject   `json:"message,omitempty"`
	MessageReaction MessageReaction `json:"message_reaction,omitempty"`
}

type MessageObject struct {
	MessageId       int               `json:"message_id"`
	From            From              `json:"from"`
	Chat            Chat              `json:"chat"`
	Date            int               `json:"date"`
	ForwardOrigin   ForwardOrigin     `json:"forward_origin"`
	ForwardFrom     ForwardFrom       `json:"forward_from"`
	ForwardDate     int               `json:"forward_date"`
	Text            string            `json:"text,omitempty"`
	PhotoEntities   []PhotoEntities   `json:"photo,omitempty"`
	Caption         string            `json:"caption,omitempty"`
	CaptionEntities []CaptionEntities `json:"caption_entities,omitempty"`
}

type PhotoEntities struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type CaptionEntities struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type From struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type User struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type SenderUser struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Chat struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type ForwardOrigin struct {
	Type       string     `json:"type"`
	SenderUser SenderUser `json:"sender_user"`
	Date       int        `json:"date"`
}

type ForwardFrom struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type MessageReaction struct {
	Chat        Chat          `json:"chat"`
	MessageId   int           `json:"message_id"`
	User        User          `json:"user"`
	Date        int           `json:"date"`
	OldReaction []OldReaction `json:"old_reaction"`
	NewReaction []NewReaction `json:"new_reaction"`
}

type OldReaction struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

type NewReaction struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}
