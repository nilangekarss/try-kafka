package trykafka


import (
	"context"
	"fmt"
	kafkago "github.com/segmentio/kafka-go"
)

func CreateTopic(ctx context.Context) {

//	conn1, err := kafkago.DialLeader(ctx, "tcp", "localhost:9092", "topic-swapnil", 3)
//	if err != nil {
//		panic(err.Error())
//	}
//	defer conn1.Close()

	conn, err := kafkago.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	fmt.Println("partitions are %#v", partitions)
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
		fmt.Println("printing p value %#v", p)
	}
	for k := range m {
		fmt.Println("Printing Topics: %v", k)
	}

}


/*
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
