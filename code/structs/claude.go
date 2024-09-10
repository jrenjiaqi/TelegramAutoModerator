package structs

// represents the JSON that Claude sends back after HTTP POST on its endpoint.
type ClaudeResponse struct {
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	Role         string          `json:"role"`
	Model        string          `json:"model"`
	Content      []ContentObject `json:"content"`
	StopReason   string          `json:"stop_reason"`
	StopSequence any             `json:"stop_sequence"`
	Usage        UsageObject     `json:"usage"`
}

type ContentObject struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type UsageObject struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// represents the parsed ratings from ContentObject.Text, as defined by system prompt given.
type ClaudeMessageRating struct {
	ScamRatingInt             int    `json:"scam_rating_int"`             // how likely message is to be a scam on a scale of 1-10.
	ScamRatingReason          string `json:"scam_rating_reason"`          // reason for above rating.
	InappropriateRatingInt    int    `json:"inappropriate_rating_int"`    // how inappropriate is the message for the context provided.
	InappropriateRatingReason string `json:"inappropriate_rating_reason"` // reason for above rating.
}
