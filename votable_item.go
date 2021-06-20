package votingplatform

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

type VotableItem struct {
	votableItemId ID     `metadata:"votableItemId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	count         int    `json:"count"`
	voterId       ID     `json:"voterId"`
	key           string `metadata:"key"`
	class         string `metadata:"class"`
	// state 		string `metadata:"currentState"`
}

type votableItemAlias VotableItem
type jsonVotableItem struct {
	*votableItemAlias
	VotableItemId string `json:"votableItemId"`
	VoterId       string `json:"voterId"`
	Count         int    `json:"count"`
	// key   string `json:"key"`
	// class string `json:"class"`
}

func (vi VotableItem) MarshalJSON() ([]byte, error) {
	jvi := jsonVotableItem{votableItemAlias: (*votableItemAlias)(&vi), VotableItemId: vi.GetVotableItemId().String(), VoterId: vi.GetVoterId().String(), Count: vi.GetVoteCount()}

	return json.Marshal(&jvi)
}

func (vi *VotableItem) UnmarshalJSON(data []byte) error {
	jvi := jsonVotableItem{votableItemAlias: (*votableItemAlias)(vi)}

	err := json.Unmarshal(data, &jvi)

	if err != nil {
		return err
	}

	vi.votableItemId = ID(uuid.MustParse(jvi.VotableItemId))
	vi.voterId = ID(uuid.MustParse(jvi.VoterId))
	vi.count = jvi.Count

	return nil
}

func (vi *VotableItem) SetVotableItemId() error {
	viuuid, err := uuid.NewRandom()

	if err != nil {
		return fmt.Errorf("Error creating a new votable item uuid. %s", err.Error())
	}

	vi.votableItemId = ID(viuuid)

	return nil
}

func (vi *VotableItem) GetVotableItemId() ID {
	return vi.votableItemId
}

func (vi *VotableItem) GetVoterId() ID {
	return vi.voterId
}

func (vi *VotableItem) SetName(name string) {
	vi.Name = name
}

func (vi *VotableItem) GetVoteCount() int {
	return vi.count
}

func (vi *VotableItem) AddVote() {
	vi.count++
}

func (vi *VotableItem) SetDescription(description string) {
	vi.Description = description
}

func (vi *VotableItem) Serialize() ([]byte, error) {
	return json.Marshal(vi)
}

func Deserialize(bytes []byte, vi *VotableItem) error {
	err := json.Unmarshal(bytes, vi)

	if err != nil {
		return fmt.Errorf("Error deserializing vote item, %s.", err.Error())
	}

	return nil
}
