---
title: "Trie"
date: 2022-11-20T09:03:20-08:00
draft: false
---
## Intuition

How would you implement an autocomplete feature? On every letter typed, check all words on a list that start with that prefix? 😱

Ideally organise the words by prefix, so all words with any prefix are clustered. Of course, many prefixes overlap, so we'll need a tree structure, branching on every letter of the prefix.

## Algorithm idea

- A Trie is a tree, where the actual data is in the branches rather than the node, via `dict[str, TrieNode]`, where `str` is the data.
- It starts with a sentinel root node, and words are inserted by inserting a branch per letter, or reusing it if it exists.
- Benefits: checking if a word exists in a Trie is O(w), and collecting all words with a given prefix is O(m*n) (m = longest word, n = words in Trie) but on average both m & n will be smaller.

## Algorithm

```python
class TrieNode:
    def __init__(self):
        self.is_end = False # True if this node is the end of a word
        self.children: dict[str, TrieNode] = defaultdict(TrieNode)

    # Time: O(w)
    def insert(self, word: str) -> None:
        if word:
            self.children[word[0]].insert(word[1:])
        else
            self.is_end = True

    # Time: O(w)
    def exists(self, word: str) -> bool:
        if not word:
            return self.is_end
        if word[0] not in self.children:
            return False
        return self.children[word[0]].exists(word[1:])
    
    # Note: same as exists, but collect words instead of return True!
    def search(self, prefix: str) -> list[str]:
        if not prefix:
            return self.collect_words("")
        if prefix[0] not in self.children:
            return []
        return self.children[prefix[0]].search(prefix[1:])

    # collect all words from this node onwards!
    def collect_words(self, prefix):
        words = []
        if self.is_end:
            words.append(prefix)
        for char in self.children:
            words.extend(self.children[char].collect_words(prefix + char))
        return words
```

## Done 🎉🎉🎉

## Appendix: Other tutorials online

[Beginner-friendly guide to Trie](https://leetcode.com/discuss/study-guide/931977/Beginner-friendly-guide-to-Trie-Tutorial-%2B-Practice-Problems): is a discussion post that explains what trie is, how to write one, and how it can be used, with some practice problems and solutions.

[GeeksForGeeks](https://www.geeksforgeeks.org/introduction-to-trie-data-structure-and-algorithm-tutorials/): lots of visuals, many languages and large descriptions.

## Appendix: Related Leetcodes

[211. Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/): This problem asks you to design a data structure that supports adding and searching words with wildcards.

[212. Word Search II](https://leetcode.com/problems/word-search-ii/): This problem asks you to find all the words in a given board that exist in a given word list, using a Trie to optimize the search.
