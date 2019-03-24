package elevator

import (
	"elevator/agent"
	"elevator/elevatorDirections"
	"errors"
	"fmt"
	"math"
)

const Speed = 0.1

type Elevator struct {
	floorTarget   int
	currentHeight float64
	people        *agent.Agent
	maxFloor      int
	doorsOpen     bool
	direction     elevatorDirections.Direction
}

func New(maxFloor int) (*Elevator, error) {
	var e Elevator
	if maxFloor <= 0 {
		return nil, errors.New("maxFloor must be > 0")
	}
	e.maxFloor = maxFloor
	return &e, nil
}

func (e *Elevator) CallElevator(a *agent.Agent) {
	e.floorTarget = a.GetCurrentFloor()
}

func (e *Elevator) Board(a *agent.Agent) {
	if e.doorsOpen {
		e.people = a
		e.floorTarget = a.GetDesiredFloor()
		a.SetCurrentFloor(-1)
	}
}

func (e *Elevator) Exit(a *agent.Agent) {
	if e.IsInElevator(a) {
		if e.doorsOpen {
			if e.isAtAgentsTargetFloor(a) {
				a.SetCurrentFloor(a.GetDesiredFloor())
				fmt.Printf("  ** %v has left the elevator ** \n", a.GetId())
				e.people = nil
			}
		}
	}
}

func (e *Elevator) Move() {
	if e.isAtTargetFloor() {
		e.doorsOpen = true
		return
	}
	e.doorsOpen = false
	if e.currentHeight < float64(e.floorTarget) {
		e.goUp()
		return
	}

	if e.currentHeight > float64(e.floorTarget) {
		e.goDown()
		return
	}
}
func (e *Elevator) SetTargetFloor(floor int) {
	if floor <= e.maxFloor && floor >= 0 {
		e.floorTarget = floor
	}
}

func (e Elevator) GetMaxFloor() int {
	return e.maxFloor
}

func (e Elevator) IsInElevator(a *agent.Agent) bool {
	if a == e.people {
		return true
	}
	return false
}

func (e Elevator) GetHeight() float64 {
	return e.currentHeight
}

func (e Elevator) PrintStatus() {
	var msg string
	msg = fmt.Sprintf("   Elevator currently at height %v", e.GetHeight())
	fmt.Println(msg)
	msg = fmt.Sprintf("   Elevator moving to %v", e.floorTarget)
	fmt.Println(msg)
}

func (e *Elevator) goUp() {
	if e.currentHeight <= float64(e.maxFloor) {
		e.currentHeight += Speed

	}
}

func (e *Elevator) goDown() {
	if e.currentHeight >= 0 {
		e.currentHeight -= Speed
	}
}

func (e *Elevator) isAtTargetFloor() bool {
	return e.isAtFloor(e.floorTarget)
}

func (e *Elevator) isAtFloor(floor int) bool {
	if math.Abs(e.currentHeight-float64(floor)) < Speed {
		e.currentHeight = float64(floor)
		return true
	}
	return false
}

func (e *Elevator) isAtAgentsTargetFloor(a *agent.Agent) bool {
	return e.isAtFloor(a.GetDesiredFloor())
}
