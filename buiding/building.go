package buiding

import (
	"elevator/elevator"
)

type Building struct {
	Elevators *elevator.Elevator
	maxFloors int
}

func New(e *elevator.Elevator) *Building{
	var b Building

	b.Elevators = e
	return &b
}