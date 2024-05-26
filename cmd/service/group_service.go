package service

import (
	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
)

type GroupService interface {
	AddStudent(asr request.AddStudentRequest) error
	AddWordToUser(awur request.AddWordToUserRequest) (response.AddWordToUserResponse, error)
	CreateGroup(cgr request.CreateGroupRequest) error
	DeleteGroup(dgr request.DeleteGroupRequest) error
	FindGroup(fgr request.FindGroupRequest) (response.GroupResponse , error)
	FindStudent(fsr request.FindStudentRequest) (response.StudentResponse, error)
	FindGroupsTeacher(fgtr request.FindGroupsTeacherRequest) ([]response.GroupResponse , error)
	FindGroupsStudent(fgsr request.FindGroupsStudentRequest) ([]response.GroupResponse , error)
	GetStatistics(gsr request.GetStatisticsRequest) (response.StatisticsResponse, error)
	RemoveStudent(rsr request.RemoveStudentRequest) error
}
