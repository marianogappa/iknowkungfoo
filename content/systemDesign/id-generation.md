---
title: "ID Generation"
date: 2022-11-20T09:03:20-08:00
draft: false
---
## Intuition

Ideally, unique ID generation is delegated to a RDBMS, e.g. MySQL's PKs. At scale, this is a bottleneck. How do big companies do it?

## Requirements

- Correct: always unique!
- Performance: many ids per second (> to RDBMS can do, e.g. 100k/s), grows over time, few ms per id.
- Reliability: if DB dies or network disconnects, needs to keep working.
- K-Sortability: ids need to be directly sortable (i.e. without the content) and always be roughly `>` than previous ids to a ~1s precision.
- Complexity: system not too complex, least moving parts.
- Size: ids not too large, most systems use 64 bits (UUID => 128 bits).

## Flickr (2010)

[DB Ticket Servers](https://code.flickr.net/2010/02/08/ticket-servers-distributed-unique-primary-keys-on-the-cheap/)

Simply have 2 MySQL instances producing even & odd numbers behind a load balancer, and if an instance is unresponsive have a backup ready.

Note that sortability may suffer if one instance is faster than the other (or the network).

## Twitter (2010)

[Twitter blogpost](https://blog.twitter.com/engineering/en_us/a/2010/announcing-snowflake)
[Twitter Snowflake Github README](https://github.com/twitter-archive/snowflake/tree/snowflake-2010)

- Many servers generate 64-bit ids: 41 bit millisecond-precision timestamp, 10-bit machine id, 12-bit autoincrement (1 reserved??).
- Autoincrement % 2^12 with protection for same millisecond rollover.
- NTP to keep server clocks accurate, with protection for non-monotonic (i.e. going back in time to correct).
- ZooKeeper to choose worker numbers (but not used for ID coordination between nodes).

## Instagram (2012)

[Sharding IDs at Instagram](https://instagram-engineering.com/sharding-ids-at-instagram-1cf5a71e5a5c)

- Many servers generate 64-bit ids: 41 bit millisecond-precision timestamp, 13-bit machine id, 10-bit autoincrement.
- Rather than special purpose, use Postgres to generate.
- Use Postgres schemas to separate logical shards (each shard is a machine-id) but keep them in fewer physical shards, and move schemas across to scale.    
