package service

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type NotificationService struct {
	topic, address string
}

func NewNotificationService(topic, address string) *NotificationService {
	return &NotificationService{topic: topic, address: address}
}

func (s *NotificationService) Consume(handler func([]byte) error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{s.address},
		Topic:    s.topic,
		GroupID:  s.topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error on read message: %v", err)
			continue
		}

		err = handler(message.Value)
		if err != nil {
			log.Print(err.Error())
		}

		// Simula tempo de processamento
		time.Sleep(time.Second)
	}
}
