package main

import (
	"context"
	"strconv"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
	)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		res := client.ProduceSync(context.Background(), &kgo.Record{
			Topic: "test",
			Value: []byte("hello piska"),
			Key:   []byte(strconv.Itoa(int(i))),
		})

		if res.FirstErr() != nil {
			panic(res.FirstErr())
		}
	}
	defer client.Close()
}
