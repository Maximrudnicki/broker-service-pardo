package utils

import (
	"broker-service/cmd/data/request"
	pb "broker-service/proto"
	"context"
	"fmt"
	"io"
	"log"
)

func CreateWord(v pb.VocabServiceClient, cwr request.CreateWordRequest) error {
	log.Println("---Delete was invoked---")

	req := &pb.CreateRequest{
		Token:      cwr.Token,
		Word:       cwr.Word,
		Definition: cwr.Definition,
	}

	_, err := v.CreateWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("Error happened while adding new word: %v\n", err)
	}

	log.Println("Word Created!")
	return nil
}

func DeleteWord(v pb.VocabServiceClient, dwr request.DeleteWordRequest) error {
	log.Println("---Delete was invoked---")

	req := &pb.DeleteRequest{
		Token:  dwr.Token,
		WordId: dwr.WordId,
	}

	_, err := v.DeleteWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("Error happened while deleting the word: %v\n", err)
	}

	log.Println("Word Deleted!")
	return nil
}

func GetWords(v pb.VocabServiceClient, token string) ([]*pb.VocabResponse, error) {
	log.Println("---Vocab was invoked---")

	var vs []*pb.VocabResponse // vocab slice

	req := &pb.VocabRequest{
		TokenType: "Bearer",
		Token:     token,
	}

	stream, err := v.GetWords(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("Error reading from stream: %v", err)
		// log.Fatalf("Error happened while reading vocablary: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("Something happened with stream %v\n", err)
			// log.Fatalf("Something happened: %v\n", err)
		}

		vs = append(vs, res)
	}

	return vs, nil
}

func UpdateWord(v pb.VocabServiceClient, uwr request.UpdateWordRequest) error {
	log.Println("---Update was invoked---")

	req := &pb.UpdateRequest{
		Token:      uwr.Token,
		Id:         uwr.WordId,
		Definition: uwr.Definition,
	}

	_, err := v.UpdateWord(context.Background(), req)
	if err != nil {
		return fmt.Errorf("Error happened while updating the word: %v\n", err)
	}

	log.Println("Word Updated!")
	return nil
}
