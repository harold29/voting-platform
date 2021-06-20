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
		return fmt.Errorf("Error creating a new votable item uuid. %s", err.Error())
	}

	e.ElectionId = ID(euuid)

	return nil
}

func (e *Election) GetElectionId() ID {
	return e.ElectionId
}
