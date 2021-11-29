package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func Consume(ctx context.Context) []Stock {

	var Stockslist []Stock
	
	mech := plain.Mechanism{
		Username: username,
		Password: password,
	}

	dial := kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mech,
	}
	
	
	read := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Dialer:  &dial,
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
