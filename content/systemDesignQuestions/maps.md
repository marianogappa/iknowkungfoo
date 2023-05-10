---
title: "Maps"
date: 2022-11-20T09:03:20-08:00
draft: false
---

(not a guide for this question; only for how this question is different from all others)

## Special functional requirements

- Calculate best routes between two locations (w/ ETA, distance, warnings)
- Construct & maintain road map of the World
- Architecture must allow add-on features

## Calculating route between A & B

Straight up Dijsktra? 🤔 Doesn't work! Needs to process all V & E in the digraph! 😱

![Cell tiers](/iknowkungfoo/maps/maps-cell-tiers.png)

- Supported territories can be subdivided in "cells" to limit the size of the graph algorithms.
- Cells can be of multiple sizes or tiers.
- Calculate aerial distance, and cells of point A & B. Have different algorithms depending on aerial distance.
- If single-cell, Dijkstra works. Otherwise, decide which cell tier applies and rely on cached inter-cell routes to optimise.

## What if A & B are very close to edge of Cell?

- Neighbouring cells may have the optimal route, so they must be considered.
- But it's very unlikely a far-away cell is involved, so can decide threshold of involved cells based on aerial distance.

## Just distance?

Factors affecting ETA & preferred route:

- Distance
- Weather
- Traffic
- Accidents, Transport disruptions & Special events

But how to condense all of this into Dijkstra? All factors could be "pluggable" into an "Avg. Speed" weighting.

## Are you sure of Dijkstra?

- Bellman-Ford also exists.
- Dijkstra works well for a single trip, but Floyd-Warshall (n^3) might work better for "every V to every other V".

## Do A & B always align to the grid?

- Calculate distance from A to closest vertex, and from B to closest vertex, and add to ETA.

## Always use Dijkstra or cached Dijkstra?

As people travel, the system gets the actual time it took from A to B! Dijkstra is approximating this source of truth! So edge Avg. Speeds are calculated based on:

- Cached Floyd Warshalls
- Cached real trip durations
- Fresh Dijkstras

All caches get outdated, or are valid for a given day of week, hour of day or special holiday timeslot. Route service decides.

## How to maintain the road map?

- Users moving around hint at new roads or roads that don't exist anymore. Update map asynchronously in map-reduce batch jobs.
- Third-party map update, transport updates, weather, special events update must be ingested in real-time and update the map.
