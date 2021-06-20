package votingplatform

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

func (c *Contract) Instantiate() {
	fmt.Println("Instantiated")
}

// func (c *Contract) CreateElection() {
// }

// func (c *Contract) CreateBallot() {
// }

// func GenerateBallot(ctx contractapi.TransactionContextInterface, votableItems VotableItem, election Election, voter Voter) {
	// var pot contractapi.TransactionContextInterface

	// ballot := Ballot{ctx, }
}
