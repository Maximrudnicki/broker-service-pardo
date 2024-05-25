package service

import (
	"errors"
	"log"

	"broker-service/cmd/data/request"
	"broker-service/cmd/data/response"
	u "broker-service/cmd/utils"
	pb "broker-service/proto"

	"github.com/go-playground/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GroupServiceImpl struct {
	Validate *validator.Validate
}

// AddStudent implements GroupService.
func (*GroupServiceImpl) AddStudent(asr request.AddStudentRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	err = u.AddStudent(c, asr)
	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}

// AddWordToUser implements GroupService.
func (*GroupServiceImpl) AddWordToUser(awur request.AddWordToUserRequest) (response.AddWordToUserResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	var addWordToUserResponse response.AddWordToUserResponse

	res, err := u.AddWordToUser(c, awur)
	if err != nil {
		return addWordToUserResponse, errors.New("cannot add word")
	}

	jsonResp := response.AddWordToUserResponse{
		WordId: res.WordId,
	}

	addWordToUserResponse = jsonResp

	return addWordToUserResponse, nil
}

// CreateGroup implements GroupService
func (g *GroupServiceImpl) CreateGroup(cgr request.CreateGroupRequest) error {
	// connect to group service as a client
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	err = u.CreateGroup(c, cgr)
	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}

// DeleteGroup implements GroupService.
func (*GroupServiceImpl) DeleteGroup(dgr request.DeleteGroupRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	err = u.DeleteGroup(c, dgr)
	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}

// FindGroup implements GroupService.
func (*GroupServiceImpl) FindGroup(fgr request.FindGroupRequest) (response.GroupResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	var groupResponse response.GroupResponse

	gr, err := u.FindGroup(c, fgr)
	if err != nil {
		return groupResponse, errors.New("cannot get groups")
	}

	jsonResp := response.GroupResponse{
		GroupId:  gr.GroupId,
		Title:    gr.Title,
		Students: gr.Students,
	}

	groupResponse = jsonResp

	return groupResponse, nil
}

// FindGroup implements GroupService.
func (*GroupServiceImpl) FindStudent(fsr request.FindStudentRequest) (response.StudentResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	var studentResponse response.StudentResponse

	sr, err := u.FindStudent(c, fsr)
	if err != nil {
		return studentResponse, errors.New("cannot find student")
	}

	jsonResp := response.StudentResponse{
		Email: sr.Email,
		Username: sr.Username,
	}

	studentResponse = jsonResp

	return studentResponse, nil
}

// FindGroupsStudent implements GroupService.
func (*GroupServiceImpl) FindGroupsStudent(fgsr request.FindGroupsStudentRequest) ([]response.GroupResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	ggr, err := u.FindGroupsStudent(c, fgsr)
	if err != nil {
		return nil, errors.New("cannot get groups")
	}

	// ggr - Group gRPC Response, groupResponse - JSON format
	var groupResponse []response.GroupResponse
	for _, grpcResp := range ggr {
		jsonResp := response.GroupResponse{
			GroupId:  grpcResp.GroupId,
			Title:    grpcResp.Title,
			Students: grpcResp.Students,
		}
		groupResponse = append(groupResponse, jsonResp)
	}

	return groupResponse, nil
}

// FindGroupsTeacher implements GroupService.
func (*GroupServiceImpl) FindGroupsTeacher(fgtr request.FindGroupsTeacherRequest) ([]response.GroupResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	ggr, err := u.FindGroupsTeacher(c, fgtr)
	if err != nil {
		return nil, errors.New("cannot get groups")
	}

	// ggr - Group gRPC Response, groupResponse - JSON format
	var groupResponse []response.GroupResponse
	for _, grpcResp := range ggr {
		jsonResp := response.GroupResponse{
			GroupId:  grpcResp.GroupId,
			Title:    grpcResp.Title,
			Students: grpcResp.Students,
		}
		groupResponse = append(groupResponse, jsonResp)
	}

	return groupResponse, nil
}

// RemoveStudent implements GroupService.
func (*GroupServiceImpl) RemoveStudent(rsr request.RemoveStudentRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupServiceClient(conn)

	err = u.RemoveStudent(c, rsr)
	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}

func NewGroupServiceImpl(validate *validator.Validate) GroupService {
	return &GroupServiceImpl{
		Validate: validate,
	}
}
