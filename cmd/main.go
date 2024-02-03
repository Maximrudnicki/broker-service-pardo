package main

import (
	"fmt"
	"log"
	"net/http"

	"broker-service/cmd/controller"
	"broker-service/cmd/router"
	"broker-service/cmd/service"

	"github.com/go-playground/validator"
	"github.com/rs/cors"
)

const (
	webPort   = "8000"
)

func main() {
	validate := validator.New()

	//Init Services
	authenticationService := service.NewAuthenticationServiceImpl(validate)
	vocabService := service.NewVocabServiceImpl(validate)

	//Init controllers
	authenticationController := controller.NewAuthenticationController(authenticationService)
	vocabController := controller.NewVocabController(vocabService)

	r := router.NewRouter(authenticationController, vocabController)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
