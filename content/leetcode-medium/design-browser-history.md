---
title: Design Browser History
date: 2022-11-20T09:03:20-08:00
---

https://leetcode.com/problems/design-browser-history



## Algorithm

```python
class BrowserHistory:

    def __init__(self, homepage: str):
        self.cursor = 0
        self.history = [homepage]
        

    def visit(self, url: str) -> None:
        self.history = self.history[:self.cursor+1]
        self.history.append(url)
        self.cursor = len(self.history)-1

    def back(self, steps: int) -> str:
        self.cursor -= min(steps, self.cursor)
        return self.history[self.cursor]

    def forward(self, steps: int) -> str:
        self.cursor = min(len(self.history)-1, self.cursor+steps)
        return self.history[self.cursor]

        


```

Your BrowserHistory object will be instantiated and called as such:
obj = BrowserHistory(homepage)
obj.visit(url)
param_2 = obj.back(steps)
param_3 = obj.forward(steps)

