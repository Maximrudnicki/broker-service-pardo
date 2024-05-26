package utils

import (
	"broker-service/cmd/data/request"
	pb "broker-service/proto"
	"context"
	"fmt"
	"io"
)

func CreateWord(v pb.VocabServiceClient, cwr request.CreateWordRequest) error {
	req := &pb.CreateRequest{
		Token:      cwr.Token,
		Word:       cwr.Word,
		Definition: cwr.Definition,
	}

	_, err := v.CreateWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while adding new word: %v", err)
	}

	return nil
}

func FindWord(v pb.VocabServiceClient, fwr request.FindWordRequest) (*pb.VocabResponse, error) {
	req := &pb.WordRequest{
		WordId: fwr.WordId,
	}

	word, err := v.FindWord(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened while finding the word: %v", err)
	}

	return word, nil
}

func DeleteWord(v pb.VocabServiceClient, dwr request.DeleteWordRequest) error {
	req := &pb.DeleteRequest{
		Token:  dwr.Token,
		WordId: dwr.WordId,
	}

	_, err := v.DeleteWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while deleting the word: %v", err)
	}

	return nil
}

func GetWords(v pb.VocabServiceClient, token string) ([]*pb.VocabResponse, error) {
	var vs []*pb.VocabResponse // vocab slice

	req := &pb.VocabRequest{
		TokenType: "Bearer",
		Token:     token,
	}

	stream, err := v.GetWords(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error reading from stream: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("something happened with stream %v", err)
		}

		vs = append(vs, res)
	}

	return vs, nil
}

func UpdateWord(v pb.VocabServiceClient, uwr request.UpdateWordRequest) error {
	req := &pb.UpdateRequest{
		Token:      uwr.Token,
		Id:         uwr.WordId,
		Definition: uwr.Definition,
	}

	_, err := v.UpdateWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while updating the word: %v", err)
	}

	return nil
}


func UpdateWordStatus(v pb.VocabServiceClient, uwsr request.UpdateWordStatusRequest) error {
	req := &pb.UpdateStatusRequest{
		Token:     uwsr.Token,
		IsLearned: uwsr.IsLearned,
		Id:        uwsr.WordId,
	}

	_, err := v.UpdateWordStatus(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while updating the word: %v", err)
	}

	return nil
}

func ManageTrainings(v pb.VocabServiceClient, mtr request.ManageTrainingsRequest) error {
	req := &pb.ManageTrainingsRequest{
		Token:    mtr.Token,
		Training: mtr.Training,
		Res:      mtr.TrainingResult,
		Id:       mtr.WordId,
	}

	_, err := v.ManageTrainings(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while managing trainings the word: %v", err)
	}

	return nil
}
