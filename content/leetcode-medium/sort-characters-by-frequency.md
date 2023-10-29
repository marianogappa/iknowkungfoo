---
title: Sort Characters By Frequency
date: 2022-11-20T09:03:20-08:00
---

https://leetcode.com/problems/sort-characters-by-frequency



## Algorithm

```python
class Solution:
    def frequencySort(self, s: str) -> str:
        sorted_tpls = sorted([(-count, char) for char, count in dict(Counter(s)).items()])
        return ''.join(
            [char*(-count) for count, char in sorted_tpls]
        )
```


