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

}

func (controller *GroupController) FindGroup(ctx *gin.Context) {

}

func (controller *GroupController) FindGroupsTeacher(ctx *gin.Context) {

}

func (controller *GroupController) FindGroupsStudent(ctx *gin.Context) {

}

func (controller *GroupController) RemoveStudent(ctx *gin.Context) {

}
