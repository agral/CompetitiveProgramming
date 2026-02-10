from typing import List
import heapq

# Runtime complexity: O(n * log(n))
# Auxiliary space complexity: O(n)
# Subjective level: easy
# Solved on: 2026-02-10
class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        # 1. Transform stones from a List[int] into a max heap (standard heapq package):
        heapq.heapify_max(stones)

        # 2. Keep retrieving a pair of stones as long as there is more than one on the heap.
        while len(stones) >= 2:
            y = heapq.heappop_max(stones)
            x = heapq.heappop_max(stones)
            if y >= x:
                heapq.heappush_max(stones, y-x)

        # 3. The answer is the weight of the remaining stone, or 0 if there are no stones.
        if len(stones) == 0:
            return 0
        return stones[0]
