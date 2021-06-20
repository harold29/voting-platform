package votingplatform

type Ballot struct {
	// votableItems
	election   Election `json:"Election"`
	voterId    ID       `json:"voterId"`
	ballotCast bool     `json:"ballotCast"`
	ballotId   ID       `json:"ballotId"`
}

// func NewBallot(ctx TransactionContextInterface)

// func (b *Ballot) ValidateBallot(voterId ID) {

// }
