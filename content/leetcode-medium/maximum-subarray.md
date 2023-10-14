---
title: Maximum Subarray
date: 2022-11-20T09:03:20-08:00
---

https://leetcode.com/problems/maximum-subarray

Keep a running sum and keep track of local maximums at every iteration.

When the running sum becomes smaller than the current number, re-start
the running sum from this number.


## Algorithm

```python
class Solution:
    # Time: O(n)
    # Space: O(1)
    def maxSubArray(self, nums: List[int]) -> int:
        max_sum = cur_sum = float('-inf')

        for right in range(len(nums)):
            # If adding the current number to the running sum results
            # in a smaller number than itself, we might as well start
            # from scratch from it.
            if cur_sum + nums[right] < nums[right]:
                cur_sum = nums[right]
            else:
                # Otherwise keep the running sum.
                cur_sum += nums[right]

            # Keep track of local maxes found.
            max_sum = max(max_sum, cur_sum)
        
        return max_sum
```


