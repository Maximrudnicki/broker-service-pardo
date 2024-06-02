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
	vocabService service.VocabService
}

func NewGroupController(service service.GroupService, vs service.VocabService) *GroupController {
	return &GroupController{groupService: service, vocabService: vs}
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

	groupResponse, err_fg := controller.groupService.FindGroup(fgr)
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

	students := make([]response.StudentInformation, 0, len(groupResponse.Students))
	for _, studentId := range groupResponse.Students {
		fsr := request.FindStudentRequest{
			Token:     token,
			StudentId: studentId,
			GroupId:   fgr.GroupId,
		}
		gsr := request.GetStatisticsRequest{
			Token:     token,
			StudentId: studentId,
			GroupId:   fgr.GroupId,
		}

		StatResp, err_gs := controller.groupService.GetStatistics(gsr)
		if err_gs != nil {
			webResponse := response.Response{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Cannot get stats",
			}
			log.Printf("Cannot get stats: %v", err_gs)
			ctx.JSON(http.StatusBadRequest, webResponse)
			return
		}

		words := make([]response.VocabResponse, 0, len(StatResp.Words))
		for _, wordId := range StatResp.Words {
			fwr := request.FindWordRequest{
				WordId: wordId,
			}
			word, err := controller.vocabService.FindWord(fwr)
			if err != nil {
				webResponse := response.Response{
					Code:    http.StatusInternalServerError,
					Status:  "Internal Server Error",
					Message: "Cannot find word",
				}
				log.Printf("Cannot find word with id %d: %v", wordId, err)
				ctx.JSON(http.StatusInternalServerError, webResponse)
				return
			}
			if word.ID != 0 {
				words = append(words, word)
			}
		}

		student, err := controller.groupService.FindStudent(fsr)
		studentInfo := response.StudentInformation{
			StudentId: studentId,
			Email:     student.Email,
			Username:  student.Username,
			Words:     words,
		}
		if err != nil {
			webResponse := response.Response{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "Cannot find student",
			}
			log.Printf("cannot find studnent with id %v: %v", student, err)
			ctx.JSON(http.StatusInternalServerError, webResponse)
			return
		}
		students = append(students, studentInfo)
	}

	res := struct {
		UserId   uint32                        `json:"user_id"`
		GroupId  string                        `json:"group_id"`
		Title    string                        `json:"title"`
		Students []response.StudentInformation `json:"students"`
	}{
		UserId:   groupResponse.UserId,
		GroupId:  groupResponse.GroupId,
		Title:    groupResponse.Title,
		Students: students,
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

func (controller *GroupController) FindTeacher(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	ftr := request.FindTeacherRequest{Token: token}
	ctx.ShouldBindJSON(&ftr)

	res, err_ft := controller.groupService.FindTeacher(ftr)
	if err_ft != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find group",
		}
		log.Printf("Cannot finds group: %v", err_ft)
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

func (controller *GroupController) GetStatistics(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(400, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := authorizationHeader[len("Bearer "):]

	gsr := request.GetStatisticsRequest{Token: token}
	ctx.ShouldBindJSON(&gsr)

	StatResp, err_gs := controller.groupService.GetStatistics(gsr)
	if err_gs != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot get stats",
		}
		log.Printf("Cannot get stats: %v", err_gs)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	words := make([]response.VocabResponse, 0, len(StatResp.Words))
	for _, wordId := range StatResp.Words {
		fwr := request.FindWordRequest{
			WordId: wordId,
		}
		word, err := controller.vocabService.FindWord(fwr)
		if err != nil {
			webResponse := response.Response{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "Cannot find word",
			}
			log.Printf("Cannot find word with id %d: %v", wordId, err)
			ctx.JSON(http.StatusInternalServerError, webResponse)
			return
		}
		if word.ID != 0 {
			words = append(words, word)
		}
	}

	student, err := controller.groupService.FindStudent(request.FindStudentRequest{
		Token: token, StudentId: StatResp.StudentId, GroupId: StatResp.GroupId,
	})
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Cannot find student",
		}
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	res := struct {
		StatId    string                   `json:"statistics_id"`
		GroupId   string                   `json:"group_id"`
		TeacherId uint32                   `json:"teacher_id"`
		Student   response.StudentInfo     `json:"student"`
		Words     []response.VocabResponse `json:"words"`
	}{
		StatId:    StatResp.StatId,
		GroupId:   StatResp.GroupId,
		TeacherId: StatResp.TeacherId,
		Student: response.StudentInfo{
			StudentId: StatResp.StudentId,
			Email:     student.Email,
			Username:  student.Username,
		},
		Words: words,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully found student!",
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
