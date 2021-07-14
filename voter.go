package votingplatform

import (
	"fmt"

	ledgerapi "votingplatform/ledger-api"
)

type Voter struct {
	VoterId       string `json:"voterId"`
	RegistrarId   string `json:"registrarId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	BallotCreated bool   `json:"ballotCreated"`
	// ObjectType    string `json:"objectType"`
}

func (v *Voter) validateVoter(voterId string) bool {
	if len(voterId) != 0 {
		return true
	}

	return false
}

func (v *Voter) validateRegistrar(registrarId string) bool {
	// if registrarId != ID(uuid.Nil) {
	// 	return true
	// }

	// return false
	return true
}

func (v *Voter) GetSplitKey() []string {
	return []string{v.RegistrarId, v.VoterId}
}

func NewVoter(firstName, lastName, voterId string, registrarId string) (*Voter, error) {
	newVoter := new(Voter)

	if newVoter.validateVoter(voterId) && newVoter.validateRegistrar(registrarId) {
		newVoter.FirstName = firstName
		newVoter.LastName = lastName
		newVoter.VoterId = voterId
		newVoter.RegistrarId = registrarId
		newVoter.BallotCreated = false

		return newVoter, nil
	} else {
		return nil, fmt.Errorf("Voter validation failed")
	}
}

func CreateVoterKey(registrarId, voterId string) string {
	return ledgerapi.MakeKey(registrarId, voterId)
}
