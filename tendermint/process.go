package tendermint

import (
	"log"
	"os"
	"time"
)

type Process struct {
	id     int
	peers  []NetworkPeer
	logger *log.Logger

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

	validValue *int
	validRound int

	// returns the value that the process is proposing
	GetValue func() int
}

type NetworkPeer interface {
	SendVote(msg voteMsg, to int)
	SendProposal(msg proposalMsg, to int)
	SendCommit(msg commitMsg, to int)
}

type NetworkNode interface {
	NetworkPeer

	onProposal(msg proposalMsg, sender int)
	onVote(msg voteMsg, sender int)
	onCommit(msg commitMsg, sender int)
}

func NewProcess() *Process {
	p := &Process{}
	p.logger = log.New(os.Stdout, "process #%d", log.LstdFlags)
	return p
}

func isValidProposal(p proposalMsg) bool {
	return true
}

func (p *Process) StartRound(round int) {
	validatorSet := []int{0, 1, 2, 3}
	F := len(validatorSet)
	majority := 2*F + 1

	proposalMsgs := make(chan proposalMsg)
	timeoutPropose := make(chan bool)

	voteMsgs := make(chan voteMsg)
	votes := []voteMsg{}
	timeoutVote := make(chan bool)

	precommitMsgs := make(chan commitMsg)
	precommits := []commitMsg{}
	timeoutCommit := make(chan bool)

	// NEW HEIGHT.
	//////////////////////////////////////////

	if proposer := getProposerForRound(round, validatorSet); proposer == p.id {
		proposal := proposalMsg{
			round: round,
		}

		// 1. Get the value to propose.
		if p.validValue != nil {
			proposal.value = *p.validValue
		} else {
			proposal.value = p.GetValue()
		}

		// 2. Broadcast proposal
		for _, peer := range p.peers {
			peer.SendProposal(proposal, p.id)
		}
		proposalMsgs <- proposal
	}

	// PROPOSE.
	//////////////////////////////////////////
	go schedTimeoutPropose(timeoutPropose)
	for {
		select {
		case proposal := <-proposalMsgs:
		case <-timeoutPropose:
			break
		}
	}

	// VOTE.
	//////////////////////////////////////////
	go schedTimeoutVote(timeoutVote)
	for {
		select {
		case vote := <-voteMsgs:
			votes = append(votes, vote)
			if majority == len(votes) {
				break
			}
		case <-timeoutVote:
			break
		}
	}

	// PRECOMMIT.
	//////////////////////////////////////////
	go schedTimeoutCommit(timeoutCommit)
	for {
		select {
		case precommit := <-precommitMsgs:
			precommits = append(precommits, precommit)
			if majority == len(votes) {
				break
			}
		case <-timeoutCommit:
			break
		}
	}
}

func schedTimeoutPropose(ch chan<- bool) {
	time.Sleep(proposeTimeout)
	ch <- true
}

func schedTimeoutVote(ch chan<- bool) {
	time.Sleep(prevoteTimeout)
	ch <- true
}

func schedTimeoutCommit(ch chan<- bool) {
	time.Sleep(precommitTimeout)
	ch <- true
}

// NETWORKING.
func (p *Process) addPeer(peer NetworkPeer) {
	p.peers = append(p.peers, peer)
}
