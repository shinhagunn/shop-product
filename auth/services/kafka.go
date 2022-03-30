package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func EmailProducer(topic string, message interface{}) {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("kafka:9092"),
		Topic: topic,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("error: ", err)
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: messageBytes,
	})

	if err != nil {
		log.Println("cannot write a message: ", err)
	}
}
