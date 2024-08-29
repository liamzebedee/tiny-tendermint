tiny-tendermint
===============

A tiny (sub 500 LOC) implementation of the Tendermint consensus algorithm for BFT state machine replication in Go.

**Work-in-progress.**

## Changes.

 - In Tendermint, a block is proposed, then nodes vote to commit on the block, after the vote passes, the nodes then confirm by sending commit messages. Rather than call these terms "prevote" and "precommit" you can just reframe them as "vote" and "commit" messages, where the property of "vote passed" and "block committed" is a reduction of these messages.


The idea of Goliath-
horizontally scalable read-write chunk log.
fault tolerance:
master - coordinates everything, allocates chunks, checks liveness, allocates file table
workers - handle shards of the dataset

basically users write to the primary and then the replicas
once they are done

how do we ensure replicas are live?
some sort of chubby blockchain server or something
just make the master do that
