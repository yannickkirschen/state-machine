package main

import (
	"fmt"

	fsm "github.com/yannickkirschen/state-machine/v2"
)

type DoorState struct {
	Id       string
	OpenedBy string
}

func main() {
	machine := fsm.NewMachine[string](&DoorState{Id: "close"})
	machine.SetTransition(&DoorState{Id: "open"}, "close-door", &DoorState{Id: "close"})
	machine.SetTransition(&DoorState{Id: "close"}, "open-door", &DoorState{Id: "open"})

	machine.SetEnterAction(func(lastState, newState *DoorState) error {
		fmt.Printf("Enter '%s' coming from '%s'\n", newState.Id, lastState.Id)
		newState.OpenedBy = "Peter"
		return nil
	})

	machine.SetExitAction(func(current, next *DoorState) error {
		fmt.Printf("Leaving '%s' going to '%s'\n", current.Id, next.Id)
		return nil
	})

	fmt.Printf("Current state is %+v\n", machine.State())

	if _, _, err := machine.Transition("open-door"); err != nil {
		panic(err)
	}

	fmt.Printf("Current state is %+v\n", machine.State())
}
