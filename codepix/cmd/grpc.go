/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/diegoclair/imersao/codepix/application/grpc"
	"github.com/diegoclair/imersao/codepix/domain/service"
	"github.com/diegoclair/imersao/codepix/infrastructure/data"
	"github.com/spf13/cobra"
)

var port int

const defaultGRPCPort int = 50051

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "To start grpc server on default port 50051",
	Run: func(cmd *cobra.Command, args []string) {
		startGrpcServer()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.
	grpcCmd.Flags().IntVarP(&port, "port", "p", defaultGRPCPort, "gRPC server port")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startGrpcServer() {
	data, err := data.Connect()
	if err != nil {
		log.Fatalf("Error to connect data repositories: %v", err)
	}

	svc := service.New(data)

	grpc.StartGrpcServer(svc, port)
}
