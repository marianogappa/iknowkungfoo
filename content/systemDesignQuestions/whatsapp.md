---
title: "WhatsApp"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

## Non-functional requirement

- Latency: once sent, messages should show in the chat immediately for the sender and very quickly for the receiver
- It’s more important that the acknowledged messages are never lost, than that they arrive at the destination quicker.
- It’s ok if there’s a small race condition in which messages sent by both clients ~ at the same time may show up differently on both clients, and potentially reordered on a refresh.

## Back-of-the-envelope estimations

- Usually systems are read heavy (e.g. Twitter, Instagram, TikTok, Facebook, Netflix), but in this case traffic is mostly 1:1.

## Websockets

- Latency nf requirement makes traditional request-response unfeasible. HTTP polling inefficient. Long polling is not full-duplex.
- Usually whenever websockets and lots of connections are involved, it's a good idea to mention the ["thundering herd"](https://en.wikipedia.org/wiki/Thundering_herd_problem) problem, and having spare capacity on each server and isolating with partitions to prevent disaster amplification.

## How do the different users connect to each other for sending messages?

- Websocket handlers talk to some Session service that keeps a hot mapping of users to handlers. Mapping will need to be in-memory, so something like Redis/Memcached.
- Messages may not be able to reach receiver if they're offline, so they must be kept server-side at least until reception, but could be permanently. Reliability requires that messages are stored before the ✅ appears.
- How to store messages? Most tutorials recommend Cassandra & HBase because of the high write throughput. However, in practice WhatsApp uses Mnesia, Facebook Messenger uses MySQL (MyRocks), Slack uses Vitess (MySQL) 🤷‍♂️. Probably best to focus on the actual schema.
- To do online presence, keep `(userId, lastSeenTimestamp)` in memory. XMPP protocol dictates how to do "is typing" and "online presence", but will require the `Session service` & `Websocket handlers` to relay messages about these events using some form of `Pub/Sub` configuration (since certain users, or consumers, are interested in the typing & online event changes of other users, or producers).


## What does Whatsapp use?

- Erlang for the backend, Mnesia for the storage (Mnesia being Erlang's storage layer). Follows XMPP protocol.
- Uses [IBM Cloud](https://en.wikipedia.org/wiki/IBM_Cloud#SoftLayer) 👀
- [100B messages sent per day end of 2020](https://techcrunch.com/2020/10/29/whatsapp-is-now-delivering-roughly-100-billion-messages-a-day/)
- [2014 High Scalability numbers](http://highscalability.com/blog/2014/3/31/how-whatsapp-grew-to-nearly-500-million-users-11000-cores-an.html)

## Discuss issues with replication lag

With strong consistency there's no problem, but with eventual consistency, what happens if the two people chatting are getting served by a primary and a replica, and the replica has a little lag?

Discuss "reading your own writes", "monotonic reads", "consistent prefix reads", "read & write quorum".

## What does Facebook Messenger use?

- In 2018 they migrated [from HBase to MyRocks](https://engineering.fb.com/2018/06/26/core-data/migrating-messenger-storage-to-optimize-performance/) and from spinning disks to flash storage.
- [MyRocks](https://engineering.fb.com/2018/06/26/core-data/migrating-messenger-storage-to-optimize-performance/) is a MySQL storage engine that replaces InnoDB, using RocksDB (embeddable persistent key value store forked off Google's LevelDB). RocksDB improves on InnoDB's write amplification and storage footprint issues. 
- Note that RocksDB by itself has no replication and is key-value (i.e. no SQL layer), so MyRocks adds this.

![Messenger - Topology](/iknowkungfoo/whatsapp/messenger-topology.png)

[MySQL for Messaging - @Scale 2014](https://www.youtube.com/watch?v=eADBCKKf8PA) (Facebook Engineer)

[Same, but blogpost from FB Engineering](https://engineering.fb.com/2014/10/09/production-engineering/building-mobile-first-infrastructure-for-messenger/)

[Initial Erlang implementation of Facebook Chat - Slide deck](http://www.erlang-factory.com/upload/presentations/31/EugeneLetuchy-ErlangatFacebook.pdf)

## What does Slack use?

- Migrated from MySQL to Vitess.
- Stack: Kubernetes, Terraform, Consul, HAProxy, Flannel (in-house proactive caching), Solr (search), Memcached, PHP/Hack & Java/Kotlin.

## Video tutorials online

- [Exponent: System Design Mock Interview: Design Facebook Messenger](https://www.youtube.com/watch?v=uzeJb7ZjoQ4)
- [codeKarle: WhatsApp System Design | FB Messenger System Design | System Design Interview Question](https://www.youtube.com/watch?v=RjQjbJ2UJDg)
