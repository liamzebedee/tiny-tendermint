// ∨ - OR
// ∧ - AND

// Function StartRound(round) :
	// roundp ← round
	// stepp ← propose
	// if proposer(hp, roundp) = p then
		// if validV aluep ̸= nil then
			// proposal ← validV aluep
		// else
			// proposal ← getV alue()
		// broadcast ⟨PROPOSAL, hp, roundp, proposal, validRoundp⟩
	// else
		// schedule OnTimeoutPropose(hp,roundp) to be executed after timeoutPropose(roundp)


/*
upon ⟨PROPOSAL, hp, roundp, v, −1⟩ from proposer(hp, roundp) while stepp = propose do 
	if valid(v) ∧ (lockedRoundp = −1 ∨ lockedValuep = v) then
		broadcast ⟨PREVOTE, hp, roundp, id(v)⟩ 
	else
		broadcast ⟨PREVOTE, hp, roundp, nil⟩ 
	stepp ← prevote

// Handles votes for proposals from previous rounds that 
// validRound: This refers to the round in which a quorum of validators (2f + 1) has already PREVOTED for the value v.
// Current Round (roundp): This is the round the node is currently participating in. It represents the ongoing phase of the consensus process.
// Rationale:
Flexibility with Earlier Rounds: By considering a PROPOSAL with a valid round vr < roundp, Tendermint allows nodes to remain flexible and possibly relock on a value that had significant support in an earlier round. This is crucial in asynchronous or partially synchronous environments where messages might be delayed.
Safety and Liveness: The conditions (lockedRoundp ≤ vr ∨ lockedValuep = v) ensure that the node either relocks on the value v if it was already considering it, or if it was locked on a round earlier than vr. This preserves safety (no conflicting values get locked) while also promoting liveness (the consensus process can continue progressing).
Valid Check: The valid(v) condition ensures that only valid proposals are considered, maintaining the integrity of the consensus process.

upon ⟨PROPOSAL, hp, roundp, v, vr⟩ from proposer(hp, roundp) AND 2f + 1 ⟨PREVOTE, hp, vr, id(v)⟩ while stepp = propose∧(vr ≥ 0∧vr < roundp) do
	if valid(v) ∧ (lockedRoundp ≤ vr ∨ lockedValuep = v) then 
		broadcast ⟨PREVOTE, hp, roundp, id(v)⟩
	else
		broadcast ⟨PREVOTE, hp, roundp, nil⟩
	stepp ← prevote

upon 2f + 1 ⟨PREVOTE, hp, roundp, ∗⟩ while stepp = prevote for the first time do
	schedule OnTimeoutPrevote(hp,roundp) to be executed after timeoutPrevote(roundp)
*/



/*
upon ⟨PROPOSAL, hp, roundp, v, ∗⟩ from proposer(hp, roundp) AND 2f + 1 ⟨PREVOTE, hp, roundp, id(v)⟩ while valid(v) ∧ stepp ≥ prevote for the first time do
	if stepp = prevote then
		lockedValuep ← v
		lockedRoundp ← roundp
		broadcast ⟨PRECOMMIT, hp, roundp, id(v))⟩ stepp ← precommit
		validValuep ← v
		validRoundp ← roundp

upon 2f + 1 ⟨PREVOTE, hp, roundp, nil⟩ while stepp = prevote do
	broadcast ⟨PRECOMMIT, hp, roundp, nil⟩
	stepp ← precommit

upon 2f + 1 ⟨PRECOMMIT, hp, roundp, ∗⟩ for the first time do
	schedule OnTimeoutPrecommit(hp,roundp) to be executed after timeoutPrecommit(roundp)
*/


// NEW HEIGHT.

/*
upon ⟨PROPOSAL, hp, r, v, ∗⟩ from proposer(hp, r) AND 2f + 1 ⟨PRECOMMIT, hp, r, id(v)⟩ while decisionp[hp] = nil do
	if valid(v) then
		decisionp [hp] = v
		hp←hp+1
		reset lockedRoundp , lockedValuep , validRoundp and validValuep to initial values and empty message log
		StartRound(0)
*/

// upon f + 1 ⟨∗, hp, round, ∗, ∗⟩ with round > roundp do
	// StartRound(round)

// Function OnTimeoutPropose(height, round) :
	// if height = hp ∧ round = roundp ∧ stepp = propose then
		// broadcast ⟨PREVOTE, hp, roundp, nil⟩
		// stepp ← prevote