package main

import (
	"fmt"

	fsm "github.com/yannickkirschen/state-machine"
)

type DoorState struct {
	Id       string
	OpenedBy string
}

func main() {
	machine := fsm.NewMachine(&DoorState{Id: "close"})
	machine.SetTransition(&DoorState{Id: "open"}, "close-door", &DoorState{Id: "close"})
	machine.SetTransition(&DoorState{Id: "close"}, "open-door", &DoorState{Id: "open"})

	machine.SetEnterAction(func(last, new *DoorState) error {
		fmt.Printf("Enter '%s' coming from '%s'\n", new.Id, last.Id)
		new.OpenedBy = "Peter"
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
