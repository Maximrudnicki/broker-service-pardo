package service

import (
	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
)

type VocabService interface {
	CreateWord(cwr request.CreateWordRequest) error
	DeleteWord(dwr request.DeleteWordRequest) error
	GetWords(vr request.VocabRequest) ([]response.VocabResponse, error)
	FindWord(fwr request.FindWordRequest) (response.VocabResponse, error)
	UpdateWord(uwr request.UpdateWordRequest) error
	UpdateWordStatus(uwsr request.UpdateWordStatusRequest) error
	ManageTrainings(mtr request.ManageTrainingsRequest) error
}
