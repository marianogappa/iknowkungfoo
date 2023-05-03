---
title: "Sloppy Quorum And Hinted Handoff"
date: 2022-11-20T09:03:20-08:00
draft: false
---

## Strong consistency

- e.g. Postgres on one node. If there are concurrent writes & concurrent reads, whether all clients see the same value for the same key depends on the Isolation Level, but assuming a strong isolation level they will.
- This 👀 the ideal scenario. Why not use it always? It doesn't scale, and if machine dies, all is lost.

## Slightly weaker consistency

- e.g. a master-slave replication scheme on top of MySQL. Only the master processes writes, and one or more replicas process reads in parallel.
- Now reads scale, but writes still don't, and new issues!
- Depending on replication lag, clients connected to different replicas see different values for the same key!
- If master dies, system is read-only until new master chosen.
- 🧠: linearizability is still guaranteed, because only one node processes writes.

## Read Quorums & Write Quorums in NoSQL replication

[This is a good explanation](https://scalablehuman.com/2021/10/24/distributed-data-intensive-systems-reading-and-writing-quorums/#:~:text=Reads%20and%20writes%20that%20obey%20this%20R%20and,to%20be%20valid.%20In%20a%20Dynamo%20style%20databases%3A).

In leader-based replication, there is one primary node and some replicas responsible for handling a read/write/delete of a given key.

Defining the `w`, `r` and `n` numbers for replication:

- `w`: number of nodes that must acknowledge a write for the success to be sent to the client (client blocked until then).
- `r`: number of nodes that must be asked for the value of a given key; if they return different values, there will be a strategy to pick winner (e.g. Last Write Wins in Cassandra, let application decide in DynamoDB).
- `n`: total nodes that are concerned with a given key (<= total nodes in cluster).

Typically `w` + `r` > `n`.

All `n` nodes that are concerned with a given key *eventually* converge to the same value.

👀: If less than `w` nodes are online (of those `n`), that key is read-only!!

## Sloppy Quorum And Hinted Handoff

⚠️ This is either not supported or disabled by default everywhere except Riak & DynamoDB ⚠️

- If less than `w` nodes are reachable (of those `n`), but there are other nodes in the cluster to fill in, accept the write anyway (sloppy quorum) and forward them to the right nodes later when they are reachable again.
- This strategy maximises availability over consistency, typically on cross-datacenter temporary network issues, but what happens to consistency if the nodes that were unreachable to this client were on for everyone else, serving reads & writes concurrently? 😱
