package main

import (
	"elevator/elevator"
	"elevator/simulation"
)

type Simulation interface {
	SimulationStep()
	Done() bool
}

func runSimulation(s Simulation) {
	for !s.Done() {
		s.SimulationStep()
	}
}

func main() {
	e, err := elevator.New(10)
	if err != nil {
		panic(err)
	}
	sim := simulation.New(e)
	runSimulation(sim)
}
