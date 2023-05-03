---
title: Top K Frequent Words
date: 2022-11-20T09:03:20-08:00
---

(https://leetcode.com/problems/top-k-frequent-words)



## Algorithm

```python
from collections import defaultdict
import heapq
from typing import List

# Time: O(w*logk)
# Space: O(w)
class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        freqs = defaultdict(int)
        for word in words:
            freqs[word] += 1
        
        h = []
        for word, freq in freqs.items():
            heapq.heappush(h, ((-freq, word), word))
        
        result = []
        for _ in range(k):
            _, word = heapq.heappop(h)
            result.append(word)
        
        return result

```


