package utils

import (
	"broker-service/cmd/data/request"
	pb "broker-service/proto"
	"context"
	"fmt"
	"io"
)

func AddStudent(g pb.GroupServiceClient, asr request.AddStudentRequest) error {
	req := &pb.AddStudentRequest{
		Token: asr.Token,
		GroupId: asr.GroupId,
	}

	_, err := g.AddStudent(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while adding student to the group: %v", err)
	}

	return nil
}

func AddWordToUser(g pb.GroupServiceClient, awur request.AddWordToUserRequest) (*pb.AddWordToUserResponse, error) {
	req := &pb.AddWordToUserRequest{
		Word: awur.Word,
		Definition: awur.Definition,
		GroupId: awur.GroupId,
		UserId: awur.UserId,
		Token: awur.Token,
	}

	res, err := g.AddWordToUser(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened while adding student to the group: %v", err)
	}

	return res, nil
}

func CreateGroup(g pb.GroupServiceClient, cgr request.CreateGroupRequest) error {
	req := &pb.CreateGroupRequest{
		Title:        cgr.Title,
		Token: cgr.Token,
	}

	_, err := g.CreateGroup(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while creating group: %v", err)
	}

	return nil
}

func DeleteGroup(g pb.GroupServiceClient, dgr request.DeleteGroupRequest) error {
	req := &pb.DeleteGroupRequest{
		Token: dgr.Token,
		GroupId: dgr.GroupId,
	}

	_, err := g.DeleteGroup(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened: %v", err)
	}

	return nil
}

func FindGroup(g pb.GroupServiceClient, dgr request.FindGroupRequest) (*pb.GroupResponse, error) {
	req := &pb.FindGroupRequest{
		Token: dgr.Token,
		GroupId: dgr.GroupId,
	}

	group, err := g.FindGroup(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return group, nil
}

func FindStudent(g pb.GroupServiceClient, fsr request.FindStudentRequest) (*pb.StudentResponse, error) {
	req := &pb.FindStudentRequest{
		Token: fsr.Token,
		StudentId: fsr.StudentId,
		GroupId: fsr.GroupId,
	}

	student, err := g.FindStudent(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return student, nil
}

func FindTeacher(g pb.GroupServiceClient, ftr request.FindTeacherRequest) (*pb.TeacherResponse, error) {
	req := &pb.FindTeacherRequest{
		Token: ftr.Token,
		GroupId: ftr.GroupId,
	}

	teacher, err := g.FindTeacher(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return teacher, nil
}

func FindGroupsTeacher(g pb.GroupServiceClient, fgtr request.FindGroupsTeacherRequest) ([]*pb.GroupResponse, error) {
	req := &pb.FindGroupsTeacherRequest{
		Token: fgtr.Token,
	}

	var grs []*pb.GroupResponse // group response slice

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

		grs = append(grs, res)
	}

	return grs, nil
}

func FindGroupsStudent(g pb.GroupServiceClient, fgsr request.FindGroupsStudentRequest) ([]*pb.GroupResponse, error) {
	req := &pb.FindGroupsStudentRequest{
		Token: fgsr.Token,
	}

	var grs []*pb.GroupResponse // group response slice

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

		grs = append(grs, res)
	}

	return grs, nil
}

func GetStatistics(g pb.GroupServiceClient, gsr request.GetStatisticsRequest) (*pb.StatisticsResponse, error) {
	req := &pb.GetStatisticsRequest{
		StudentId: gsr.StudentId,
		GroupId: gsr.GroupId,
		Token: gsr.Token,
	}

	res, err := g.GetStatistics(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error happened: %v", err)
	}

	return res, nil
}

func RemoveStudent(g pb.GroupServiceClient, rsr request.RemoveStudentRequest) error {
	req := &pb.RemoveStudentRequest{
		Token: rsr.Token,
		GroupId: rsr.GroupId,
		UserId: rsr.UserId,
	}

	_, err := g.RemoveStudent(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error happened while removing student from the group: %v", err)
	}

	return nil
}
