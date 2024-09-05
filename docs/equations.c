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