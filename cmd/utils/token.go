package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetToken(ctx *gin.Context) (string, error) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return "", errors.New("authorization header is missing")
	}

	token := authorizationHeader[len("Bearer "):]

	return token, nil
}
