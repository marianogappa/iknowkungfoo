---
title: "Instagram"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

![Instagram](/iknowkungfoo/instagram/instagram-design.png)

## Main differences

There is a lot of overlap with [Twitter design](/iknowkungfoo/systemdesignquestions/twitter):

- [Timeline design](/iknowkungfoo/systemdesignquestions/newsfeed)
- Social graph
- Search

But Instagram/Tiktok is image/video based, so:

- Bandwidth can be an issue in some localities, so on upload, producing videos at different bitrates is important. How about consistent loudness between videos?
- Access pattern (a lot in the first few days, then fade away) becomes important for cost saving. Need a "cached timelines" storage on Redis/Memcached, and a "archived timelines" on Cassandra/HBase.
- "The algorithm" is very important, so analytics feeding back to timelines MUST be mentioned.

## "The algorithm"

- Users need to be profiled thoroughly to best recommend them posts that will boost engagement & retention.
- Their social graph & activity must be analyzed through collaborative filtering & other data science algorithms, and output some profiling artifact (e.g. tags) to be kept in some User Service's cache.
- If elaborating on "The algorithm": need `User Service`, `User cache`, `Trends service`, `Profiling Service`, etc.
- Then, use the feedback `Post-process service` in the diagram to explain how posts are sent to users timelines. 

## Video

- [codeKarle - Facebook System Design | Instagram System Design | System Design Interview Question](https://www.youtube.com/watch?v=9-hjBGxuiEs)
