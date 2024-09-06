package structs

// JSON -> Go Struct: https://transform.tools/json-to-go

type Update_response struct {
	Ok     bool `json:"ok"`
	Result []Data `json:"result"`
}

type Delete_response struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}
