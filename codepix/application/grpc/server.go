package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/diegoclair/imersao/codepix/application/factory"
	"github.com/diegoclair/imersao/codepix/application/grpc/pb"
	"github.com/diegoclair/imersao/codepix/application/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//StartGrpcServer is responsibile to start a gRPC server
func StartGrpcServer(port int) {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to get listener: %v", err)
	}

	registerGRPCServices(grpcServer)

	log.Println("gRPC server is listening on port: ", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Could not start gRPC server: ", err)
	}
}

func registerGRPCServices(grpcServer *grpc.Server) {

	factory := factory.GetDomainServices()

	pixServer := server.NewPixServer(factory)
	pb.RegisterPixServiceServer(grpcServer, pixServer)
}
