package votingplatform

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Election struct {
	ElectionId ID        `json:"electionId"`
	Name       string    `json:"name"`
	Year       string    `json:"year"`
	StartDate  time.Time `json:"startTime"`
	EndDate    time.Time `json:"endTime"`
	Type       string    `json:"type"`
}

func (e *Election) SetElectionId() error {
	euuid, err := uuid.NewRandom()

	if err != nil {
		return fmt.Errorf("error creating a new votable item uuid. %s", err.Error())
	}

	valid := e.validateElection(ID(euuid))

	if !valid {
		return fmt.Errorf("error during uuid validation on set election id")
	}

	e.ElectionId = ID(euuid)

	return nil
}

func (e *Election) GetElectionId() ID {
	return e.ElectionId
}

func (e *Election) validateElection(electionId ID) bool {
	return electionId != ID(uuid.Nil)
}

func NewElection(name, year string, startDate, endDate time.Time) (*Election, error) {
	election := new(Election)

	err := election.SetElectionId()

	if err != nil {
		return nil, fmt.Errorf("error during the creation of a new election: %s", err.Error())
	}

	election.Name = name
	election.Year = year
	election.StartDate = startDate
	election.EndDate = endDate
	election.Type = "election"

	return election, nil
}
