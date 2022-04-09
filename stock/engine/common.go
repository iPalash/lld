package engine

import (
	"fmt"
	"math/rand"
)

type Stock interface {
	Ticker() string
}

type User interface {
	PlaceBuyOrder(Stock, int, int) Order
	PlaceSellOrder(Stock, int, int) Order
}

type OrderStatus int

const (
	CREATED OrderStatus = iota
	PENDING
	EXECUTED
	CANCELLED
)

type OrderType string

const (
	BUY  OrderType = "Buy"
	SELL OrderType = "Sell"
)

type Order struct {
	ID     int
	Stock  Stock
	Price  int
	Volume int
	Status OrderStatus
	Type   OrderType
}

func NewOrder(price int, volume int, orderType OrderType) Order {
	return Order{ID: rand.Int(), Price: price, Volume: volume, Status: CREATED, Type: orderType}
}

func (o Order) String() string {
	return fmt.Sprintf("%s order for %d qty at %d price", o.Type, o.Volume, o.Price)
}
