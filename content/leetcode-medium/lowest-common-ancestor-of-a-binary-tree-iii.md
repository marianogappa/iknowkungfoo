---
title: Lowest Common Ancestor Of A Binary Tree Iii
date: 2022-11-20T09:03:20-08:00
---

(https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii)



## Algorithm

```python
class Node:
    def _init_(self, val):
        self.val = val
        self.left = None
        self.right = None
        self.parent = None

class Solution:
    def lowestCommonAncestor(self, p: 'Node', q: 'Node') -> 'Node':
        visited = {p.val, q.val}
        while True:
            if p.parent:
                p = p.parent 
                if p.val in visited:
                    return p
                visited.add(p.val)
            if q.parent:
                q = q.parent
                if q.val in visited:
                    return q
                visited.add(q.val)

```


