---
title: Nested List Weight Sum
date: 2022-11-20T09:03:20-08:00
---

(https://leetcode.com/problems/nested-list-weight-sum)

Time: O(n)
Space: O(d) depth

## Algorithm

```python
class Solution:
    def depthSum(self, nestedList: List[NestedInteger]) -> int:
        return sum([ flatten(item, 1) for item in nestedList ])

def flatten(item: NestedInteger, depth: int) -> int:
    if item.isInteger():
        return item.getInteger() * depth
    return sum([flatten(nested_item, depth+1) for nested_item in item.getList()])

```


