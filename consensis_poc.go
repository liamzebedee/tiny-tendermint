// GOAT ---- https://arxiv.org/pdf/1807.04938
// https://jzhao.xyz/thoughts/Tendermint
// https://groups.csail.mit.edu/tds/papers/Lynch/jacm88.pdf
// https://tendermint.com/static/docs/tendermint.pdf
package tendermint

import "fmt"

// timeout constnats.
// The timeouts prevent the algorithm from blocking and waiting forever for some condition to be true, ensure that processes continuously transition between rounds, and guarantee that eventually (after GST) communication between correct processes is timely and reliable so they can decide
// The last role is achieved by increasing the timeouts with every new round r, i.e, timeoutX(r) = initT imeoutX + r ∗ timeoutDelta; they are reset for every new height (consensus instance).
const (
	proposeTimeout = 1000
	prevoteTimeout = 1000
	precommitTimeout = 1000
)

func getProposerForRound(i int) int {
	return 0
}
func getValue() *int {
	v := 0
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
	
	// the most recent value (with respect to a round number) for which a PRECOMMIT message has been sent
	lockedValue int
	// the last round in which the process sent a PRECOMMIT message that is not nil
	lockedRound int

	validValue int
	validRound int
}

func process() {
}

func process_startRound() {
	// Propose
	// schedule a timeout for the propose message
}

func upon_proposal() {
	// broacast prevote for the current locked value or nil
}

func upon_proposalAndPrevote() {
	// send prevote
}

// 
// Pre-vote: Validators broadcast their votes for a proposed block. If a block is proposed and receives more than two-thirds of the votes, it moves to the next round.
// 
// Pre-commit: Validators broadcast pre-commits for the block that got sufficient pre-votes. If two-thirds of the validators pre-commit to a block, it's locked.
// 
// Commit: If a block receives enough pre-commits, it's committed as the finalized block. If it fails, the process restarts.
// 

func tendermint_run() {
	height := 0
	round := 0
	
	for {
		fmt.Printf("h=%d r=%d\n", height, round)
		round++
		height++

		if height == 10 {
			break
		}
	}
}