package main

import (
	"context"
	"github.com/nilangekarss/trykafka"
)

const (
	topic          = "message-log"
	broker1Address = "localhost:9093"
	broker2Address = "localhost:9094"
	broker3Address = "localhost:9095"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go trykafka.Produce(ctx)
	trykafka.Consume(ctx)
}

