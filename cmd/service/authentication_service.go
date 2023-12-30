package service

import "broker-service/cmd/data/request"

type AuthenticationService interface {
	Login(user request.LoginRequest) (string, error)
	Register(user request.CreateUserRequest) error
}
