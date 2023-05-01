---
title: "Twitter"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

![Twitter Feed](/iknowkungfoo/newsfeed/twitter-feed.png)

## Main specific discussion points

- [Timeline cannot be computed on read](/iknowkungfoo/systemdesignquestions/newsfeed): 🧠 start by showing with estimations & requirements that it wouldn't scale.
- **Social graph**: RDBMS table to keep track of who follows who, cached for min latency. 🧠 Show difference in latency cache vs RDBMS `SELECT`.
- **Search**: use Lucene to tag documents on write. On search, all search shards must be queried. Why?
- **Active/passive users/Celebrities** get different treatment: precalculate timelines, fan-out vs multi-get on read.
- When discussing timeline, talk about post-processing: filtering the precalculated timeline.
- Optional: Write path has mostly pull but potentially push model in case of push notifications.

## What does Twitter use?

- Storing tweets: [Manhattan](https://blog.twitter.com/engineering/en_us/a/2014/manhattan-our-real-time-multi-tenant-distributed-database-for-twitter-scale), in-house eventually consistent database (with strong-consistency for some workloads), with 3 storage backends: read-only for Hadoop data, LSM tree for high-write, BTree for high-read/low-write. Low-level storage is [Apache BookKeeper](https://bookkeeper.apache.org/docs/overview/). It started with MySQL, then built a MySQL clustering solution, then Manhattan.
- Caching tweets: [Memcached](https://blog.twitter.com/engineering/en_us/a/2012/caching-with-twemcache).
- Caching timelines: Redis.
- Provisioning IDs: [Snowflake](https://github.com/twitter-archive/snowflake/tree/snowflake-2010).

