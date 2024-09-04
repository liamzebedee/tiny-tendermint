package tendermint

// timeout constnats.
// The timeouts prevent the algorithm from blocking and waiting forever for some condition to be true, ensure that processes continuously transition between rounds, and guarantee that eventually (after GST) communication between correct processes is timely and reliable so they can decide
// The last role is achieved by increasing the timeouts with every new round r, i.e, timeoutX(r) = initT timeoutX + r * timeoutDelta; they are reset for every new height (consensus instance).
const (
	proposeTimeout   = 1000
	prevoteTimeout   = 1000
	precommitTimeout = 1000
)

const (
	STEP_PROPOSE = iota
	STEP_VOTE
	STEP_PRECOMMIT
)

type RoundState = int
