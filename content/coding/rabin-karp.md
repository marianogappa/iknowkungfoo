---
title: "Rabin-Karp"
date: 2022-11-20T09:03:20-08:00
draft: false
---
## Intuition

Substring search (e.g. is "log" in "catalog"?) is `O(t*p)`:

```python
def is_substr(text: str, pattern: str) -> bool:
    return any([text[i:i+len(pattern)] == pattern for i in range(len(text)-len(pattern)+1)])
```

Can we do it faster?

## Algorithm idea

- Start by doing the same, but instead of `text[] == pattern`, do `hash(text[]) == hash(pattern)`. Which `hash()`? 🤔
- Note that `hash(pattern)` never changes, so only compute it once: `O(p)` so far.
- First time, compare with `hash(text[:len(pattern)])`, again `O(p)`.
- 🤯: instead of a new `O(p)` hash as we slide, *update* the hash in `O(1)` by removing the evicting character and adding the incoming character to the hash.
- Hash collisions exist, so one still has to check equality to confirm (another `O(p)`).
- 💥 BOOM 💥! Now it's `O(p + t)` on average!

## Algorithm

```python

def compute_hash(text, pattern):
    hash = 0
    for i in range(len(pattern)):
        hash = (hash * 256 + ord(text[i])) % 101
    return hash

def update_hash(hash, old_char, new_char, pattern):
    h = pow(256, len(pattern)-1) # highest power of base
    hash -= (ord(old_char) * h) % 101 # remove char
    hash = (hash * 256 + ord(new_char)) % 101 # 👀 same as compute_hash
    return hash

def rabin_karp_search(text, pattern):
    p_hash = compute_hash(pattern, pattern)
    t_hash = compute_hash(text, pattern)

    for i in range(len(text)-len(pattern)+1):
        if p_hash == t_hash and pattern == text[i:i+len(pattern)]:
            return i

        if i < len(text)-len(pattern): # As in "Every time except last"
            t_hash = update_hash(t_hash, text[i], text[i+len(pattern)], pattern)

    return -1
```


## 🤔 Understanding the hash functions

### `compute_hash`

It's tricky to see in reverse, but the hash is just a weighted sum (the % is just for convenience: positive & bounded):

```python
hash("HELLO") = ( H *256^4 + E *256^3 + L *256^2 + L *256^1 + O *256^0 ) % 101
```

### `update_hash`

- Adding the new char works the same as in `compute_hash`
- To remove `H` in the example above we'd need to `-= H*256^4`, but how do we figure out the `4` in `update_hash`?
- It will always be `len(pattern-1)`! Because it goes from `0` to `n` right to left, and the size is `len(pattern)`.

## 🧠 Pro tips

- 👀 worst case still `O(t*p)` if all hashes collide and have to check equality every time.
- 256 is the "base" and 101 is the "mod". Chosen because there are 256 characters and 101 is a large prime, which reduces chance of collision, but other numbers can be used.
- `%` is distributive over sum, so you can `% 101` after every operation or at the end and it should return the same.

## Done 🎉🎉🎉

## Appendix: Other tutorials online

DON'T LOOK 🙈. They ALL SUCK 😱.

## Appendix: Related Leetcodes

[28. Implement strStr()](https://leetcode.com/problems/implement-strstr/): This problem asks you to return the index of the first occurrence of a given needle in a given haystack, or -1 if the needle is not part of the haystack.

[49. Group Anagrams](https://leetcode.com/problems/group-anagrams/): This problem asks you to group all the strings in a given array that are anagrams of each other, using a hash function to encode each string.

[1044. Longest Duplicate Substring](https://leetcode.com/problems/longest-duplicate-substring/): This problem asks you to find the longest substring of a given string that occurs at least twice, using binary search and Rabin Karp to check for duplicates.

[187. Repeated DNA Sequences](https://leetcode.com/problems/repeated-dna-sequences/): This problem asks you to find all the 10-letter-long sequences that occur more than once in a given DNA molecule, using Rabin Karp to hash each sequence.

[686. Repeated String Match](https://leetcode.com/problems/repeated-string-match/): This problem asks you to find the minimum number of times you need to repeat a given string A such that another string B is a substring of it, using Rabin Karp to check for substring match.
