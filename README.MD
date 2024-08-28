tiny-tendermint
===============

A tiny (sub 500 LOC) implementation of the Tendermint consensus algorithm for BFT state machine replication in Go.

## Changes.

 - In Tendermint, a block is proposed, then nodes vote to commit on the block, after the vote passes, the nodes then confirm by sending commit messages. Rather than call these terms "prevote" and "precommit" you can just reframe them as "vote" and "commit" messages, where the property of "vote passed" and "block committed" is a reduction of these messages.

