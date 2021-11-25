package main

import (
	"context"
	"fmt"
	"sync"
)

type Stock struct {
	Name              string
	Currprice         float64
	Fiftytwoweekshigh float64
	Fiftytwoweekslow  float64
	Exchangeid        string
}

var names = []string{"AAPL", "AMZN", "NVDA", "TSLA", "MSFT", "GME", "AMC", "FB", "NFLX", "MRNA"}

func main() {

	var wtg sync.WaitGroup
	ctx := context.Background()
	wtg.Add(1)
	go Produce(ctx, names, &wtg)
	wtg.Wait()
	data := Consume(ctx)
	fmt.Println(data)

}
