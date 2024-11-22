# State Machine

[![Lint commit message](https://github.com/yannickkirschen/state-machine/actions/workflows/commit-lint.yml/badge.svg)](https://github.com/yannickkirschen/state-machine/actions/workflows/commit-lint.yml)
[![Push](https://github.com/yannickkirschen/state-machine/actions/workflows/push.yml/badge.svg)](https://github.com/yannickkirschen/state-machine/actions/workflows/push.yml)
[![Release](https://github.com/yannickkirschen/state-machine/actions/workflows/release.yml/badge.svg)](https://github.com/yannickkirschen/state-machine/actions/workflows/release.yml)
[![GitHub release](https://img.shields.io/github/release/yannickkirschen/state-machine.svg)](https://github.com/yannickkirschen/state-machine/releases/)

State Machine is a Golang library implementing a [finite-state machine](https://en.wikipedia.org/wiki/Finite-state_machine).

## Usage

Basic usage works as follows:

1. Create a machine with an initial state.
2. Set all possible transitions between states.
3. Optional: Set enter end exit actions.
4. Transition from state to state ;)

### Basic example

Here is an example code based on the very simple use-case of a door:

```go
// Signature: <initial state>
machine := fsm.NewMachine("close")

// Signature: <current state, input event, next state>
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

// Signature: <input event>
if err := machine.Transition("open-door"); err != nil {
    panic(err)
}

fmt.Printf("Current state is %s\n", machine.State())
```

### Using complex state types

Define a struct that holds all the data you need. Here is an example containing
the relevant changes based on the door example:

```go
type DoorState struct {
    Id       string
    OpenedBy string
}
```

```go
machine := fsm.NewMachine(&DoorState{Id: "close"})
machine.SetTransition(&DoorState{Id: "open"}, "close-door", &DoorState{Id: "close"})
machine.SetTransition(&DoorState{Id: "close"}, "open-door", &DoorState{Id: "open"})

machine.SetEnterAction(func(last, new *DoorState) error {
    new.OpenedBy = "Peter"
    return nil
})

// ...
```

When working with complex states, there two important notes to consider:

1. Equality of states is checked with `reflect.DeepEqual`. This works both for values and pointers.
2. When changing a state's fields in an action function, a pointer must be used. Otherwise, the change will be lost.
