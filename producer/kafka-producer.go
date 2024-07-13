package main

import (
	"crypto/tls"
	"log"

	"github.com/IBM/sarama"
)

func main() {

	config := sarama.NewConfig()

	config.Net.SASL.Enable = true

	config.Net.SASL.User = "controller_user"
	config.Net.SASL.Password = "bitnami"
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext

	config.Net.TLS.Enable = true
	config.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true, // Change to false in production
	}

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// Create a producer --------------------------------------------------------------
	producer, err := sarama.NewSyncProducer([]string{"example.kafka.com:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing producer: %v", err)
		}
	}()

	// Send a message -----------------------------------------------------------------
	msg := &sarama.ProducerMessage{
		Topic: "your_topic",
		Key:   sarama.StringEncoder("key"),
		Value: sarama.StringEncoder("value 99"),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
