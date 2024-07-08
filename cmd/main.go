package main

import (
	"fmt"
	"log"
	"net/http"

	"broker-service/cmd/config"
	"broker-service/cmd/controller"
	"broker-service/cmd/router"
	"broker-service/cmd/service"

	"github.com/go-playground/validator"
	"github.com/rs/cors"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	validate := validator.New()

	//Init Services
	authenticationService := service.NewAuthenticationServiceImpl(validate)
	vocabService := service.NewVocabServiceImpl(validate)
	groupService := service.NewGroupServiceImpl(validate)

	//Init controllers
	authenticationController := controller.NewAuthenticationController(authenticationService)
	vocabController := controller.NewVocabController(vocabService)
	groupController := controller.NewGroupController(groupService, vocabService)

	r := router.NewRouter(authenticationController, vocabController, groupController)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", loadConfig.PORT),
		Handler: handler,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
