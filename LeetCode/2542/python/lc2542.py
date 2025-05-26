from typing import List
import heapq

# Runtime complexity: O(n*logn)
# Auxiliary space complexity: O(n)
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        ans = 0
        sum = 0
        minHeap = []
        combined = sorted([(n2, n1) for n1, n2 in zip(nums1, nums2)], reverse=True)

        for (n2, n1) in combined:
            heapq.heappush(minHeap, n1)
            sum += n1
            if len(minHeap) > k:
                sum -= heapq.heappop(minHeap)
            if len(minHeap) == k:
                ans = max(ans, n2 * sum)
        return ans
