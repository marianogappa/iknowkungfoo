---
title: "Partitioning Concepts"
date: 2022-11-20T09:03:20-08:00
draft: false
---

`partition == shard (MongoDB/ElasticSearch) == region (HBase) == tablet (BigTable) == vnode (Cassandra/Riak)`

Typically:
- many partitions per node.
- each partition replicated on many nodes (fault tolerance).

## Why partition?

For scalability: 

- More data capacity that one machine can hold.
- Query throughput scales with the number of machines.

## Explain consistent hashing

[Consistent Hashing article](/iknowkungfoo/systemdesignconcepts/consistent-hashing)

(DDIA recommends calling it `hash partitioning`, because `consistent` is confusing)

## Hotspots

If data is partitioned by `userId` and some users are celebrities, the partition that owns the hash of the celebrity will be a hotspot (skewed workload).

**Fix**: keep track of the few "hot userIds", and whenever a write happens, prefix or postfix a random number (e.g. `rand(10)` or `rand(100)`), then the posts end up on 10/100 different partitions. But reads for celebrity posts need to go to many partitions instead of 1 (in parallel).

## Rebalancing: automatic or manual

- Automatic is convenient but dangerous: if instances rebalancing get too busy and start failing heartbeats, they'll become offline and cause further rebalancing 😱😱😱
- CouchBase, Riak & Voldemort suggest partition assignment, but require a human to approve.

## Secondary indexes

For partitioned data, two flavours:

- **Local to partition**, but then searches scatter-gather to all partitions!
- **Global**, but then it doesn't fit on a machine so it must itself be partitioned.

Note that secondary indexes typically update asynchronously.

## Request routing (or "Service discovery")

Options:

- Client knows which node has the key.
- Client sends to routing layer, who knows which node has the key.
- Client sends to any node, and nodes know who to route to.

Nodes get added & removed all the time, so how to keep knowledge of who owns which key?

- Every node heartbeats every node? 😱 scales quadratically! NO!
- Gossip protocol (e.g. Cassandra).
- External service implementing Raft/Paxos consensus (e.g. ZooKeeper).
