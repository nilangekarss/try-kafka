package trykafka

/*
import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"time"
)

func CreateTopic(){
	topic := "my-topic"
	partition := 0

	conn, err := kafkago.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafkago.Message{Value: []byte("one!")},
		kafkago.Message{Value: []byte("two!")},
		kafkago.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}


	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
*/