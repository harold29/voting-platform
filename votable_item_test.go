package votingplatform

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestSetVotableItemId(t *testing.T) {
	votableItem := new(VotableItem)

	votableItem.SetVotableItemId()
	got := votableItem.GetVotableItemId()

	if got == ID(uuid.Nil) {
		t.Errorf("got nil want UUID")
	}
}

func TestSetName(t *testing.T) {
	votableItem := new(VotableItem)

	votableItem.SetName("test_name")

	got := votableItem.Name
	want := "test_name"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAddVote(t *testing.T) {
	votableItem := new(VotableItem)

	for i := 0; i < 5; i++ {
		votableItem.AddVote()
	}

	got := votableItem.GetVoteCount()
	want := 5

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSetDescription(t *testing.T) {
	votableItem := new(VotableItem)

	votableItem.SetDescription("test description")

	got := votableItem.Description
	want := "test description"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSerialize(t *testing.T) {
	vi := new(VotableItem)
	vi.Name = "test name"
	vi.Description = "test description"
	vi.AddVote()

	got, err := vi.Serialize()
	want := `{"name":"test name","description":"test description","votableItemId":"00000000-0000-0000-0000-000000000000","voterId":"00000000-0000-0000-0000-000000000000","count":1}`

	if err != nil {
		t.Errorf("got an error during serializing")
	}

	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDeserialize(t *testing.T) {
	var vi *VotableItem
	var err error

	validJSON := `{"name":"test name","description":"test description","votableItemId":"00000000-0000-0000-0000-000000000000","voterId":"00000000-0000-0000-0000-000000000000","count":1}`
	expectedVi := new(VotableItem)
	expectedVi.Name = "test name"
	expectedVi.Description = "test description"
	expectedVi.AddVote()

	vi = new(VotableItem)
	err = Deserialize([]byte(validJSON), vi)

	if err != nil {
		t.Errorf("got an error during deserializing")
	}

	if reflect.DeepEqual(vi, expectedVi) == false {
		t.Errorf("The deserialized Votable Item is not equal to the original votable item")
	}

}
