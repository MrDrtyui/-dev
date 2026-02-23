package main

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumerGroup("test-group"),
		kgo.ConsumeTopics("test"),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for {
		fetches := client.PollFetches(context.Background())
		if errs := fetches.Errors(); len(errs) > 0 {
			fmt.Printf("piska")
			continue
		}

		fetches.EachRecord(func(record *kgo.Record) {
			fmt.Printf("Topic=%s Partition=%d Offset=%d Key=%s Value=%s\n",
				record.Topic,
				record.Partition,
				record.Offset,
				record.Key,
				record.Value)
		})

	}
}
