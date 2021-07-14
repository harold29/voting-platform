package votingplatform

import (
	"reflect"
	"testing"
	"time"

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

func TestValidateElection(t *testing.T) {
	t.Run("when the election Id is valid and not nil returns true", func(t *testing.T) {
		election := new(Election)
		newUuid, err := uuid.NewRandom()

		if err != nil {
			t.Errorf("Error creating uuid for test")
		}

		got := election.validateElection(ID(newUuid))
		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("when the election Id is nil or empty", func(t *testing.T) {
		election := new(Election)
		nilUuid := uuid.Nil

		got := election.validateElection(ID(nilUuid))
		want := false

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestNewElection(t *testing.T) {
	t.Run("when all the attributes are correct, returns a new election", func(t *testing.T) {
		startDate := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2021, time.Month(1), 2, 0, 0, 0, 0, time.UTC)

		got, err := NewElection("Test Name", "2021", startDate, endDate)

		if err != nil {
			t.Errorf("Error creating election %s", err.Error())
		}

		want := new(Election)

		want.Name = "Test Name"
		want.Year = "2021"
		want.StartDate = startDate
		want.EndDate = endDate
		want.ElectionId = got.ElectionId
		want.Type = "election"

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
