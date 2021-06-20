package votingplatform

type Voter struct {
	voterId       ID     `json:"voterId"`
	registrarId   ID     `json:"registrarId"`
	firstName     string `json:"firstName"`
	lastName      string `json:"lastName"`
	ballotCreated bool   `json:"ballotCreated"`
	// ObjectType    string `json:"objectType"`
}
