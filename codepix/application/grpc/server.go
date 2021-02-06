package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/imersao/codepix/application/grpc/pb"
	"github.com/diegoclair/imersao/codepix/application/grpc/server"
	"github.com/diegoclair/imersao/codepix/domain/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//StartGrpcServer is responsibile to start a gRPC server
func StartGrpcServer(svc *service.Service, port int) {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to get listener: %v", err)
	}

	registerGRPCServices(grpcServer, svc)

	fmt.Println("gRPC server is listening on port: ", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Could not start gRPC server: ", err)
	}
}

func registerGRPCServices(grpcServer *grpc.Server, svc *service.Service) {
	mapper := mapper.New()
	svm := service.NewServiceManager()

	pixService := svm.PixService(svc)
	pixServer := server.NewPixServer(pixService, mapper)
	pb.RegisterPixServiceServer(grpcServer, pixServer)
}
