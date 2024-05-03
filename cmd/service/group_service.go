package service

import (
	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
)

type GroupService interface {
	AddStudent(asr request.AddStudentRequest) error
	CreateGroup(cgr request.CreateGroupRequest) error
	DeleteGroup(dgr request.DeleteGroupRequest) error
	FindGroup(fgr request.FindGroupRequest) (response.GroupResponse , error)
	FindGroupsTeacher(fgtr request.FindGroupsTeacherRequest) ([]response.GroupResponse , error)
	FindGroupsStudent(fgsr request.FindGroupsStudentRequest) ([]response.GroupResponse , error)
	RemoveStudent(rsr request.RemoveStudentRequest) error
}
