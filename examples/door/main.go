package main

import (
	"fmt"

	fsm "github.com/yannickkirschen/state-machine"
)

func main() {
	machine := fsm.NewMachine("close")
	machine.SetTransition("open", "close-door", "close")
	machine.SetTransition("close", "open-door", "open")

	machine.SetEnterAction(func(last, new string) error {
		fmt.Printf("Enter '%s' coming from '%s'\n", new, last)
		return nil
	})

	machine.SetExitAction(func(current, next string) error {
		fmt.Printf("Leaving '%s' going to '%s'\n", current, next)
		return nil
	})

	fmt.Printf("Current state is %s\n", machine.State())

	if err := machine.Transition("open-door"); err != nil {
		panic(err)
	}

	fmt.Printf("Current state is %s\n", machine.State())
}
