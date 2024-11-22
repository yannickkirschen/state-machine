package fsm

import (
	"fmt"
	"testing"
)

const (
	Open  = "open"
	Close = "close"

	OpenDoor  = "open-door"
	CloseDoor = "close-door"
)

func getMachine() *Machine {
	machine := NewMachine(Close)
	machine.SetTransition(Open, CloseDoor, Close)
	machine.SetTransition(Close, OpenDoor, Open)

	machine.SetEnterAction(func(last, new State) error {
		fmt.Printf("Enter '%s' coming from '%s'\n", new, last)
		return nil
	})

	machine.SetExitAction(func(current, next State) error {
		fmt.Printf("Leaving '%s' going to '%s'\n", current, next)
		return nil
	})

	return machine
}

func TestState(t *testing.T) {
	machine := getMachine()

	if machine.State() != Close {
		t.Error("expecting state to be close")
	}

	machine.Transition(OpenDoor)

	if machine.State() != Open {
		t.Error("expecting state to be open")
	}
}

func TestTransition(t *testing.T) {
	machine := getMachine()

	if err := machine.Transition(OpenDoor); err != nil {
		t.Error(err)
	}

	if err := machine.Transition(OpenDoor); err == nil {
		t.Error("expecting door to be opened")
	}
}

func TestCanTransition(t *testing.T) {
	machine := getMachine()

	if !machine.CanTransition(OpenDoor) {
		t.Error("expecting machine.CanTransition(\"open-door\") to return true")
	}

	if machine.CanTransition(CloseDoor) {
		t.Error("expecting machine.CanTransition(\"close-door\") to return false")
	}
}
