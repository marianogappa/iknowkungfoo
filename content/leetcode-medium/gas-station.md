---
title: Gas Station
date: 2022-11-20T09:03:20-08:00
---

(https://leetcode.com/problems/gas-station)



## Algorithm

```python
class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        def change(i: int) -> int:
            return gas[i%len(gas)] - cost[i%len(gas)]
        left = 0
        tank = 0
        right = 0
        while right < 2 * len(gas):
            tank += change(right)
            
            if right-left + 1 == len(gas) and tank >= 0:
                return left    
            while tank < 0:
                tank -= change(left)
                left += 1
            right += 1
        return -1

```


