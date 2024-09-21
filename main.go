package main

import (
	"github.com/Netflix/go-env"
	"github.com/the-great-checkout/transactions-cdc/internal/client"
	"github.com/the-great-checkout/transactions-cdc/internal/service"
)

type Environment struct {
	Kafka struct {
		Topic   string `env:"KAFKA_TOPIC,default=transactions"`
		Address string `env:"KAFKA_ADDRESS,default=localhost:9092"`
	}

	Client struct {
		TransactionLog struct {
			Url string `env:"TRANSACTION_LOG_URL,default=http://localhost:8082"`
		}
	}
}

func main() {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		panic(err)
	}

	transactionLogClient := client.NewTransactionLogClient(environment.Client.TransactionLog.Url)
	transactionLogService := service.NewCDCService(transactionLogClient)

	notificationService := service.NewNotificationService(environment.Kafka.Topic, environment.Kafka.Address)
	notificationService.Consume(transactionLogService.ChangeDataCapture)
}
