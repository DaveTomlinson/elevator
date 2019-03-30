package simulation

import (
	"elevator/agent"
	"elevator/elevator"
	"fmt"
	"math/rand"
	"time"
)

const (
	maxSteps  = 1000
	numPeople = 5
)

type ElevatorSimulation struct {
	elev     *elevator.Elevator
	people   []Agent
	maxFloor int
	stepNum  int64
}

type Agent interface {
	elevator.Agent
	IsOnDesiredFloor() bool
	PrintStatus()
}

func New(e *elevator.Elevator) *ElevatorSimulation {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var es ElevatorSimulation
	es.elev = e
	es.maxFloor = e.GetMaxFloor()
	for i := 0; i < numPeople; i++ {
		currentFloor := myRand.Intn(es.maxFloor + 1)
		targetFloor := myRand.Intn(es.maxFloor + 1)
		for targetFloor == currentFloor {
			targetFloor = myRand.Intn(es.maxFloor)
		}
		es.people = append(es.people, agent.New(currentFloor, targetFloor))
	}
	return &es
}

func (es *ElevatorSimulation) SimulationStep() {
	es.printStatus()

	es.performAgentActions()
	es.elev.Move()
	es.stepNum++
}

func (es *ElevatorSimulation) Done() bool {
	if es.stepNum >= maxSteps || es.isEveryoneOnDesiredFloor() {
		return true
	}
	return false
}

func (es ElevatorSimulation) isEveryoneOnDesiredFloor() bool {
	good := true
	for _, person := range es.people {
		if !person.IsOnDesiredFloor() {
			good = false
		}
	}
	return good
}

func (es *ElevatorSimulation) performAgentActions() {
	for _, person := range es.people {
		es.elev.Board(person)
		if !es.elev.IsInElevator(person) {
			es.elev.CallElevator(person)

		} else {
			es.elev.Exit(person)
		}
	}
}

func (es ElevatorSimulation) printStatus() {
	msg := fmt.Sprintf("Runing simulation step %v", es.stepNum)
	fmt.Println(msg)

	es.elev.PrintStatus()
	es.printPersonStatus()
}

func (es ElevatorSimulation) printPersonStatus() {
	for _, person := range es.people {
		person.PrintStatus()
	}
}
