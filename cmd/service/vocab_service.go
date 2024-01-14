package service

import (
	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
)

type VocabService interface {
	CreateWord(cwr request.CreateWordRequest) error
	DeleteWord(dwr request.DeleteWordRequest) error
	GetWords(vr request.VocabRequest) ([]response.VocabResponse, error)
	UpdateWord(uwr request.UpdateWordRequest) error
	ManageTrainings(mtr request.ManageTrainingsRequest) error
}
