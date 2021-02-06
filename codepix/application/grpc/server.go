package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/diegoclair/imersao/codepix/domain/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//StartGrpcServer is responsibile to start a gRPC server
func StartGrpcServer(db contract.DataManager, port int) {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to get listener: %v", err)
	}

	// pixGrpcService := getPixGrpcService(db)
	// pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	fmt.Println("gRPC server is listening on port: ", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Could not start gRPC server: ", err)
	}
}

// func getPixGrpcService(db contract.RepoManager) *service.PixGrpcService {

// 	pixRepository := PixKeyReposirotyDB{
// 		DB: db,
// 	}
// 	accountRepository := AccountReposirotyDB{
// 		DB: db,
// 	}
// 	pixUseCase := usecase.PixUseCase{
// 		PixRepository:     pixRepository,
// 		AccountRepository: accountRepository,
// 	}

// 	return service.NewPixGrpcService(pixUseCase)
// }
