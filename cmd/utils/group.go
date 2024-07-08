package utils

import (
	"broker-service/cmd/data/request"
	pb "broker-service/proto"
	"context"
	"fmt"
	"io"
)

func AddStudent(g pb.GroupServiceClient, addStudentRequest request.AddStudentRequest) error {
	req := &pb.AddStudentRequest{
		Token:   addStudentRequest.Token,
		GroupId: addStudentRequest.GroupId,
	}

	_, err := g.AddStudent(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while adding student to the group: %v", err)
	}

	return nil
}

func AddWordToUser(g pb.GroupServiceClient, addWordToUserRequest request.AddWordToUserRequest) (*pb.AddWordToUserResponse, error) {
	req := &pb.AddWordToUserRequest{
		Word:       addWordToUserRequest.Word,
		Definition: addWordToUserRequest.Definition,
		GroupId:    addWordToUserRequest.GroupId,
		UserId:     addWordToUserRequest.UserId,
		Token:      addWordToUserRequest.Token,
	}

	res, err := g.AddWordToUser(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened while adding student to the group: %v", err)
	}

	return res, nil
}

func CreateGroup(g pb.GroupServiceClient, createGroupRequest request.CreateGroupRequest) error {
	req := &pb.CreateGroupRequest{
		Title: createGroupRequest.Title,
		Token: createGroupRequest.Token,
	}

	_, err := g.CreateGroup(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while creating group: %v", err)
	}

	return nil
}

func DeleteGroup(g pb.GroupServiceClient, deleteGroupRequest request.DeleteGroupRequest) error {
	req := &pb.DeleteGroupRequest{
		Token:   deleteGroupRequest.Token,
		GroupId: deleteGroupRequest.GroupId,
	}

	_, err := g.DeleteGroup(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened: %v", err)
	}

	return nil
}

func FindGroup(g pb.GroupServiceClient, findGroupRequest request.FindGroupRequest) (*pb.GroupResponse, error) {
	req := &pb.FindGroupRequest{
		Token:   findGroupRequest.Token,
		GroupId: findGroupRequest.GroupId,
	}

	group, err := g.FindGroup(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return group, nil
}

func FindStudent(g pb.GroupServiceClient, findStudentRequest request.FindStudentRequest) (*pb.StudentResponse, error) {
	req := &pb.FindStudentRequest{
		Token:     findStudentRequest.Token,
		StudentId: findStudentRequest.StudentId,
		GroupId:   findStudentRequest.GroupId,
	}

	student, err := g.FindStudent(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return student, nil
}

func FindTeacher(g pb.GroupServiceClient, findTeacherRequest request.FindTeacherRequest) (*pb.TeacherResponse, error) {
	req := &pb.FindTeacherRequest{
		Token:   findTeacherRequest.Token,
		GroupId: findTeacherRequest.GroupId,
	}

	teacher, err := g.FindTeacher(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return teacher, nil
}

func FindGroupsTeacher(g pb.GroupServiceClient, findGroupsTeacherRequest request.FindGroupsTeacherRequest) ([]*pb.GroupResponse, error) {
	req := &pb.FindGroupsTeacherRequest{
		Token: findGroupsTeacherRequest.Token,
	}

	var groupResponseSlice []*pb.GroupResponse

	stream, err := g.FindGroupsTeacher(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error reading from stream: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("something happened with stream %v", err)
		}

		groupResponseSlice = append(groupResponseSlice, res)
	}

	return groupResponseSlice, nil
}

func FindGroupsStudent(g pb.GroupServiceClient, findGroupsStudentRequest request.FindGroupsStudentRequest) ([]*pb.GroupResponse, error) {
	req := &pb.FindGroupsStudentRequest{
		Token: findGroupsStudentRequest.Token,
	}

	var groupResponseSlice []*pb.GroupResponse

	stream, err := g.FindGroupsStudent(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error reading from stream: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("something happened with stream %v", err)
		}

		groupResponseSlice = append(groupResponseSlice, res)
	}

	return groupResponseSlice, nil
}

func GetStatistics(g pb.GroupServiceClient, getStatisticsRequest request.GetStatisticsRequest) (*pb.StatisticsResponse, error) {
	req := &pb.GetStatisticsRequest{
		StudentId: getStatisticsRequest.StudentId,
		GroupId:   getStatisticsRequest.GroupId,
		Token:     getStatisticsRequest.Token,
	}

	res, err := g.GetStatistics(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return res, nil
}

func RemoveStudent(g pb.GroupServiceClient, removeStudentRequest request.RemoveStudentRequest) error {
	req := &pb.RemoveStudentRequest{
		Token:   removeStudentRequest.Token,
		GroupId: removeStudentRequest.GroupId,
		UserId:  removeStudentRequest.UserId,
	}

	_, err := g.RemoveStudent(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while removing student from the group: %v", err)
	}

	return nil
}
