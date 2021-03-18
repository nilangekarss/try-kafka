package trykafka
import (
	"context"
	"fmt"
	"time"
	//      "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	//kafkago "github.com/segmentio/kafka-go"
)


func Consume(ctx context.Context) {



	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	//c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
	c.SubscribeTopics([]string{"users3"}, nil)
	cgmetadata, _ := c.GetConsumerGroupMetadata()
	fmt.Println("Consumer group metadata associated with consumer is %#v", cgmetadata)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		fmt.Println("Consumer sleeping for 5 seconds")
		time.Sleep(time.Second * 5)
	}

	c.Close()
}

/*
import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
		GroupID: "my-group",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
 */
