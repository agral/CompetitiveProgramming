from typing import List
import heapq

# Runtime complexity: O(k * log(k)) generally; where k is less than or equal len(nums1)*len(nums2);
#                     usually a lot less than that product.
# Auxiliary space complexity: O( max(len(nums1), k) ) (for keeping the min-heap, and for the answer (k)).
# Subjective level: easy+
# Solved on: 2026-02-10
class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        # 1. Create a min-heap of all possible pairs formed from entries of 1st list paired with nums2[0]:
        minHeap = [] # minHeap holds tuples of (sum, i, j), such that nums1[i]+nums2[j] = sum.

        for i in range(min(len(nums1), k)):
            heapq.heappush(minHeap, (nums1[i] + nums2[0], i, 0))

        # 2. Keep constructing the answer until it is of length `k`.
        #    It is guaranteed that the first pair is `nums1[0]`, `nums2[0]`;
        #    the first entry on the `minHeap`. The other entries are added
        #    to the `minHeap` as previous are removed, if possible.
        #    Therefore the `minHeap` has at all times at most `len(nums1)` entries.
        #    Also worth noting is the fact that `i` gets increased "naturally", as soon as
        #    it is used with `nums2[0]`, which is guaranteed to be the least of all `nums2`.
        ans = []
        while len(ans) < k and minHeap:
            _, i, j = heapq.heappop(minHeap) # sum (the first entry of the tuple) is unused here.
            ans.append([nums1[i], nums2[j]])
            if j+1 < len(nums2):
                heapq.heappush(minHeap, (nums1[i] + nums2[j+1], i, j+1))

        return ans
