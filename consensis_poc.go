// GOAT ---- https://arxiv.org/pdf/1807.04938
// https://jzhao.xyz/thoughts/Tendermint
// https://groups.csail.mit.edu/tds/papers/Lynch/jacm88.pdf
// https://tendermint.com/static/docs/tendermint.pdf
package tendermint

import (
	"fmt"
	"time"
)

// timeout constnats.
// The timeouts prevent the algorithm from blocking and waiting forever for some condition to be true, ensure that processes continuously transition between rounds, and guarantee that eventually (after GST) communication between correct processes is timely and reliable so they can decide
// The last role is achieved by increasing the timeouts with every new round r, i.e, timeoutX(r) = initT timeoutX + r * timeoutDelta; they are reset for every new height (consensus instance).
const (
	proposeTimeout   = 1000
	prevoteTimeout   = 1000
	precommitTimeout = 1000
)

func getProposerForRound(i int, validatorSet []int) int {
	return validatorSet[0]
}

func getValue() *int {
	MAGIC_SEED := 42
	v := MAGIC_SEED
	return &v
}

type processState struct {
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

type proposalMsg struct {
	value  int
	sender int
}
type voteMsg struct {
	sender int
}
type commitMsg struct {
	sender int
}

//
// Pre-vote: Validators broadcast their votes for a proposed block. If a block is proposed and receives more than two-thirds of the votes, it moves to the next round.
//
// Pre-commit: Validators broadcast pre-commits for the block that got sufficient pre-votes. If two-thirds of the validators pre-commit to a block, it's locked.
//
// Commit: If a block receives enough pre-commits, it's committed as the finalized block. If it fails, the process restarts.
//

const (
	ROUND_STATE_PROPOSE = iota
	ROUND_STATE_VOTE
	ROUND_STATE_PRECOMMIT
)

type RoundState = int

func tm_run_rounds(validatorId int, rounds int, network Network) {
	height := 0
	round := 0

	nvalidators := 4
	validatorsSet := []int{}
	for i := 0; i < nvalidators; i++ {
		validatorsSet = append(validatorsSet, i)
	}

	roundState := ROUND_STATE_PROPOSE
	for {
		if roundState == ROUND_STATE_PROPOSE {
			// Select proposer.
			proposer := getProposerForRound(round, validatorsSet)

			// If we are proposer, send proposal.
			if proposer == validatorId {
				proposal := proposalMsg{
					value:  *getValue(),
					sender: validatorId,
				}
				// Broadcast proposal.
				network.broadcast(proposal)
			}

			// Wait for proposal.
			time.Sleep(proposeTimeout)
		}
		if roundState == ROUND_STATE_VOTE {
			// Get proposal ID.
			// Decide whether to vote or not.
		}
		if roundState == ROUND_STATE_PRECOMMIT {

		}

		fmt.Printf("h=%d r=%d proposer=%d\n", height, round, proposer)

		round++
		height++

		if height == 10 {
			break
		}
	}
}

type Network interface {
	broadcast(msg interface{})
	recv() interface{}
}

type Process struct {
	state   processState
	network Network
}

func tendermint_run() {
	// setup network
	// setup 4 processes each with network
	// go process.run()
	// process1 should propose a block
	// it will getValue, create a proposal, and broadcast it
	// process2-4 should receive proposal, and if they don't, a timeout should occur wherein they vote for a nil proposal
	// process2-4 should then broadcast their votes
	// upon 2/3 votes for a proposal, the proposal is now accepted and we move to commit it
	// process2-4 should then broadcast their commit votes
	// if 2/3 commit votes are received, the block is committed
	// we increment the height, and start a new round
	// else we start a new round
	// each round we do timeouts
	// and the timeouts are pure functions of the (height, round)
	// where the timeouts reset for each height, but increase in the face of no progress (increasing round id)
	go tm_run_rounds(0, 5)
	go tm_run_rounds(1, 5)
	go tm_run_rounds(2, 5)
	go tm_run_rounds(3, 5)
}

// Ok this is great.
// Now how do we bundle this library?
// Simply:
// - setup network
// - setup process
// - processs.sync which downloads decisions list
// - now run tendermint consensus instance
// - now we can use this for running a server for example. just use tendermint to elect a master.

// Each round begins with a proposal
//
// Once a complete proposal is received by a validator, it signs a pre-vote for that proposal and broadcasts it to the network. If a validator does not receive a correct proposal within ProposalTimeout, it pre-votes for nil instead.

// If a proposal is received for a lower round, or from an incorrect proposer, it is rejected
