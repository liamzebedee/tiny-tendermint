package tendermint

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

}
