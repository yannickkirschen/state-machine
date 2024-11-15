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

Here is an example code based on the very simple use-case of a door:

```go
// Signature: <initial state>
machine := fsm.NewMachine("close")

// Signature: <current state, input event, next state>
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

// Signature: <input event>
if err := machine.Transition("open-door"); err != nil {
    panic(err)
}
```
