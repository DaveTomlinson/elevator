package simulation

import (
	"elevator/agent"
	"elevator/elevator"
	"fmt"
	"math/rand"
	"time"
)

const maxSteps = 1000

type ElevatorSimulation struct {
	elev     *elevator.Elevator
	person   *agent.Agent
	maxFloor int
	stepNum  int64
}

func New(e *elevator.Elevator) *ElevatorSimulation {
	myrand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var es ElevatorSimulation
	es.elev = e
	es.maxFloor = e.GetMaxFloor()
	currentFloor := myrand.Intn(es.maxFloor + 1)
	targetFloor := myrand.Intn(es.maxFloor + 1)
	for targetFloor == currentFloor {
		targetFloor = myrand.Intn(es.maxFloor)
	}
	es.person = agent.New(currentFloor, targetFloor)
	return &es
}

func (es *ElevatorSimulation) SimulationStep() {
	es.printStatus()

	es.elev.Board(es.person)
	if !es.elev.IsInElevator(es.person) {
		es.elev.CallElevator(es.person)

	} else{
		es.elev.Exit(es.person)
	}
	es.elev.Move()
	es.stepNum++
}

func (es *ElevatorSimulation) Done() bool {
	if es.stepNum >= maxSteps || es.person.IsOnDesiredFloor() {
		return true
	}
	return false
}

func (es ElevatorSimulation) printStatus() {
	msg := fmt.Sprintf("Runing simulation step %v", es.stepNum)
	fmt.Println(msg)

	es.elev.PrintStatus()
	es.person.PrintStatus()
}
