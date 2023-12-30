package utils

import (
	pb "broker-service/proto"
	"context"
	"errors"
	"log"
)

func Login(c pb.AuthenticationServiceClient, email string, password string) (*pb.LoginResponse, error) {
	log.Println("---Login was invoked---")

	req := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := c.Login(context.Background(), req)
	if err != nil {
		return nil, errors.New("invalid username or Password")
	}

	log.Printf("Login: %v\n", res)
	return res, nil
}
