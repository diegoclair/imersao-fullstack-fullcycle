package server

import (
	"context"
	"log"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/imersao/codepix/application/grpc/pb"
	"github.com/diegoclair/imersao/codepix/domain/contract"
)

//PixServer is a struct to interact with services
type PixServer struct {
	pixService                       contract.PixService
	mapper                           mapper.Mapper
	pb.UnimplementedPixServiceServer //its necessary when we are using grpc
}

//NewPixServer to handle requests
func NewPixServer(pixService contract.PixService, mapper mapper.Mapper) *PixServer {
	return &PixServer{
		pixService: pixService,
		mapper:     mapper,
	}
}

func (s *PixServer) AddPixKey(ctx context.Context, req *pb.AddPixKeyRequest) (*pb.AddPixKeyResponse, error) {

	key, err := s.pixService.RegisterKey(req.Key, req.Kind, req.AccountID)
	if err != nil {
		return nil, err
	}

	return &pb.AddPixKeyResponse{
		Id:     key.ID,
		Status: "created",
	}, nil
}

func (s *PixServer) FindPixKeyByID(ctx context.Context, req *pb.FindPixKeyByIDRequest) (*pb.FindPixKeyByIDResponse, error) {

	pix, err := s.pixService.FindKeyByID(req.Key, req.Kind)
	if err != nil {
		return nil, err
	}

	response := &pb.FindPixKeyByIDResponse{}
	err = s.mapper.From(pix).To(&response)
	if err != nil {
		log.Println("Error to mapper response: ", err)
		return nil, err
	}

	response.Account.CreatedAt = pix.Account.CreatedAt.String()
	response.CreatedAt = pix.CreatedAt.String()

	return response, err
}
