---
title: Binary Tree Right Side View
date: 2022-11-20T09:03:20-08:00
---

(https://leetcode.com/problems/binary-tree-right-side-view)

Veeery trivial. Just dfs in-order checking right side first.


## Algorithm

```python
# Time: O(n)
# Space: O(n)
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        return dfs(root, [], 0)
        

def dfs(root: Optional[TreeNode], right_side: list[int], level: int) -> list[int]:
    if not root:
        return right_side
    
    if len(right_side) == level:
        right_side.append(root.val)
        
    dfs(root.right, right_side, level+1)
    dfs(root.left, right_side, level+1)
    
    return right_side

```


