package service

import (
	"errors"
	"log"

	"broker-service/cmd/config"
	"broker-service/cmd/data/request"
	u "broker-service/cmd/utils"
	pb "broker-service/proto"

	"github.com/go-playground/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthenticationServiceImpl struct {
	Validate *validator.Validate
}

func NewAuthenticationServiceImpl(validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		Validate: validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(user request.LoginRequest) (string, error) {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// connect to auth service as a client
	conn, err := grpc.NewClient(loadConfig.AUTH_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAuthenticationServiceClient(conn)

	loginResponse, err := u.Login(c, user.Email, user.Password)
	if err != nil {
		return "", errors.New("invalid username or Password")
	}

	return loginResponse.Token, nil
}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(user request.CreateUserRequest) error {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// connect to auth service as a client
	conn, err := grpc.NewClient(loadConfig.AUTH_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthenticationServiceClient(conn)
	
	register_err := u.Register(c, user.Username, user.Email, user.Password)
	if register_err != nil {
		return errors.New("cannot register")
	}

	return nil
}
