---
title: "Quickselect"
date: 2022-11-20T09:03:20-08:00
draft: false
---
## Intuition

How do you get the kth smallest element of an array?

- Quicksort the array, then it's the kth! But it's `O(n*logn)`.
- Heapify the array, then heappop k times! But it's `O(k*logn)`.

Is there a faster way?

## Algorithm idea

- Start with the quicksort idea: pick pivot and partition.
- If the pivot ends at the kth position, done! Pivot is the one.
- Otherwise, do the same process but only on the left side if pivot index > k, or the right otherwise.

```python
def partition(ns: list[int], left: int, right: int) -> int:
    pivot = randint(left, right)

    ns[pivot], ns[right] = ns[right], ns[pivot]

    for i in range(left, right):
        if ns[i] <= ns[right]:
            ns[i], ns[left] = ns[left], ns[i]
            left += 1

    ns[left], ns[right] = ns[right], ns[left]

    return left

# Time: O(n) average, O(n^2) worst case with very bad pivots
# Space: O(1) in-place
def _quickselect(ns: list[int], left: int, right: int, k: int) -> list[int]:
    if left == right:
        return ns[left]

    pivot_idx = partition(ns, left, right)
    if pivot_idx == k:
        return ns[pivot_idx]
    elif pivot_idx < k:
        return _quickselect(ns, pivot_idx+1, right, k)
    # pivot_idx > k
    return _quickselect(ns, left, pivot_idx-1, k)

def quickselect(ns: list[int], k: int) -> int: # k is 1-indexed here!
    return _quickselect(ns, 0, len(ns)-1, k-1) # make k 0-indexed
```

## Done 🎉🎉🎉

## Appendix: Other tutorials online

[Quickselect Algorithm - GeeksforGeeks](https://www.geeksforgeeks.org/quickselect-algorithm/): This tutorial introduces the Quickselect algorithm and gives some examples of its usage. It also compares it with other selection algorithms and provides a C++ implementation.

[Complexity Analysis of QuickSelect | Baeldung on Computer Science](https://www.baeldung.com/cs/quickselect): This tutorial explains the worst-case, the best-case, and the average-case time complexity of QuickSelect. It also shows how to implement QuickSelect using Lomuto partitioning with random pivot selection.

[QuickSelect Algorithm Understanding - Stack Overflow](https://stackoverflow.com/questions/10846482/quickselect-algorithm-understanding): This tutorial answers a question about how QuickSelect works and provides a Python implementation. It also explains the difference between QuickSelect and QuickSort and gives some tips on how to choose a good pivot element.

## Appendix: Related Leetcodes

[215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array): This problem asks you to find the k-th largest element in an array of integers. You can use quickselect to find the element in O(n) time on average.

[324. Wiggle Sort II](https://leetcode.com/problems/wiggle-sort-ii): This problem asks you to reorder an array of integers such that nums[0] < nums[1] > nums[2] < nums[3] … You can use quickselect to find the median of the array in O(n) time on average, and then use a three-way partitioning scheme to rearrange the elements around the median.

[347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements): This problem asks you to find the k most frequent elements in an array of integers. You can use a hash map to count the frequencies of each element, and then use quickselect to find the k-th largest frequency in O(n) time on average.

[973. K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin): This problem asks you to find the k closest points to the origin (0, 0) in a list of points on a plane. You can use quickselect to find the k-th smallest distance from the origin in O(n) time on average, and then return all points with distances less than or equal to that.

[1738. Find Kth Largest XOR Coordinate Value](https://leetcode.com/problems/find-kth-largest-xor-coordinate-value): This problem asks you to find the k-th largest value of a matrix where each cell is the XOR of all cells from (0, 0) to (i, j). You can use dynamic programming to compute the matrix values in O(mn) time, where m and n are the dimensions of the matrix, and then use quickselect to find the k-th largest value in O(mn) time on average.

[1985. Find the Kth Largest Integer in the Array](https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array): This problem asks you to find the k-th largest integer in an array of strings representing integers. You can use quickselect to find the element in O(n) time on average, using a custom comparator function to compare strings as integers.

[2343. Query Kth Smallest Trimmed Number](https://leetcode.com/problems/query-kth-smallest-trimmed-number): This problem asks you to answer queries that ask for the k-th smallest number after removing some numbers from both ends of a sorted array of integers. You can use binary search and quickselect to answer each query in O(logn + k) time on average, where n is the length of the array.
