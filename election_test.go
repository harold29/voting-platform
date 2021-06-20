package votingplatform

import (
	"testing"

	"github.com/google/uuid"
)

func TestSetElectionId(t *testing.T) {
	election := new(Election)

	election.SetElectionId()
	got := election.GetElectionId()

	if got == ID(uuid.Nil) {
		t.Errorf("got nil want UUID")
	}
}

func TestGetElectionId(t *testing.T) {
	election := new(Election)

	election.SetElectionId()
	got := election.GetElectionId()
	want := election.ElectionId

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
