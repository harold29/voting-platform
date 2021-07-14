package votingplatform

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestValidateVoter(t *testing.T) {
	t.Run("when the attribute voterId is not nil returns true", func(t *testing.T) {
		voter := new(Voter)

		got := voter.validateVoter("1234")
		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("when the attribute voterId is nil returns false", func(t *testing.T) {
		voter := new(Voter)

		got := voter.validateVoter("")
		want := false

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestValidateRegistrar(t *testing.T) {
	t.Run("when the attribute registrarId is not nil returns true", func(t *testing.T) {
		voter := new(Voter)
		regUuid, err := uuid.NewRandom()

		if err != nil {
			t.Errorf("Error creating resource for test uuid")
		}

		got := voter.validateRegistrar(ID(regUuid))
		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("when the attribute registrarId is nil return false", func(t *testing.T) {
		voter := new(Voter)
		regUuid := uuid.Nil

		got := voter.validateRegistrar(ID(regUuid))
		want := false

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestNewVoter(t *testing.T) {
	t.Run("when all the attributes are validated, returns a Voter object", func(t *testing.T) {
		newUuid, uuidErr := uuid.NewRandom()

		if uuidErr != nil {
			t.Errorf("Error creating new uuid for test")
		}

		got, err := NewVoter("TestName", "TestLastName", "123456", ID(newUuid))
		want := new(Voter)
		want.firstName = "TestName"
		want.lastName = "TestLastName"
		want.voterId = "123456"
		want.registrarId = ID(newUuid)
		want.ballotCreated = false

		if err != nil {
			t.Errorf("Error creating new voter object")
		}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("when the validations fail, returns an error", func(t *testing.T) {
		nilUuid := uuid.Nil

		_, err := NewVoter("TestName", "TestLastName", "", ID(nilUuid))

		if err == nil {
			t.Errorf("Expecting an error creating the Voter but didn't get one")
		}
	})
}
