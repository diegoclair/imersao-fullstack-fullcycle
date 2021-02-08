package server

import (
	"context"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/imersao/codepix/application/factory"
	"github.com/diegoclair/imersao/codepix/application/grpc/pb"
	"github.com/diegoclair/imersao/codepix/domain/contract"
)

//PixServer holds pix server functions
type PixServer struct {
	pixService                       contract.PixService
	mapper                           mapper.Mapper
	pb.UnimplementedPixServiceServer //its necessary when we are using grpc
}

//NewPixServer to handle requests
func NewPixServer(factory *factory.Services) *PixServer {
	return &PixServer{
		pixService: factory.PixService,
		mapper:     factory.Mapper,
	}
}

//AddPixKey handle an add pix key request
func (s *PixServer) AddPixKey(ctx context.Context, req *pb.AddPixKeyRequest) (*pb.AddPixKeyResponse, error) {

	key, err := s.pixService.RegisterKey(req.Key, req.KeyType, req.AccountID)
	if err != nil {
		return nil, err
	}

	return &pb.AddPixKeyResponse{
		Id:     key.ID,
		Status: "created",
	}, nil
}

//FindPixKeyByID handle a find pix key request
func (s *PixServer) FindPixKeyByID(ctx context.Context, req *pb.FindPixKeyByIDRequest) (*pb.FindPixKeyByIDResponse, error) {

	pix, err := s.pixService.FindKeyByID(req.Key, req.KeyType)
	if err != nil {
		return nil, err
	}

	response := &pb.FindPixKeyByIDResponse{
		Id:      pix.ID,
		KeyType: pix.KeyType,
		Key:     pix.Key,
		Account: &pb.Account{
			AccountID:     pix.Account.ID,
			AccountNumber: pix.Account.Number,
			BankID:        pix.Account.Bank.ID,
			BankName:      pix.Account.Bank.Name,
			OwnerName:     pix.Account.OwnerName,
			CreatedAt:     pix.Account.CreatedAt.String(),
		},
		CreatedAt: pix.CreatedAt.String(),
	}

	return response, err
}
