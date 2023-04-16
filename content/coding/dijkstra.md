---
title: "Dijkstra"
date: 2022-11-20T09:03:20-08:00
draft: false
---
## Intuition

How do you find the shortest path from your home to your work in the city map? (assuming distance as only factor)

Dijkstra calculates this, but requires an "adjacency list" (check below section). 

## Algorithm

```python
# Time: O((V+E)*logV)
# Space: O(V)
def dijkstra(
    adjacency: dict[int, list[list[int]]],
    num_vertices: int, # != len(adjacency) if there are orphan vertices!!
    start_vertex: int
) -> list[int]:
    # We start with default assumption that distance from start_vertex to every vertex
    # is infinite, except for start_vertex.
    distances = [float('inf')] * (num_vertices+1) # +1 optional; depends on zero-indexed
    distances[start_vertex] = 0
    
    # We only want to visit each vertex once.
    visited: set[int] = set()

    # Use a heap to iteratively pop the shortest distance vertex, 
    # until all vertices are visited
    h = []
    heapq.heappush(h, (0, start_vertex))

    # O(v): While there are still unvisited vertices...
    while len(h) > 0:
        # O(log v): Pop the next shortest distance vertex
        cur_dist, vertex = heapq.heappop(h)
        
        # O(1): Get the vertex's edges and mark the vertex as visited
        edges = adjacency[vertex]
        visited.add(vertex)
        
        # O(e): For every edge...
        for edge in edges:
            [dest_vertex, distance] = edge

            # If the vertex is visited, or we already found a shortest distance, 
            # ignore the edge
            if dest_vertex in visited or distances[dest_vertex] <= cur_dist + distance:
                continue

            # O(log v): Store the shortest distance and add it to the heap, 
            # so that we visit it later
            distances[dest_vertex] = cur_dist + distance
            heapq.heappush(h, (distances[dest_vertex], dest_vertex))
        
    return distances
```

### 🤔 adjacency list? wtf? 

It contains a list of "Going from Vertex A to Vertex B requires X cost". In the code example it's a `dict[FromVertex, list[tuple[ToVertex, Cost]]]`, but the `tuple` is simplified as a `list`. It can be structured differently: could be a matrix.

### 🧠 Pro tips!

- When exercise involves a graph, or a shortest path, or a minimum cost, always think of Dijkstra!
- Note that it calculates minimum distance from `start_vertex` to ALL vertices!
- To think about TC: `E*logV` + `V*logV` because we put ~all edges and ~all vertices on the heap!

## Done 🎉🎉🎉

## Appendix: Other tutorials online

[Leetcode Solutions](https://leetcode.com/problems/network-delay-time/): Dijkstra pattern for many exercises

[Dijkstra’s Algorithm | LeetCode The Hard Way](https://leetcodethehardway.com/tutorials/graph-theory/dijkstra): This tutorial explains the overview, implementation, and suggested problems of Dijkstra’s algorithm with C++ code snippets.

[A guide to Dijkstra’s Algorithm - LeetCode Discuss](https://leetcode.com/discuss/general-discussion/1059477/A-guide-to-Dijkstra's-Algorithm): This tutorial provides a detailed explanation and intuition behind Dijkstra’s algorithm with diagrams and examples.

## Appendix: Related Leetcodes

743. [Network Delay Time](https://leetcode.com/problems/network-delay-time): This problem asks you to find the time it takes for a signal to reach all nodes in a network given a list of edges with delays. You can use Dijkstra’s algorithm to find the shortest time from the source node to each node and return the maximum time among them.

1514. [Path with Maximum Probability](https://leetcode.com/problems/path-with-maximum-probability/): This problem asks you to find the path with the maximum probability of success between two nodes in a graph given a list of edges with probabilities. You can use Dijkstra’s algorithm to find the path with the maximum product of probabilities from the source node to each node and return the product for the destination node.

787. [Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/): This problem asks you to find the cheapest price from a source city to a destination city within a given number of stops using flights. You can use Dijkstra’s algorithm to find the minimum cost from the source city to each city and keep track of the number of stops along the way.
