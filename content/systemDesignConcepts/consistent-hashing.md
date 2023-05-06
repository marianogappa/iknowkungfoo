---
title: "Consistent Hashing"
date: 2022-11-20T09:03:20-08:00
draft: false
---

## What does it solve?

- Data doesn't fit on a single machine 🤔: just partition/shard? ✅
- Data is not safe on a single machine 🤔: just replicate? ✅
- As machines are added/removed, gain scalability and don't rebalance everything 🤔: 💥 CONSISTENT HASHING 💥

## Modulo sharding strategy

- If you have `n` servers, to decide where to fetch a key, do `id % n`. Very simple!
- 🤔 if you add machines, or one dies, `n` changes, *all* keys need to be rebalanced!!! 😱😱😱 NO!

## Range sharding strategy

- Hash the `id` to some large number. Assign available shards to a range of the hash space.
- If instance is added/removed/dies, reassign ownerless range to remaining shards.
- 😱 cannot extend range! Careful rebalance when machines added/removed.
- 😱 replication necessary, or instance dying means data is lost!

## Consistent hashing

- Hash `id` to some large number. Assign hash space to a really large number of LOGICAL shards.
- Use a consensus-based system (e.g. ZooKeeper/etcd) to map logical shards to a much smaller number of PHYSICAL shards.
- A shard doesn't map to 1 machine but to a number of machines, for replication.
- As machines added/removed/died, logical shards reassigned to MANY different physical shards.
- Minimise blast radius strategy: keep each physical shard not too busy, so that they can take extra load when others die. Minimise machines involved in rebalancing, to isolate blast.
- Better if some `id`s don't work, than if whole cluster explodes in massive rebalancing effort 💥!
