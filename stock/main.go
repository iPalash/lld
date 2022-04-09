package main

import (
	"context"
	"stock/engine"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ob := engine.NewOrderBook(ctx)
	ob.Buy(engine.NewOrder(10, 10, engine.BUY))

	// ob.Sell(engine.NewOrder(11, 10, engine.SELL))
	time.Sleep(time.Second)
	ob.Buy(engine.NewOrder(11, 5, engine.BUY))

	ob.Sell(engine.NewOrder(10, 20, engine.SELL))
	time.Sleep(time.Second * 2)

	ob.Buy(engine.NewOrder(6, 5, engine.BUY))
	// cancel()
	time.Sleep(time.Second * 1)
}
