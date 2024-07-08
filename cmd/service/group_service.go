package service

import (
	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
)

type GroupService interface {
	AddStudent(addStudentRequest request.AddStudentRequest) error
	AddWordToUser(addWordToUserRequest request.AddWordToUserRequest) (response.AddWordToUserResponse, error)
	CreateGroup(createGroupRequest request.CreateGroupRequest) error
	DeleteGroup(deleteGroupRequest request.DeleteGroupRequest) error
	FindGroup(findGroupRequest request.FindGroupRequest) (response.GroupResponse, error)
	FindStudent(findStudentRequest request.FindStudentRequest) (response.StudentResponse, error)
	FindTeacher(findTeacherRequest request.FindTeacherRequest) (response.TeacherResponse, error)
	FindGroupsTeacher(findGroupsTeacherRequest request.FindGroupsTeacherRequest) ([]response.GroupResponse , error)
	FindGroupsStudent(findGroupsStudentRequest request.FindGroupsStudentRequest) ([]response.GroupResponse , error)
	GetStatistics(getStatisticsRequest request.GetStatisticsRequest) (response.StatisticsResponse, error)
	RemoveStudent(removeStudentRequest request.RemoveStudentRequest) error
}
