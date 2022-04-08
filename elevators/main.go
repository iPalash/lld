package main

import (
	"elevators/elevator"
	"fmt"
)

func main() {
	fmt.Println("Implementing an elevator")
	var e elevator.Elevator
	e = elevator.NewElevator()
}
