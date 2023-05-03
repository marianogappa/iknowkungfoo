---
title: Find Largest Value In Each Tree Row
date: 2022-11-20T09:03:20-08:00
---

(https://leetcode.com/problems/find-largest-value-in-each-tree-row)

Straight up BFS, no tricks.


## Algorithm

```python
# Time: O(n)
# Space: O(n) one row at a time, plus O(h) for solution space
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def largestValues(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []

        queue = deque()
        queue.append(root)
        result = []
        while queue:
            max_val = float("-inf")
            for _ in range(len(queue)):
                elem = queue.popleft()
                max_val = max(max_val, elem.val)
                
                if elem.left:
                    queue.append(elem.left)

                if elem.right:
                    queue.append(elem.right)

            result.append(max_val)

        return result

```


