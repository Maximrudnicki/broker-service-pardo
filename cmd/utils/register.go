package utils

import (
	pb "broker-service/proto"
	"context"
	"errors"
	"log"
)

func Register(c pb.AuthenticationServiceClient, username string, email string, password string) error {
	req := &pb.RegisterRequest{
		Username: username,
		Email: email,
		Password: password,
	}	
	
	_, err := c.Register(context.Background(), req)
	if err != nil {
		log.Printf("error happened while register: %v", err)
		return errors.New("error happened while register")
	}

	return nil
}
