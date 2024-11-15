package fsm

import "fmt"

type State string
type Event string

type StateActionFunc func(State, State) error

type TransitionEvent struct {
	Current State
	Input   Event
}

type Machine struct {
	transitions map[TransitionEvent]State
	current     State

	enterAction StateActionFunc
	exitAction  StateActionFunc
}

func NewMachine(initial State) *Machine {
	return &Machine{
		transitions: map[TransitionEvent]State{},
		current:     initial,
	}
}

func (machine *Machine) SetTransition(current State, input Event, next State) {
	machine.transitions[TransitionEvent{Current: current, Input: input}] = next
}

func (machine *Machine) SetEnterAction(f StateActionFunc) {
	machine.enterAction = f
}

func (machine *Machine) SetExitAction(f StateActionFunc) {
	machine.exitAction = f
}

func (machine *Machine) Transition(input Event) error {
	next, ok := machine.transitions[TransitionEvent{Current: machine.current, Input: input}]
	if !ok {
		return fmt.Errorf("there is no state to transition to from state '%s' on event '%s'", machine.current, input)
	}

	if machine.exitAction != nil {
		if err := machine.exitAction(machine.current, next); err != nil {
			return err
		}
	}

	if machine.enterAction != nil {
		if err := machine.enterAction(machine.current, next); err != nil {
			return err
		}
	}

	// We will only enter the next state if the enter action was successful.
	machine.current = next
	return nil
}
