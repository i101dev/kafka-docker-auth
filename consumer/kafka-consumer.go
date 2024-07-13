package main

import (
	"fmt"
	"kafka-docker-auth/util"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func main() {

	// -------------------------------------------------------------------------------
	// Connect to broker and create new consumer -------------------------------------

	tlsConfig := util.LoadTLSConfig()
	brokerConsumer, err := sarama.NewConsumer([]string{"localhost:9092"}, tlsConfig)
	if err != nil {
		fmt.Printf("Error creating [brokerConsumer]: %v\n", err)
		return
	}
	defer func() {
		if err := brokerConsumer.Close(); err != nil {
			fmt.Printf("Error closing [brokerConsumer]: %v\n", err)
		}
	}()

	// -------------------------------------------------------------------------------
	// Subscribe to topics -----------------------------------------------------------

	topic := "your_topic"
	topicConsumer, err := brokerConsumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("Error creating [topicConsumer]: %v\n", err)
		return
	}
	defer func() {
		if err := topicConsumer.Close(); err != nil {
			fmt.Printf("Error closing [topicConsumer]: %v\n", err)
		}
	}()

	// -------------------------------------------------------------------------------
	// Handle messages ---------------------------------------------------------------

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	log.Println("Consumer connection success")

	for {
		select {
		case msg := <-topicConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
		case <-signals:
			fmt.Println("Interrupted, shutting down...")
			return
		}
	}
}
