package service

import (
	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
)

type VocabService interface {
	CreateWord(createWordRequest request.CreateWordRequest) error
	DeleteWord(deleteWordRequest request.DeleteWordRequest) error
	GetWords(vocabRequest request.VocabRequest) ([]response.VocabResponse, error)
	FindWord(findWordRequest request.FindWordRequest) (response.VocabResponse, error)
	UpdateWord(updateWordRequest request.UpdateWordRequest) error
	UpdateWordStatus(updateWordStatusRequest request.UpdateWordStatusRequest) error
	ManageTrainings(manageTrainingsRequest request.ManageTrainingsRequest) error
}
