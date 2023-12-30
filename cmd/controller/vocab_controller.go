package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
	"broker-service/cmd/service"

	"github.com/gin-gonic/gin"
)

type VocabController struct {
	vocabService service.VocabService
}

func NewVocabController(service service.VocabService) *VocabController {
	return &VocabController{vocabService: service}
}

func (controller *VocabController) CreateWord(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	cwr := request.CreateWordRequest{Token: token}
	err := ctx.ShouldBindJSON(&cwr)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add word",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	
	err_cw := controller.vocabService.CreateWord(cwr)
	if err_cw != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add word",
		}
		log.Printf("Cannot add: %v", err_cw)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully added!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *VocabController) DeleteWord(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	wordId := ctx.Param("wordId")
	id, err_id := strconv.Atoi(wordId)

	dwr := request.DeleteWordRequest{
		Token:  token,
		WordId: uint32(id),
	}

	err_dw := controller.vocabService.DeleteWord(dwr)
	fmt.Println(err_dw)
	if err_dw != nil || err_id != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot delete word",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *VocabController) GetWords(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	vocabRequest := request.VocabRequest{
		TokenType: "Bearer",
		Token:     token,
	}

	res, err_words := controller.vocabService.GetWords(vocabRequest)
	fmt.Println(err_words)
	if err_words != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot get words",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully got words!",
		Data:    res,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *VocabController) UpdateWord(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	wordId := ctx.Param("wordId")
	id, err_id := strconv.Atoi(wordId)

	uwr := request.UpdateWordRequest{
		Token: token,
		WordId: uint32(id),
	}
	err := ctx.ShouldBindJSON(&uwr)
	if err != nil || err_id != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add word",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	
	err_uw := controller.vocabService.UpdateWord(uwr)
	if err_uw != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update word",
		}
		log.Printf("Cannot update: %v", err_uw)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
