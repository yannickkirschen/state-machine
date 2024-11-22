package fsm

import (
	"testing"
)

const (
	Open  = "open"
	Close = "close"

	OpenDoor  = "open-door"
	CloseDoor = "close-door"
)

type ComplexDoorState struct {
	Id       string
	OpenedBy string
}

func getMachine() *Machine[string] {
	machine := NewMachine(Close)
	machine.SetTransition(Open, CloseDoor, Close)
	machine.SetTransition(Close, OpenDoor, Open)
	return machine
}

func getComplexMachine() *Machine[*ComplexDoorState] {
	machine := NewMachine(&ComplexDoorState{Id: "close"})
	machine.SetTransition(&ComplexDoorState{Id: "open"}, "close-door", &ComplexDoorState{Id: "close"})
	machine.SetTransition(&ComplexDoorState{Id: "close"}, "open-door", &ComplexDoorState{Id: "open"})

	machine.SetEnterAction(func(last, new *ComplexDoorState) error {
		new.OpenedBy = "Peter"
		return nil
	})

	machine.SetExitAction(func(old, next *ComplexDoorState) error {
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

	if _, _, err := machine.Transition(OpenDoor); err != nil {
		t.Error(err)
	}

	if _, _, err := machine.Transition(OpenDoor); err == nil {
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

func TestComplexTransition(t *testing.T) {
	machine := getComplexMachine()

	var last *ComplexDoorState
	var current *ComplexDoorState
	if l, c, err := machine.Transition(OpenDoor); err != nil {
		t.Error(err)
	} else {
		last = *l
		current = *c
	}

	if _, _, err := machine.Transition(OpenDoor); err == nil {
		t.Error("expecting door to be opened")
	}

	state := machine.State()
	if state.OpenedBy != "Peter" {
		t.Errorf("expecting state %s to be 'Peter', not %s", state.Id, state.OpenedBy)
	}

	if state.Id != current.Id {
		t.Errorf("States are not equal (%s vs %s)", state.Id, current.Id)
	}

	last.Id = "this should have no effect"
	if state.Id == last.Id {
		t.Errorf("State and last state are equal")
	}
}
