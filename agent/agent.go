package agent

import (
	"elevator/elevatorDirections"
	"fmt"
	"github.com/google/uuid"
)

type Agent struct {
	id           uuid.UUID
	desiredFloor int
	currentFloor int
}

func New(currentFloor int, desiredFloor int) *Agent {
	var agent Agent

	agent.id = uuid.New()
	agent.desiredFloor = desiredFloor
	agent.currentFloor = currentFloor

	return &agent
}

func (a *Agent) SetCurrentFloor(floor int) {
	a.currentFloor = floor
}

func (a Agent) GetFullId() string {
	return a.id.String()
}

func (a Agent) GetId() string {
	return a.id.String()[:5]
}

func (a Agent) GetCallDirection() elevatorDirections.Direction {
	if a.GetFloorDifference() < 0 {
		return elevatorDirections.DOWN
	}
	if a.GetFloorDifference() > 0 {
		return elevatorDirections.UP
	}
	return elevatorDirections.STAY
}

func (a Agent) IsOnDesiredFloor() bool {
	if a.GetFloorDifference() == 0 {
		return true
	}
	return false
}

func (a Agent) GetFloorDifference() int {
	return a.desiredFloor - a.currentFloor
}

func (a Agent) GetDesiredFloor() int {
	return a.desiredFloor
}

func (a Agent) GetCurrentFloor() int {
	return a.currentFloor
}

func (a Agent) PrintStatus() {
	var msg string
	if a.currentFloor != a.desiredFloor {
		msg = fmt.Sprintf("   Passenger %v is on floor %v and has %v as the desired floor", a.GetId(), a.currentFloor, a.desiredFloor)
	} else {
		msg = fmt.Sprintf("   Passenger %v is on their desired floor", a.GetId())
	}
	fmt.Println(msg)
	//msg = fmt.Sprintf("   Passenger %v has %v as the desired floor", a.GetId(), a.desiredFloor)
	//fmt.Println(msg)
}
