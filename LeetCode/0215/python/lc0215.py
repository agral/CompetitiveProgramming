from typing import List
import heapq

# Runtime complexity: O(n * logk)
# Auxiliary space complexity: O(k)
# Subjective level: easy+ (a known problem, easy when using standard minHeap is allowed)
# Solved on: 2026-02-28
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        minHeap = []
        for num in nums:
            heapq.heappush(minHeap, num)
            if len(minHeap) > k:
                heapq.heappop(minHeap)
        return minHeap[0]
