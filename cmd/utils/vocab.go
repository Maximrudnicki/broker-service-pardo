package utils

import (
	"broker-service/cmd/data/request"
	pb "broker-service/proto"
	"context"
	"fmt"
	"io"
)

func CreateWord(v pb.VocabServiceClient, createWordRequest request.CreateWordRequest) error {
	req := &pb.CreateRequest{
		Token:      createWordRequest.Token,
		Word:       createWordRequest.Word,
		Definition: createWordRequest.Definition,
	}

	_, err := v.CreateWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while adding new word: %v", err)
	}

	return nil
}

func FindWord(v pb.VocabServiceClient, findWordRequest request.FindWordRequest) (*pb.VocabResponse, error) {
	req := &pb.WordRequest{
		WordId: findWordRequest.WordId,
	}

	word, err := v.FindWord(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened while finding the word: %v", err)
	}

	return word, nil
}

func DeleteWord(v pb.VocabServiceClient, deleteWordRequest request.DeleteWordRequest) error {
	req := &pb.DeleteRequest{
		Token:  deleteWordRequest.Token,
		WordId: deleteWordRequest.WordId,
	}

	_, err := v.DeleteWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while deleting the word: %v", err)
	}

	return nil
}

func GetWords(v pb.VocabServiceClient, token string) ([]*pb.VocabResponse, error) {
	var vocabSlice []*pb.VocabResponse

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

		vocabSlice = append(vocabSlice, res)
	}

	return vocabSlice, nil
}

func UpdateWord(v pb.VocabServiceClient, updateWordRequest request.UpdateWordRequest) error {
	req := &pb.UpdateRequest{
		Token:      updateWordRequest.Token,
		Id:         updateWordRequest.WordId,
		Definition: updateWordRequest.Definition,
	}

	_, err := v.UpdateWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while updating the word: %v", err)
	}

	return nil
}


func UpdateWordStatus(v pb.VocabServiceClient, updateWordStatusRequest request.UpdateWordStatusRequest) error {
	req := &pb.UpdateStatusRequest{
		Token:     updateWordStatusRequest.Token,
		IsLearned: updateWordStatusRequest.IsLearned,
		Id:        updateWordStatusRequest.WordId,
	}

	_, err := v.UpdateWordStatus(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while updating the word: %v", err)
	}

	return nil
}

func ManageTrainings(v pb.VocabServiceClient, manageTrainingsRequest request.ManageTrainingsRequest) error {
	req := &pb.ManageTrainingsRequest{
		Token:    manageTrainingsRequest.Token,
		Training: manageTrainingsRequest.Training,
		Res:      manageTrainingsRequest.TrainingResult,
		Id:       manageTrainingsRequest.WordId,
	}

	_, err := v.ManageTrainings(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while managing trainings the word: %v", err)
	}

	return nil
}
