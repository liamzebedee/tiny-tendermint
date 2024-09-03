tiny-tendermint
===============

A tiny (sub 500 LOC) implementation of the Tendermint consensus algorithm for BFT state machine replication in Go.

**Work-in-progress.**

 - [x] Understand intuition of protocol. Understand all key concepts.
 - [x] Document key concepts - process, proposal, propose msg, prevote msg, precommit msg, timeouts, height, round, lock change
 - [ ] Implement basic model without timeouts.
   - [ ] Process thread.
   - [ ] Networking between processes.
   - [ ] Create proposal, vote proposal, commit proposal.
   - [ ] Tests.
 - [ ] Implement model with timeouts.
   - [ ] Propose timeout.
   - [ ] Vote timeout.
   - [ ] Commit timeout.

## Readings.

 - https://arxiv.org/pdf/1807.04938
 - https://jzhao.xyz/thoughts/Tendermint
 - https://groups.csail.mit.edu/tds/papers/Lynch/jacm88.pdf
 - https://tendermint.com/static/docs/tendermint.pdf
 - https://github.com/tendermint/tendermint/blob/main/spec/consensus/consensus.md#common-exit-conditions
 - `docs/tendermint-0.1/` - extracted from the Tendermint git repo at tag `0.1`. Contains v useful docs.

## Changes.

 - In Tendermint, a block is proposed, then nodes vote to commit on the block, after the vote passes, the nodes then confirm by sending commit messages. Rather than call these terms "prevote" and "precommit" you can just reframe them as "vote" and "commit" messages, where the property of "vote passed" and "block committed" is a reduction of these messages.
