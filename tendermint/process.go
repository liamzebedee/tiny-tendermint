package tendermint

type Process struct {
	// the current height of the consensus instance
	height int
	// the current round number
	round int
	// the array of decisions
	decisions []int

	// the current state of the internal Tendermint state machine
	step int

	// the most recent value (with respect to a round number) for which a commit message has been sent
	lockedValue int
	// the last round in which the process sent a commit message that is not nil
	lockedRound int

	validValue int
	validRound int
}

type NetworkNode interface {
	onProposal(msg proposalMsg, sender int)
	onVote(msg voteMsg, sender int)
	onCommit(msg commitMsg, sender int)

	SendVote(msg voteMsg, to int)
	SendProposal(msg proposalMsg, to int)
	SendCommit(msg commitMsg, to int)
}

func NewProcess() *Process {
	return &Process{}
}

func addPeer(peer NetworkNode) {}
