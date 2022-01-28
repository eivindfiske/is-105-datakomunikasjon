package state_test

import (
	"testing"

	"main.go/event"
	"main.go/state"
)

func Test_state(t *testing.T) {
	event.KyllingtoEast()
	wanted := true
	state := state.KyllingEast

	if state != wanted {
		t.Errorf("Error")
	}
}
func Test_state2(t *testing.T) {
	event.KyllingtoBoat()
	wanted := true
	state := state.KyllingEast

	if state != wanted {
		t.Errorf("Error")
	}
}
