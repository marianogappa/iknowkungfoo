---
title: "Uber"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

# Main discussion points specific to Uber

- Read/Write throughput is huge: cabs & riders need to broadcast location constantly (websockets are a must for full-duplex). Recent data must be in-memory for quick access, but also pushed to Kafka for batch processing to improve maps & ETA.
- How to scale with so many cabs/riders? Divide Earth in cells with S2 & geofence cities, partition this way. A cab/rider has a lat-long which can be mapped to a cell, which determines which partition has them (consistent hashing cluster).
- There will be a Websocket manager mapping users to hosts.
- How does the dispatch operation work? Talk about cabs within radius of rider, sort based on ETA and other factors.
- How does ETA work? Talk about running Dijkstra on the map's digraph, where the weights depend on distance but also traffic, weather, etc, and should be cached (80/20) and predicted. Dijkstra is too expensive, so graph has to be partitioned and calculations combined & aggregated across them.
- Discuss calculating cost? Distance * (gas/km) + Elapsed time * (wage/hour) + Service fee (include special days / late night / service tier)
- GPS-Map matching: GPS gives inaccurate results in general, and specifically in urban canyons (two many skyscrapers) or sometimes sparse datapoints. Apparently there's a model called "Hidden Markov Model (HMM)" and a dynamic programming algorithm called Viterbi for this.

# What does Uber do?

- Uses [Google S2 library](https://opensource.googleblog.com/2017/12/announcing-s2-library-geometry-on-sphere.html) for building global geometric databases.
- Uses ["Schemaless"](https://www.uber.com/en-JP/blog/schemaless-part-two-architecture/) on top of MySQL: no schema (JSON objects), high write/read throughput, horizontal scaling, fault-tolerance, immutable (updates by inserting new version). "Each Schemaless shard is a separate MySQL database, and each MySQL database server contains a set of MySQL databases."
- In 2019 they [migrated from InnoDB to MyRocks](https://www.uber.com/en-JP/blog/mysql-to-myrocks-migration-in-uber-distributed-datastores/)
- [Apparently](https://www.uber.com/en-JP/blog/schemaless-sql-database/), Schemaless wasn't general-purpose so efficiency suffered, then they tried Cassandra but it "lacked operational maturity at scale" and the eventual consistency impeded developer efficiency, complicating application architecture, so they created Docstore (2021) which can enforce schemas and have strict linearizability per partition.

# Video tutorials online

I watched all tutorials and I'm only linking to useful ones; they all miss points and add relevant points, so you really need to watch them all to get the full picture:

- [Tech Dummies: UBER System design | OLA system design | uber architecture | amazon interview question](https://www.youtube.com/watch?v=umWABit-wbk&t=1s&ab_channel=TechDummiesNarendraL)
- [codeKarle: Uber System Design | Ola System Design | System Design Interview Question - Grab, Lyft](https://www.youtube.com/watch?v=Tp8kpMe-ZKw&t=427s&ab_channel=codeKarle)
- [Success in Tech: System Design: Uber Lyft ride sharing services - Interview question](https://www.youtube.com/watch?v=J3DY3Te3A_A&t=2357s&ab_channel=SuccessinTech)

# Other sources

[High Scalability article: Designing Uber](http://highscalability.com/blog/2022/1/25/designing-uber.html)
