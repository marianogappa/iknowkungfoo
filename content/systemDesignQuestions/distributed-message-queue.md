---
title: "Distributed Message Queue"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

![Distributed Message Queue](/iknowkungfoo/distributed_message_queue/distributed_message_queue.png)

# Delivery semantics

- Exactly once, at most once or at least once? Why most systems do at least once?
- Is it a queue or a topic? In a queue, once a message is sent to a consumer, it is deleted or at least not sent to others. In a topic, different consumer groups can receive all messages, and deletion is subject to a retention policy.

# Failure modes: when is a message actually "consumed"?
- What if the message was received but the consumer crashed right after?
- What if sending the message to the consumer times out? Was it actually "processed" or not? Maybe the network glitched but the consumer did "ack" it.
- If we retry on failure, is it ok to process other messages to maintain throughput or is it more important to maintain strict ordering?
- SQS makes a message "invisible" while it is being processed by a consumer, but this invisibility times out if the message is not acked (or failed) before a deadline and then it can be retried by a different consumer.
- Common approaches: consumers "ack" messages, and then producer can consider them "done", and consumer keeps state of consumed messages by some id, so consumption can be idempotent.

# Concurrency pattern for scaling consumers

- If strict ordering is required for consumption, then there's no way to scale reads horizontally, but typically absolute ordering is not required.
- More commonly, there are many queues and partitions within each queue, and only ordering within partition needs to support ordering.
- In order to keep track of which backend instance contains which queue contents, discuss consistent hashing.

# Where to store messages?

- High performance requires storing messages in memory, but memory is volatile so replication is needed (high reliability).
- Durability may be a requirement, cost or scale may mean messages don't fit in memory or are too expensive, so disk may be required.
- The above suggests LSM trees which is exactly what Kafka uses.

# Video tutorials

- https://www.youtube.com/watch?v=iJLL-KPqBpM&ab_channel=SystemDesignInterview
