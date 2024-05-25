package controller

import (
	"log"
	"net/http"

	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
	"broker-service/cmd/service"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupService service.GroupService
}

func NewGroupController(service service.GroupService) *GroupController {
	return &GroupController{groupService: service}
}

func (controller *GroupController) AddStudent(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	asr := request.AddStudentRequest{Token: token}
	err := ctx.ShouldBindJSON(&asr)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add student",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	
	err_as := controller.groupService.AddStudent(asr)
	if err_as != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add student",
		}
		log.Printf("Cannot add student: %v", err_as)
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

func (controller *GroupController) AddWordToUser(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	awur := request.AddWordToUserRequest{Token: token}
	ctx.ShouldBindJSON(&awur)

	res, err_fg := controller.groupService.AddWordToUser(awur)
	if err_fg != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add word",
		}
		log.Printf("Cannot add word: %v", err_fg)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully added word!",
		Data:    res,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *GroupController) CreateGroup(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	cgr := request.CreateGroupRequest{Token: token}
	err := ctx.ShouldBindJSON(&cgr)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot create group",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err_cg := controller.groupService.CreateGroup(cgr)
	if err_cg != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot create group",
		}
		log.Printf("Cannot create: %v", err_cg)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *GroupController) DeleteGroup(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	dgr := request.DeleteGroupRequest{
		Token:   token,
		GroupId: ctx.Param("groupId"),
	}

	err_dg := controller.groupService.DeleteGroup(dgr)
	if err_dg != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add word",
		}
		log.Printf("Cannot delete group: %v", err_dg)
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

func (controller *GroupController) FindGroup(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	fgr := request.FindGroupRequest{Token: token}
	ctx.ShouldBindJSON(&fgr)

	res, err_fg := controller.groupService.FindGroup(fgr)
	if err_fg != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find group",
		}
		log.Printf("Cannot finds group: %v", err_fg)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully found groups!",
		Data:    res,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *GroupController) FindStudent(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	fsr := request.FindStudentRequest{Token: token}
	ctx.ShouldBindJSON(&fsr)

	res, err_fs := controller.groupService.FindStudent(fsr)
	if err_fs != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find group",
		}
		log.Printf("Cannot finds group: %v", err_fs)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully found student!",
		Data:    res,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *GroupController) FindGroupsTeacher(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	fgtr := request.FindGroupsTeacherRequest{Token: token}

	res, err_fgt := controller.groupService.FindGroupsTeacher(fgtr)
	if err_fgt != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find groups",
		}
		log.Printf("Cannot finds groups: %v", err_fgt)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully found groups!",
		Data:    res,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *GroupController) FindGroupsStudent(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	fgsr := request.FindGroupsStudentRequest{Token: token}

	res, err_fgs := controller.groupService.FindGroupsStudent(fgsr)
	if err_fgs != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find groups",
		}
		log.Printf("Cannot finds groups: %v", err_fgs)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully found groups!",
		Data:    res,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *GroupController) RemoveStudent(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	rsr := request.RemoveStudentRequest{Token: token}
	err := ctx.ShouldBindJSON(&rsr)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add student",
		}
		log.Printf("Cannot bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err_rs := controller.groupService.RemoveStudent(rsr)
	if err_rs != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot remove student",
		}
		log.Printf("Cannot remove student: %v", err_rs)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully removed!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
