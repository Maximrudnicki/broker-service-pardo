package response

type VocabResponse struct {
	ID          uint32 `json:"id"`
	Word        string `json:"word"`
	Definition  string `json:"definition"`
}
