package consensus

import (
	"testing"

	_ "github.com/tendermint/tendermint/config/tendermint_test"
)

func TestEnterProposeNoPrivValidator(t *testing.T) {
	cs, _ := randConsensusState()
	cs.EnterPropose(1, 0)
	rs := cs.GetRoundState()
	if rs.Proposal != nil {
		t.Error("Expected to make no proposal, since no privValidator")
	}
}

func TestEnterPropose(t *testing.T) {
	cs, privValidators := randConsensusState()
	val0 := privValidators[0]
	cs.SetPrivValidator(val0)

	cs.EnterPropose(1, 0)
	rs := cs.GetRoundState()

	// Check that Proposal, ProposalBlock, ProposalBlockParts are set.
	if rs.Proposal == nil {
		t.Error("rs.Proposal should be set")
	}
	if rs.ProposalBlock == nil {
		t.Error("rs.ProposalBlock should be set")
	}
	if rs.ProposalBlockParts.Total() == 0 {
		t.Error("rs.ProposalBlockParts should be set")
	}
}

// TODO write better consensus state tests
