package engine

import (
	"container/heap"
	"context"
	"fmt"
	"time"
)

type Orderbook interface {
	Buy(Order)
	Sell(Order)
	Cancel(Order)
}

type Trade struct {
	BuyOrderID  int
	SellOrderID int
	Volume      int
	Price       int
}

func (t *Trade) String() string {
	return fmt.Sprintf("Matched BUY:%d with SELL:%d for %d qty with %d price", t.BuyOrderID, t.SellOrderID, t.Volume, t.Price)
}

type OrderBookImpl struct {
	buy     (chan Order)
	sell    (chan Order)
	cancel  (chan Order)
	buys    MaxHeap
	sells   MinHeap
	history []*Trade
}

func NewOrderBook(ctx context.Context) Orderbook {
	ob := &OrderBookImpl{
		buy:    make(chan Order),
		sell:   make(chan Order),
		cancel: make(chan Order),
		buys:   make(MaxHeap, 0),
		sells:  make(MinHeap, 0),
	}

	go func(ctx context.Context) {
		ob.process(ctx)
	}(ctx)

	return ob
}

func (ob *OrderBookImpl) addTrade(buy, sell *Order) {
	var vol int
	if buy.Volume < sell.Volume {
		vol = buy.Volume
	} else {
		vol = sell.Volume
	}
	trade := &Trade{BuyOrderID: buy.ID, SellOrderID: sell.ID, Volume: vol, Price: buy.Price}
	ob.history = append(ob.history, trade)
	fmt.Println(trade)
	buy.Volume -= vol
	sell.Volume -= vol

}

func (ob *OrderBookImpl) _buy(o Order) {
	for len(ob.sells) > 0 && ob.sells[0].Price <= o.Price {
		ob.addTrade(&o, &ob.sells[0])

		if ob.sells[0].Volume == 0 {
			heap.Pop(&ob.sells)
		}

		if o.Volume == 0 {
			return
		}
	}
	heap.Push(&ob.buys, o)
}

func (ob *OrderBookImpl) Buy(o Order) {
	ob.buy <- o
}

func (ob *OrderBookImpl) _sell(o Order) {

	for len(ob.buys) > 0 && ob.buys[0].Price >= o.Price {
		ob.addTrade(&ob.buys[0], &o)

		if ob.buys[0].Volume == 0 {
			heap.Pop(&ob.buys)
		}

		if o.Volume == 0 {
			return
		}
	}

	heap.Push(&ob.sells, o)
}

func (ob *OrderBookImpl) Sell(o Order) {
	ob.sell <- o
}

func (ob *OrderBookImpl) _cancel(o Order) {

	//TODO: Cancel order
	if o.Type == BUY {
		for i, order := range ob.buys {
			if order == o {
				heap.Remove(&ob.buys, i)
				break
			}
		}
	} else if o.Type == SELL {
		for i, order := range ob.sells {
			if order == o {
				heap.Remove(&ob.buys, i)
				break
			}
		}
	}
}

func (ob *OrderBookImpl) Cancel(o Order) {
	ob.cancel <- o
}

func (ob *OrderBookImpl) String() string {
	return fmt.Sprintf("BUYS: %v \n SELLS: %v", ob.buys, ob.sells)

}

func (ob *OrderBookImpl) process(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	loop := true
	for loop {
		select {
		case o := <-ob.buy:
			ob._buy(o)
		case o := <-ob.sell:
			ob._sell(o)
		case o := <-ob.cancel:
			ob._cancel(o)
		case <-ticker.C:
			fmt.Println("Orderbook\n", ob)
		case <-ctx.Done():
			fmt.Println("Shutting down")
			loop = false
		}
	}
}
