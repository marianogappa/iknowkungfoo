---
title: "Youtube"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

## Special functional requirements

- Support all kinds of devices

(this translates to supporting different formats, dimensions, bandwidths, audio bitrates)

## Special non-functional requirements

- Minimise buffering
- Maximise user session time

(these translate to low latency, high availability & "the algorithm")

## Video processing service responsibilities

- **File chunker**: allow parallelism for processing pipeline, and for deduping.
- **Moderation**: filter out illegal content ASAP.
- **Tagging**: tagging for "the algorithm", thumbnails, captions (could be used for search).
- **Transcoding, bitrate, aspect ratio & resolution**
- **Upload to CDN**
- **Publish event to Kafka**

(there is overlap between components, e.g. moderation requires classifier from tagging step, classifier returning identical video descriptions would hint that the video is exactly the same, which deduping could have missed).

## Asset service

- Thumbnails for videos and other metadata (stored & cached).
- Object storage links.
- CDN links for each version.

## S3 is a bottleneck to CDNs

- All CDNs that need a copy of chunks of videos need to request it from S3 😱.
- Rather than making S3 a bottleneck, decentralise with some "gossip protocol" topology, and get chunks from different CDNs.
- CDNs are local to timezones that have peak times and idle times. Using idle times for "rebalancing" reduces traffic.
- Netflix gives hardware to ISPs to cache CDN content so they don't even go to the public Internet to see it.

# Analytics service

- **Collaborative filtering** to figure out what to recommend to the user.
- **Traffic prediction** to co-locate videos with the users most likely to watch them.
- **Device & user fingerprinting** to prevent credential sharing, fraud, etc.
- **Machine learning model iteration** on content to improve classifiers, translators, video descriptions, tagging.
- **User profiling** based on video preference, length of video watched, likes, subscribes to improve "the algorithm".
- **Accounting** on videos liked, watched, subscribed, ads clicked for paying channels.

## Misc points to mention!

- Client owns "adaptive bitrate streaming" responsibility!
- Log stream stats (engagement with video) for "the algorithm"!

## Video tutorials

[Netflix System Design | YouTube System Design | System Design Interview Question](https://www.youtube.com/watch?v=lYoSd2WCJTo)
