package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context) []Stock {

	var Stockslist []Stock
	read := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	for {
		var val Stock
		out, err := read.ReadMessage(ctx)

		if err != nil {
			panic("Receiving failed")
		}
		json.Unmarshal(out.Value, &val)
		if val.Name == "" {
			break
		}
		Stockslist = append(Stockslist, val)

	}
	return Stockslist
}
