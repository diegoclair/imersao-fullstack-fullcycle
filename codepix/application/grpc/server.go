package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/diegoclair/imersao-fullstack-fullcycle/application/grpc/pb"
	"github.com/diegoclair/imersao-fullstack-fullcycle/application/grpc/service"
	"github.com/diegoclair/imersao-fullstack-fullcycle/application/usecase"
	"github.com/diegoclair/imersao-fullstack-fullcycle/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//StartGrpcServer is responsibile to start a gRPC server
func StartGrpcServer(db *gorm.DB, port int) {

	address := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to get listener: %v", err)
	}

	pixRepository := repository.PixKeyReposirotyDB{
		DB: db,
	}
	accountRepository := repository.AccountReposirotyDB{
		DB: db,
	}
	pixUseCase := usecase.PixUseCase{
		PixRepository:     pixRepository,
		AccountRepository: accountRepository,
	}
	pixGrpcService := service.NewGrpcService(pixUseCase)

	s := grpc.NewServer()
	pb.RegisterPixServiceServer(s, pixGrpcService)

	reflection.Register(s)

	fmt.Println("gRPC server is listening on port: ", port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatal("Could not start gRPC server: ", err)
	}
}
