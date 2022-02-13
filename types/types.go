package types

type Blog struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Anous    string `json:"anous"`
	FullText string `json:"full_text"`
	Now      string `json:"time"`
	Username string `json:"username"`
}
