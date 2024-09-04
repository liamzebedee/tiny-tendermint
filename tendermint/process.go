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

func getProposerForRound(round int, validatorSet []int) int {
	return validatorSet[round%len(validatorSet)]
}

func (p *Process) awaitProposal(
	proposalMsgs <-chan proposalMsg,
) {
}

func (p *Process) voteProposal(
	round int,
	proposal *proposalMsg,
) {
	vote := voteMsg{
		value:  -1,
		height: p.height,
		round:  round,
		sender: p.id,
	}
	if proposal != nil && isValidProposal(*proposal) {
		vote.value = proposal.value
	}
	go func() {
		for _, peer := range p.peers {
			peer.SendVote(vote, p.id)
		}
	}()
}

func (p *Process) commitProposal(
	round int,
	value *int,
) {
	vote := commitMsg{
		value:  -1,
		height: p.height,
		round:  round,
		sender: p.id,
	}
	if value != nil {
		vote.value = *value
	}
	go func() {
		for _, peer := range p.peers {
			peer.SendCommit(vote, p.id)
		}
	}()
}

func (p *Process) awaitVotes(
	voteMsgs <-chan voteMsg,
) {
}

func (p *Process) awaitVoteMajority() {
	// if vote.value != proposal.value {
	// 	// ignore.
	// 	continue
	// }

	// votes = append(votes, vote)
	// if majority == len(votes) {
	// 	break
	// }
}

func (p *Process) awaitCommitMajority(
	commitMajorityCh <-chan bool,
) {
	// if commit.value != proposal.value {
	// 	// ignore.
	// 	continue
	// }

	// precommits = append(precommits, commit)
	// if majority == len(precommits) {
	// 	break
	// }
}

func (p *Process) StartRound(round int) {
	validatorSet := []int{0, 1, 2, 3}
	// F := len(validatorSet)
	// majority := 2*F + 1

	proposalMsgs := make(chan proposalMsg)
	timeoutPropose := make(chan bool)

	voteMsgs := make(chan voteMsg)
	var preparedValue *int
	preparedValueChan := make(chan *int)
	// votes := []voteMsg{}
	timeoutVote := make(chan bool)

	// precommitMsgs := make(chan commitMsg)
	// precommits := []commitMsg{}
	commitMajorityCh := make(chan bool)
	timeoutCommit := make(chan bool)

	// NEW HEIGHT.
	//////////////////////////////////////////

	if proposer := getProposerForRound(round, validatorSet); proposer == p.id {
		proposal := proposalMsg{
			value: int(time.Now().UnixMilli()),
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

	var proposal *proposalMsg
	go p.awaitProposal(proposalMsgs)
	go schedTimeoutPropose(timeoutPropose)
	select {
	case proposalR := <-proposalMsgs:
		proposal = &proposalR
		break
	case <-timeoutPropose:
		break
	}

	// VOTE.
	//////////////////////////////////////////

	// 1. Vote for proposal or nil.
	go p.voteProposal(round, proposal)

	// 2. Wait for vote majority or timeout.
	go p.awaitVotes(voteMsgs)
	go schedTimeoutVote(timeoutVote)
	select {
	case v := <-preparedValueChan:
		preparedValue = v
	case <-timeoutVote:
		break
	}

	// COMMIT.
	//////////////////////////////////////////

	// 1. Commit to value.
	go p.commitProposal(round, preparedValue)

	// 2. Wait for commit majority or timeout.
	go p.awaitCommitMajority(commitMajorityCh)
	go schedTimeoutCommit(timeoutCommit)
	select {
	case <-commitMajorityCh:
		// New height.
		// if valid(v) then
		// 	decisionp [hp] = v
		// 	hp←hp+1
		// 	reset lockedRoundp , lockedValuep , validRoundp and validValuep to initial values and empty message log
		// 	StartRound(0)
	case <-timeoutCommit:
		// New round.
		break
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
