package request

type CreateWordRequest struct {
	Token      string `json:"token"`
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type DeleteWordRequest struct {
	Token  string `json:"token"`
	WordId uint32 `json:"word_id"`
}

type FindWordRequest struct {
	WordId uint32 `json:"word_id"`
}

type VocabRequest struct {
	TokenType string `json:"token_type"` // Bearer
	Token     string `json:"token"`
}

type UpdateWordRequest struct {
	Token      string `json:"token"`
	WordId     uint32 `json:"word_id"`
	Definition string `json:"definition"`
}

type UpdateWordStatusRequest struct {
	Token     string `json:"token"`
	WordId    uint32 `json:"word_id"`
	IsLearned bool   `json:"is_learned"`
}

type ManageTrainingsRequest struct {
	Token          string `json:"token"`
	TrainingResult bool   `json:"result"`
	Training       string `json:"training"`
	WordId         uint32 `json:"word_id"`
}
