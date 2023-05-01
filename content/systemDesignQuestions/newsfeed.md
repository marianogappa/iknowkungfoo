---
title: "Newsfeed"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

## Applies to

- Twitter
- Facebook
- Instagram
- Tiktok

![Twitter Feed](/iknowkungfoo/newsfeed/twitter-feed.png)

## Content amount and structure

- ~100:1 read/write ratio. Should be optimised for reads. 
- Computing newsfeed on read? 🤔 ~75 followers * 100 last tweets multi-get, sort & filter 😱😱 TOO SLOW! Can't do it in real-time.
- Newsfeed must be cached for all "Active users". Avg_tweet_sec * Avg_followers = 😱! Tricky but necessary; it can be done async. For users not active in last 30 days, compute at read-time.

## Fan out service

- On write, fan out the `(tweet_id, user_id, ...metadata)` to all followers' timelines. Timelines must be kept in cache for performance, so make each entry as small as possible.
- Data structure for one cached timeline? Should stay sorted by timestamp and evict after certain size, so heap makes sense. If inserts into it won't be out of order then a sorted list doesn't require a heap, and with a map it could allow deletes in `O(1)`.
- On read, a multi-get will be needed to fetch all tweets. At Twitter that reader service caches all tweets for the last month. Here, the question on how to shard tweets is important.

## Social graph service

- Need to keep knowledge of who's following who, to fan out tweets.
- Graph databases seem optimal? But no mature ones, everyone uses RDBMS 🤷‍♂️.

## How to shard Tweets? Explore options!

- By Timestamp? 🤔 NOO! Most requested tweets will be in the recent shards, making a few shards ultra hot and all others unused!
- By TweetID? Maximises data spread (& disk usage!), minimises hot shards, but most multi-get requests could involve all shards (scatter-gather), so slower shard determines p99 latency! Note that there's a "User Timeline" request with all tweets of one user.
- By UserID? User timeline requests involve a single shard, but popular users create hot shards, some users post more than others making disk usage unevenly spread.
- When discussing sharding, mention Consistent Hashing. Ask if you should elaborate or keep going.

## Blob storage!

- Don't forget that feed has image/video too. Object storage multi-get by id takes care of it, but important to mention CDN (data-locality!). Blob requests happen in parallel to text requests and can be asynchronously loaded as they arrive.

## Must see video

[Rafi Krikorian (Twitter) @ InfoQ](https://www.infoq.com/presentations/Twitter-Timeline-Scalability/)
