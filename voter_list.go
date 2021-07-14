package votingplatform

import ledgerapi "votingplatform/ledger-api"

type VoterListInterface interface {
	AddVoter(*Voter) error
	GetVoter(string, string) (*Voter, error)
	UpdateVoter(*Voter) error
}

type list struct {
	voterList ledgerapi.StateListInterface
}

func (vl *list) AddVoter(vt *Voter) error {
	return vl.voterList.AddState(vt)
}

func (vl *list) GetVoter(voterId, registrarID string) (*Voter, error) {
	vt := new(Voter)

	err := vl.voterList.GetState(CreateVoterKey(registrarID, voterId), vt)

	if err != nil {
		return nil, err
	}

	return vt, nil
}

func (vl *list) UpdateVoter(voter *Voter) error {
	return vl.voterList.UpdateState(voter)
}

func newList(ctx TransactionContextInterface) *list {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "testing.voteplatform.voterlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*Voter))
	}

	list := new(list)
	list.voterList = stateList

	return list
}
