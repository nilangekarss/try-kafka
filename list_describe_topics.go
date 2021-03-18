package trykafka

import (
	"context"
	"fmt"
	confkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"time"
)

func getAllTopics(admin *confkafka.AdminClient) []string{
	mdata, terr := admin.GetMetadata(nil,true, 100)
	if terr != nil {
		fmt.Printf("Failed to get metadata for all topics: %s\n", terr)
		os.Exit(1)
	}
	topicsdetail := mdata.Topics
	//mymap := make(map[int]string)
	keys := make([]string, 0, len(topicsdetail))
	for k := range topicsdetail {
		keys = append(keys, k)
		fmt.Println("name of topic is : %s", k)
	}
return keys
}

func DescribeTopic(ctx context.Context){
	a, err := confkafka.NewAdminClient(&confkafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	//here specify the resource type "any/topic/group/broker/else(becomes unknown)"
	resourceType, err := confkafka.ResourceTypeFromString("topic")
	resourceName := "users3"

	topics := getAllTopics(a)
	fmt.Println("Topics associated with broker are : %v", topics)
	
	metadata, merr := a.GetMetadata(&resourceName, false, 100)

	if merr != nil {
		fmt.Printf("Failed to get metadata for topic: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("metadata for topics %#v", metadata.Topics)
	for k,v := range metadata.Topics{
		fmt.Println("k and v is %s %#v", k,v)
		for index, partitions := range v.Partitions{
			fmt.Println("index and list of partitions are %s %#v", index, partitions)
		}
	}

	ctxt, cancel := context.WithCancel(ctx)
	defer cancel()
	dur, _ := time.ParseDuration("20s")
	results, err := a.DescribeConfigs(ctxt,
		[]confkafka.ConfigResource{{Type: resourceType, Name: resourceName}},
		confkafka.SetAdminRequestTimeout(dur))
	if err != nil {
		fmt.Printf("Failed to DescribeConfigs(%s, %s): %s\n",
			resourceType, resourceName, err)
		os.Exit(1)
	}
	// Print results
	for _, result := range results {
		fmt.Printf("%s %s: %s:\n", result.Type, result.Name, result.Error)
		for _, entry := range result.Config {
			// Truncate the value to 60 chars, if needed, for nicer formatting.
			fmt.Printf("%60s = %-60.60s   %-20s Read-only:%v Sensitive:%v\n",
				entry.Name, entry.Value, entry.Source,
				entry.IsReadOnly, entry.IsSensitive)
		}
	}
}
