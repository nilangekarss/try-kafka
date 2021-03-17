
package trykafka


/*
	import (
		"context"
		"fmt"
		kafkago "github.com/segmentio/kafka-go"
		"time"
	)
func newKafkaWriter(kafkaURL, topic string) *kafkago.Writer {
	return &kafkago.Writer{
		Addr:         kafkago.TCP(kafkaURL),
		Topic:        topic,
		Balancer:     &kafkago.LeastBytes{},
		MaxAttempts:  0,
		BatchSize:    0,
		BatchBytes:   0,
		BatchTimeout: 0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		RequiredAcks: 0,
		Async:        false,
		Completion:   nil,
		Compression:  0,
		Logger:       nil,
		ErrorLogger:  nil,
		Transport:    nil,
	}
}


func CreateProducer(){
	kafkaURL := "localhost:9092"
	topic := "TopicA"
	kafkago.CreatePartitionsRequest{
		Addr:         nil,
		Topics:       nil,
		ValidateOnly: false,
	}
	writer := newKafkaWriter(kafkaURL, topic)


}

*/