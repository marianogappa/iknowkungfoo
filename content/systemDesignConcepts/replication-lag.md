---
title: "Replication Lag"
date: 2022-11-20T09:03:20-08:00
draft: false
---

Master-slave replication is eventually consistent, because clients making reads of the same "key" at the same time to the primary or the replicas can get different values (if a write took place, and hasn't been replicated yet).

## Reading Your Own Writes

![Read your own Writes](/iknowkungfoo/replication-lag/read-your-own-writes.png)

**Fix:**

- Client always reads from the leader when it's a read of something that they can have written to.
- Client keeps track of the logical timestamp they wrote at, and sends the timestamp with the read request (logical timestamp as in sequence number; real timestamp would depend on trusting clocks). If replica hasn't arrived to it, wait/redirect/fail.

## Monotonic Reads

`Strong consistency` > `Monotonic Reads` > `Eventual consistency`

![Monotonic Reads](/iknowkungfoo/replication-lag/monotonic-reads.png)

Fix? Each user (e.g. via `userID` hash) does reads always to the same replica. Works unless replica goes offline 🤷‍♂️.

## Consistent Prefix Reads

![Consistent Prefix Reads](/iknowkungfoo/replication-lag/consistent-prefix-reads.png)

**Fix?**

Writes go to the same machine.
