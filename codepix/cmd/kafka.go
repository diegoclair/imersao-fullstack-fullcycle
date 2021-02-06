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
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegoclair/imersao/codepix/application/kafka"
	"github.com/diegoclair/imersao/codepix/infrastructure/data"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Starting consuming transactions using apache kafka",
	Run: func(cmd *cobra.Command, args []string) {
		startApacheKafkaServer()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startApacheKafkaServer() {

	fmt.Println("Producing a message")
	producer := kafka.NewKafkaProducer()

	deliveryChan := make(chan ckafka.Event)
	kafka.Publish("Olá Kafka", "teste", producer, deliveryChan)
	go kafka.DeliveryReport(deliveryChan)

	data, err := data.Connect()
	if err != nil {
		log.Fatalf("Error to connect data repositories: %v", err)
	}

	kafkaProcessor := kafka.NewKafkaProcessor(data, producer, deliveryChan)
	kafkaProcessor.Consume()
}
