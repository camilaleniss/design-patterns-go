package state

// State
// A pattern which object behavior is determined by its state

// A formalized construct which manages state and transitions is called
// a state machine

import "fmt"

type Switch struct {
	State State
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

// our state machine
type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

// base state allows us if no transition it's set
type BaseState struct{}

func (s *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

// in on state we just can go off
func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	// reassign state
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

// in off state we just can go on
func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	// reassign state
	sw.State = NewOnState()
}

func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
