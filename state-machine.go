package fsm

import (
	"fmt"
	"reflect"
)

type Event string

// Transition is a three-tuple of a transition in a finite-state machine
// containing a current state, an event and a next state that they lead to.
type Transition[T comparable] struct {
	Current T
	Input   Event
	Next    T
}

// StateActionFunc is a callback function that will be called when entering or
// leaving a state.
type StateActionFunc[T comparable] func(T, T) error

// Machine is a finite-state machine handling states of a given comparable type.
type Machine[T comparable] struct {
	transitions []*Transition[T]
	current     T

	enterAction StateActionFunc[T]
	exitAction  StateActionFunc[T]
}

// NewMachine creates a new machine with an initial state and no transitions.
func NewMachine[T comparable](initial T) *Machine[T] {
	return &Machine[T]{
		transitions: []*Transition[T]{},
		current:     initial,
	}
}

// State returns the current state.
func (machine *Machine[T]) State() T {
	return machine.current
}

// SetTransition defines an allowed transition from a current state to a next
// state on a specific event.
func (machine *Machine[T]) SetTransition(current T, input Event, next T) {
	transition := machine.findTransition(current, input)
	if transition == nil {
		machine.transitions = append(machine.transitions, &Transition[T]{
			Current: current,
			Input:   input,
			Next:    next,
		})

		return
	}

	transition.Current = current
	transition.Input = input
}

// SetEnterAction defines a callback function for entering a state.
func (machine *Machine[T]) SetEnterAction(f StateActionFunc[T]) {
	machine.enterAction = f
}

// SetExitAction defines a callback function for leaving a state.
func (machine *Machine[T]) SetExitAction(f StateActionFunc[T]) {
	machine.exitAction = f
}

// Transition performs a transition for a given event. If the transition is not
// possible, an error is returned.
// When leaving the old state, the exit action is called (if provided). When
// entering the new state, the enter action is called (if provided).
func (machine *Machine[T]) Transition(input Event) (*T, *T, error) {
	transition := machine.findTransition(machine.current, input)
	if transition == nil {
		return nil, nil, fmt.Errorf("there is no state to transition to from state '%+v' on event '%s'", machine.current, input)
	}

	next := transition.Next

	if machine.exitAction != nil {
		if err := machine.exitAction(machine.current, next); err != nil {
			return nil, nil, err
		}
	}

	if machine.enterAction != nil {
		if err := machine.enterAction(machine.current, next); err != nil {
			return nil, nil, err
		}
	}

	// We will only enter the next state if the enter action was successful.
	last := copy(machine.current)
	machine.current = next
	return &last, &machine.current, nil
}

// CanTransition checks if the machine can transition to a new state when the
// given event occurs.
// Important note: Machine methods are not thread-safe! When multiple processes
// change states on a machine, it may happen that CanTransition returns true but
// when calling Transition afterwards, an error is returned because meanwhile
// another process already made that transition.
func (machine *Machine[T]) CanTransition(input Event) bool {
	return machine.findTransition(machine.current, input) != nil
}

func (machine *Machine[T]) findTransition(current T, input Event) *Transition[T] {
	for _, transition := range machine.transitions {
		if reflect.DeepEqual(transition.Current, current) && transition.Input == input {
			return transition
		}
	}

	return nil
}

func copy[T any](t T) T {
	return t
}
