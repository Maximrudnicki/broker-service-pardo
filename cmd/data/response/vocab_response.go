package response

import "google.golang.org/protobuf/types/known/timestamppb"

type VocabResponse struct {
	ID              uint32                 `json:"id"`
	Word            string                 `json:"word"`
	Definition      string                 `json:"definition"`
	CreatedAt       *timestamppb.Timestamp `json:"created_at"`
	IsLearned       bool                   `json:"is_learned"`
	Cards           bool                   `json:"cards"`
	WordTranslation bool                   `json:"word_translation"`
	Constructor     bool                   `json:"constructor"`
	WordAudio       bool                   `json:"word_audio"`
}
