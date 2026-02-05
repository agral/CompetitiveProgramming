from typing import List
import heapq

# Runtime complexity: O(n*log(k))
# Auxiliary space complexity: O(k) (for keeping a min-heap)
# Subjective level: easy+
# Solved on: 2026-02-05
class KthLargest:
    def __init__(self, k: int, nums: List[int]) -> None:
        self.minHeap = []
        self.capacity = k
        for num in nums:
            self.heapify(num)

    def add(self, val: int) -> int:
        self.heapify(val)
        return self.minHeap[0]

    def heapify(self, num: int) -> None:
        heapq.heappush(self.minHeap, num)
        if len(self.minHeap) > self.capacity:
            heapq.heappop(self.minHeap)
