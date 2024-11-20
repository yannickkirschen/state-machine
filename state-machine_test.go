package fsm

import (
	"fmt"
	"testing"
)

func getMachine() *Machine {
	machine := NewMachine("close")
	machine.SetTransition("open", "close-door", "close")
	machine.SetTransition("close", "open-door", "open")

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

func TestTransition(t *testing.T) {
	machine := getMachine()

	if err := machine.Transition("open-door"); err != nil {
		t.Error(err)
	}

	if err := machine.Transition("open-door"); err == nil {
		t.Error("expecting door to be opened")
	}
}

func TestCanTransition(t *testing.T) {
	machine := getMachine()

	if !machine.CanTransition("open-door") {
		t.Error("expecting machine.CanTransition(\"open-door\") to return true")
	}

	if machine.CanTransition("close-door") {
		t.Error("expecting machine.CanTransition(\"close-door\") to return false")
	}
}
