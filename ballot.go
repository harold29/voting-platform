package votingplatform

import (
	"fmt"

	"github.com/google/uuid"
)

type Ballot struct {
	VotableItems []VotableItem `json:"votableItems"`
	Election     Election      `json:"Election"`
	VoterId      string        `json:"voterId"`
	BallotCast   bool          `json:"ballotCast"`
	BallotId     ID            `json:"ballotId"`
	Type         string        `json:"type"`
}

func (b *Ballot) SetBallotId() error {
	euuid, err := uuid.NewRandom()

	if err != nil {
		return fmt.Errorf("error creating a new votable item uuid. %s", err.Error())
	}

	e.BallotId = ID(euuid)

	return nil
}

func NewBallot(ctx TransactionContextInterface, items []VotableItem, election Election, voterId string) (Ballot, error) {
	ballot := new(Ballot)

	// if b.validateBallot()
	ballot.Election = election
	ballot.VotableItems = items
	ballot.VoterId = voterId
	ballot.BallotCast = false
	ballot.SetBallotId()
	ballot.Type = "ballot"

	return ballot, nil
}

func (b *Ballot) validateBallot(ctx TransactionContextInterface, voterId string) bool {
	// voter :=
	// if b.Ballo
}

// func NewBallot(ctx TransactionContextInterface)

// func (b *Ballot) ValidateBallot(voterId ID) {

// }
