package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
	"broker-service/cmd/service"
	"broker-service/cmd/utils"

	"github.com/gin-gonic/gin"
)

type VocabController struct {
	vocabService service.VocabService
}

func NewVocabController(service service.VocabService) *VocabController {
	return &VocabController{vocabService: service}
}

func (controller *VocabController) CreateWord(ctx *gin.Context) {
	token, _ := utils.GetToken(ctx)

	req := request.CreateWordRequest{Token: token}
	err := ctx.ShouldBindJSON(&req)
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

	err_cw := controller.vocabService.CreateWord(req)
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
	token, _ := utils.GetToken(ctx)

	wordId := ctx.Param("wordId")
	id, err_id := strconv.Atoi(wordId)

	req := request.DeleteWordRequest{
		Token:  token,
		WordId: uint32(id),
	}

	err_dw := controller.vocabService.DeleteWord(req)
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
	token, _ := utils.GetToken(ctx)

	vocabRequest := request.VocabRequest{
		TokenType: "Bearer",
		Token:     token,
	}

	res, err_words := controller.vocabService.GetWords(vocabRequest)
	if err_words != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot get words",
		}
		log.Printf("err_words: %v", err_words)
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
	token, _ := utils.GetToken(ctx)

	wordId := ctx.Param("wordId")
	id, err_id := strconv.Atoi(wordId)

	req := request.UpdateWordRequest{
		Token:  token,
		WordId: uint32(id),
	}
	err := ctx.ShouldBindJSON(&req)
	if err != nil || err_id != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update word",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err_uw := controller.vocabService.UpdateWord(req)
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

func (controller *VocabController) UpdateWordStatus(ctx *gin.Context) {
	token, _ := utils.GetToken(ctx)

	wordId := ctx.Param("wordId")
	id, err_id := strconv.Atoi(wordId)

	req := request.UpdateWordStatusRequest{
		Token:  token,
		WordId: uint32(id),
	}
	err := ctx.ShouldBindJSON(&req)
	if err != nil || err_id != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update status",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err_uws := controller.vocabService.UpdateWordStatus(req)

	if err_uws != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update status",
		}
		log.Printf("Cannot update status: %v", err_uws)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Status successfully updated!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *VocabController) ManageTrainings(ctx *gin.Context) {
	token, _ := utils.GetToken(ctx)

	wordId := ctx.Param("wordId")
	id, err_id := strconv.Atoi(wordId)

	req := request.ManageTrainingsRequest{
		Token:  token,
		WordId: uint32(id),
	}
	err := ctx.ShouldBindJSON(&req)
	if err != nil || err_id != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot manage training",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err_mt := controller.vocabService.ManageTrainings(req)

	if err_mt != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot manage training",
		}
		log.Printf("Cannot manage training: %v", err_mt)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully managed!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
