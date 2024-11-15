package main

import (
	"fmt"

	fsm "github.com/yannickkirschen/state-machine"
)

func main() {
	machine := fsm.NewMachine("close")
	machine.SetTransition("open", "close-door", "close")
	machine.SetTransition("close", "open-door", "open")

	machine.SetEnterAction(func(last, new fsm.State) error {
		fmt.Printf("Enter '%s' coming from '%s'\n", new, last)
		return nil
	})

	machine.SetExitAction(func(current, next fsm.State) error {
		fmt.Printf("Leaving '%s' going to '%s'\n", current, next)
		return nil
	})

	if err := machine.Transition("open-door"); err != nil {
		panic(err)
	}
}
