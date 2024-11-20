package fsm

import "fmt"

type State string
type Event string

// StateActionFunc is a callback function that will be called when entering or
// leaving a state.
type StateActionFunc func(State, State) error

type transitionEvent struct {
	Current State
	Input   Event
}

// Machine is a finite-state machine.
type Machine struct {
	transitions map[transitionEvent]State
	current     State

	enterAction StateActionFunc
	exitAction  StateActionFunc
}

// NewMachine creates a new machine with an initial state and no transitions.
func NewMachine(initial State) *Machine {
	return &Machine{
		transitions: map[transitionEvent]State{},
		current:     initial,
	}
}

// SetTransition defines an allowed transition from a current state to a next
// state on a specific event.
func (machine *Machine) SetTransition(current State, input Event, next State) {
	machine.transitions[transitionEvent{Current: current, Input: input}] = next
}

// SetEnterAction defines a callback function for entering a state.
func (machine *Machine) SetEnterAction(f StateActionFunc) {
	machine.enterAction = f
}

// SetExitAction defines a callback function for leaving a state.
func (machine *Machine) SetExitAction(f StateActionFunc) {
	machine.exitAction = f
}

// Transition performs a transition for a given event. If the transition is not
// possible, an error is returned.
// When leaving the old state, the exit action is called (if provided). When
// entering the new state, the enter action is called (if provided).
func (machine *Machine) Transition(input Event) error {
	next, ok := machine.transitions[transitionEvent{Current: machine.current, Input: input}]
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

// CanTransition checks if the machine can transition to a new state when the
// given event occurs.
// Important note: Machine methods are not thread-safe! When multiple processes
// change states on a machine, it may happen that CanTransition returns true but
// when calling Transition afterwards, an error is returned because meanwhile
// another process already made that transition.
func (machine *Machine) CanTransition(input Event) bool {
	_, ok := machine.transitions[transitionEvent{Current: machine.current, Input: input}]
	return ok
}
