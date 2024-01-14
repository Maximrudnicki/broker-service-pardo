package utils

import (
	pb "broker-service/proto"
	"context"
	"errors"
)

func Login(c pb.AuthenticationServiceClient, email string, password string) (*pb.LoginResponse, error) {
	req := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := c.Login(context.Background(), req)
	if err != nil {
		return nil, errors.New("invalid username or Password")
	}

	return res, nil
}
