package elevator

import (
	"fmt"
	"time"
)

type Direction int

const (
	Stationary Direction = iota
	Up
	Down
)

type Elevator interface {
	GetFloor() int
	GetDirection() Direction
	Stop()
	Press(floor int)
	Open()
	Close()
}

// type ElevatorInPanel interface {

// }

type ElevatorOutPanel interface {
	CallUp()
	CallDown()
}

type ElevatorInPanelImpl struct {
}

type State int

const (
	Closed State = iota
	Closing
	Open
	Opening
)

type ElevatorImpl struct {
	State     State
	Direction Direction
	Floor     Floor
}

func NewElevator() Elevator {
	return &ElevatorImpl{}
}

func (e *ElevatorImpl) GetFloor() int {
	return e.Floor.Number
}

func (e *ElevatorImpl) GetDirection() Direction {
	return e.Direction
}

func (e *ElevatorImpl) Stop() {
	t := time.NewTicker(1000 * time.Millisecond)
	t.C
	// Move to nearest direction floor and stop
}

func (e *ElevatorImpl) Press(floor int) {
	// Take request for the floor and add to queue and start moving if applicable
}

func (e *ElevatorImpl) Open() {
	// If door is closing, open"
	if e.State == Closing {
		fmt.Println("Opening doors")
		e.State = Opening
	}
}

func (e *ElevatorImpl) Close() {
	// If door is open, close
	if e.State == Open {
		fmt.Println("Closing doors")
		e.State = Closing
	}
}
