package util

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
)

func LogFilesAtPath(path string) {
	files, _ := os.ReadDir(path)
	for _, file := range files {
		fmt.Println("*** >>> ", file.Name())
	}
}

func LoadCertificates() (tls.Certificate, *x509.CertPool) {

	cert, err := tls.LoadX509KeyPair("../cert/keystore/kafka.pem", "../cert/keystore/kafka.pem")
	if err != nil {
		log.Fatalf("Error loading TLS certificate: %v", err)
	}

	caCert, err := os.ReadFile("../cert/truststore/kafka-truststore.pem")
	if err != nil {
		log.Fatalf("Error loading CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return cert, caCertPool
}

func LoadTLSConfig() *sarama.Config {

	cert, caCertPool := LoadCertificates()

	config := sarama.NewConfig()

	config.ClientID = "kafka-go-consumer"
	config.Net.SASL.Enable = true

	config.Net.SASL.User = "controller_user"
	config.Net.SASL.Password = "bitnami"
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext

	config.Net.TLS.Enable = true
	config.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true, // Change to false in production and configure proper certificates
		ClientAuth:         tls.RequireAndVerifyClientCert,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
	}

	return config
}
