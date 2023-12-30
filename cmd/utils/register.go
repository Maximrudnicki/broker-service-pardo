package utils

import (
	pb "broker-service/proto"
	"context"
	"errors"
	"log"
)

func Register(c pb.AuthenticationServiceClient, username string, email string, password string) error {
	log.Println("---Register was invoked---")

	req := &pb.RegisterRequest{
		Username: username,
		Email: email,
		Password: password,
	}	
	
	_, err := c.Register(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while register: %v\n", err)
		return errors.New("Error happened while register")
	}

	log.Println("Registered")
	return nil
}
