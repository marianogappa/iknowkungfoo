---
title: "Union Find"
date: 2022-11-20T09:03:20-08:00
draft: false
---
## Intuition

Union-Find solves any problem with an undirected graph for which you need to answer one of these:

- How many "clustered vertices" (or disjoint sets, or "groups") are there?
- Do vertex A & vertex B belong to the same group?

Examples:

- [Number of islands](https://leetcode.com/problems/number-of-islands/) (or "Battleships"): Count connected '1’s in a binary grid.
- [Accounts Merge](https://leetcode.com/problems/accounts-merge): Merge accounts by email. Each account has a name and some emails.

## Algorithm idea

- Given a graph of vertices and edges, we think of it as a "forest": many isolated/disjoint trees of vertices, where the edges define parent/children relationships.
- Use `add` to add a vertex, which will start isolated as its own tree in the forest.
- Use `union` to add an edge, which will connect 2 vertices (and maybe recursively two trees). Largest tree's parent will become parent of the union.
- Use `find` to answer "which set/tree/group does a vertex belong to?".
- Use `set_count` to answer "how many disjoint sets are there?".

## Algorithm

```python
class UnionFind: # Used to keep track of disjointed sets on graphs.
    parent: list[int] # Sets are trees. Each vertex has a parent.
    size: list[int] # Optimization: on union, larger set becomes parent.
    idx: dict[any, int] # Optional: elements are mapped to int "labels".

    def __init__(self):
        self.parent = []
        self.size = []
        self.idx = {}

    # Add a new vertex to the forest, as a new tree of size 1
    def add(self, vertex: any):
        idx = len(self.parent) # Next available idx to use for this vertex.
        self.idx[vertex] = idx
        self.parent.append(idx) # New vertex is its own parent in own tree.
        self.size.append(1) # Tree of one vertex: size = 1

    # Union two vertices: if they belong to != sets, make larger parent
    # of the smaller. Keep track of sizes.
    # vertices must exist (via add)!
    def union(self, vertex1: any, vertex2: any):
        set1 = self.find(vertex1)
        set2 = self.find(vertex2)
        if set1 != set2:
            if self.size[set1] > self.size[set2]:
                self.size[set1] += self.size[set2]
                self.parent[set2] = set1
            else:
                self.size[set2] += self.size[set1]
                self.parent[set1] = set2

    def _find(self, idx: int) -> int:
        return self._find(self.parent[idx]) if self.parent[idx] != idx else idx
    
    # Find the set a vertex belongs to.
    # vertex must exist (via add)!
    def find(self, vertex: any) -> int:
        return self._find(self.idx[vertex])
    
    # Count how many different sets exist (O(n))
    def set_count(self) -> int:
        return sum(1 for idx, parent in enumerate(self.parent) if idx == parent)
```

### 🧠 Pro tips!

- There are many implementations, but learn the above: "Weighted Union Find w/ Path compression", because all 3 base methods are ~O(1).
- Why O(1)? `union` is `O(2*find)`, and `find` calls `find` recursively, but the depth of that tree will be better than `log(n)` because the path compression flattens the tree, so `ALMOST O(1)`.
- This version starts with empty "forest" and you add vertices, and there's a mapping from "element" to "vertex index". If you know the size beforehand and "elements" are already numbers 0 to n, you can bypass `add` & `find` and use `union` & `_find` directly.

## Done 🎉🎉🎉

## Appendix: Other tutorials online

[Swift Algorithm Club](https://aquarchitect.github.io/swift-algorithm-club/Union-Find/): good article explaining the same algorithm variant, but with some visuals and a longer explanation.

[Disjoint Set Union (DSU)/Union-Find - A Complete Guide - LeetCode Discuss](https://leetcode.com/discuss/general-discussion/1072418/Disjoint-Set-Union-(DSU)Union-Find-A-Complete-Guide): This post explains the concept and implementation of Union-Find with examples and code snippets. It also covers some variations and optimizations of Union-Find such as union by rank, path compression, path halving and path splitting.

[Union-Find - LeetCode Discuss](https://leetcode.com/problems/number-of-operations-to-make-network-connected/discuss/477806/python-union-find): This post shows a simple and concise Python implementation of Union-Find with comments. It also explains the time complexity analysis of Union-Find operations.

## Appendix: Related Leetcodes

684. [Redundant Connection](https://leetcode.com/problems/redundant-connection/): This problem asks you to find an edge that can be removed from a graph to make it a tree. You can use Union-Find to detect cycles in the graph and return the last edge that creates a cycle.

685. [Number of Provinces](https://leetcode.com/problems/number-of-provinces/): This problem asks you to find the number of connected components in a graph given an adjacency matrix. You can use Union-Find to group the nodes that are connected and return the number of distinct groups.
