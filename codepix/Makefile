#To make sure that this command will execute with success, execute it inside of your docker container, using docker exec command..
pb_generate:
	protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

run-grpc:
	go run main.go grpc

run-kafka:
	go run main.go kafka

run-all:
	go run main.go all