package service

import (
	"context"
	"log"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/imersao-fullstack-fullcycle/application/grpc/pb"
	"github.com/diegoclair/imersao-fullstack-fullcycle/application/usecase"
)

type PixGrpcService struct {
	pixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func NewGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		pixUseCase: usecase,
	}
}

func (s *PixGrpcService) AddPixKey(ctx context.Context, req *pb.AddPixKeyRequest) (*pb.AddPixKeyResponse, error) {

	key, err := s.pixUseCase.RegisterKey(req.Key, req.Kind, req.AccountID)
	if err != nil {
		return nil, err
	}

	return &pb.AddPixKeyResponse{
		Id:     key.ID,
		Status: "created",
	}, nil
}

func (s *PixGrpcService) FindPixKeyByID(ctx context.Context, req *pb.FindPixKeyByIDRequest) (*pb.FindPixKeyByIDResponse, error) {
	mapper := mapper.New()
	pix, err := s.pixUseCase.FindKeyByID(req.Key, req.Kind)
	if err != nil {
		return nil, err
	}

	response := &pb.FindPixKeyByIDResponse{}
	err = mapper.From(pix).To(&response)
	if err != nil {
		log.Println("Error to mapper response: ", err)
		return nil, err
	}

	response.Account.CreatedAt = pix.Account.CreatedAt.String()
	response.CreatedAt = pix.CreatedAt.String()

	return response, err
}
