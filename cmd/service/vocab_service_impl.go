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

type VocabServiceImpl struct {
	Validate *validator.Validate
}

func NewVocabServiceImpl(validate *validator.Validate) VocabService {
	return &VocabServiceImpl{
		Validate: validate,
	}
}

// CreateWord implements VocabService
func (v *VocabServiceImpl) CreateWord(cwr request.CreateWordRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVocabServiceClient(conn)

	create_err := u.CreateWord(c, cwr)
	if create_err != nil {
		return create_err
	}

	return nil
}

// DeleteWord implements VocabService
func (v *VocabServiceImpl) DeleteWord(dwr request.DeleteWordRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVocabServiceClient(conn)

	delete_err := u.DeleteWord(c, dwr)
	if delete_err != nil {
		return delete_err
	}

	return nil
}

// DeleteWord implements VocabService
func (v *VocabServiceImpl) FindWord(fwr request.FindWordRequest) (response.VocabResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVocabServiceClient(conn)

	var findWordReponse response.VocabResponse

	word, fw_err := u.FindWord(c, fwr)
	if fw_err != nil {
		return findWordReponse, fw_err
	}

	jsonResp := response.VocabResponse{
		ID:              word.Id,
		Word:            word.Word,
		Definition:      word.Definition,
		CreatedAt:       word.CreatedAt,
		IsLearned:       word.IsLearned,
		Cards:           word.Cards,
		WordTranslation: word.WordTranslation,
		Constructor:     word.Constructor,
		WordAudio:       word.WordAudio,
	}

	findWordReponse = jsonResp

	return findWordReponse, nil
}

// GetWords implements VocabService
func (v *VocabServiceImpl) GetWords(vr request.VocabRequest) ([]response.VocabResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVocabServiceClient(conn)

	gr, err := u.GetWords(c, vr.Token)
	if err != nil {
		return nil, errors.New("cannot get words")
	}

	// gr - gRPC Response, vocabResponse - JSON format
	var vocabResponse []response.VocabResponse
	for _, grpcResp := range gr {
		jsonResp := response.VocabResponse{
			ID:              grpcResp.Id,
			Word:            grpcResp.Word,
			Definition:      grpcResp.Definition,
			CreatedAt:       grpcResp.CreatedAt,
			IsLearned:       grpcResp.IsLearned,
			Cards:           grpcResp.Cards,
			WordTranslation: grpcResp.WordTranslation,
			Constructor:     grpcResp.Constructor,
			WordAudio:       grpcResp.WordAudio,
		}
		vocabResponse = append(vocabResponse, jsonResp)
	}

	return vocabResponse, nil
}

// UpdateWord implements VocabService
func (v *VocabServiceImpl) UpdateWord(uwr request.UpdateWordRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVocabServiceClient(conn)

	update_err := u.UpdateWord(c, uwr)
	if update_err != nil {
		return update_err
	}

	return nil
}

// UpdateWordStatus implements VocabService
func (v *VocabServiceImpl) UpdateWordStatus(uwsr request.UpdateWordStatusRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVocabServiceClient(conn)

	update_err := u.UpdateWordStatus(c, uwsr)
	if update_err != nil {
		return update_err
	}

	return nil
}

// ManageTrainings implements VocabService.
func (*VocabServiceImpl) ManageTrainings(mtr request.ManageTrainingsRequest) error {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVocabServiceClient(conn)

	manage_trainings_err := u.ManageTrainings(c, mtr)
	if manage_trainings_err != nil {
		return manage_trainings_err
	}
	
	return nil
}
