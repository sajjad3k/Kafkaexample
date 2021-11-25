package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/piquette/finance-go/quote"
	"github.com/segmentio/kafka-go"
)

const (
	topic         = "stocks"
	brokerAddress = "localhost:9092"
)

func getdata(sym string) Stock {
	//var symb string
	var stck Stock
	out, err := quote.Get(sym)
	if err != nil {
		log.Fatal("error")
	}
	stck.Name = out.ShortName
	stck.Currprice = out.Ask
	stck.Fiftytwoweekshigh = out.FiftyTwoWeekHigh
	stck.Fiftytwoweekslow = out.FiftyTwoWeekLow
	stck.Exchangeid = out.ExchangeID
	return stck
}

func Produce(ctx context.Context, lst []string, wg *sync.WaitGroup) {

	write := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	for _, val := range lst {
		stockdata := getdata(val)
		req := new(bytes.Buffer)
		json.NewEncoder(req).Encode(stockdata)
		err := write.WriteMessages(ctx, kafka.Message{
			Key:   []byte(val),
			Value: req.Bytes(),
		})
		if err != nil {
			panic("failed")
		}

	}
	defer wg.Done()
}
