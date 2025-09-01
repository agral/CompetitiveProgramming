import heapq
from typing import List

# Runtime complexity: O((len(classes)+extraStudents)*log(len(classes)))
# Auxiliary space complexity: O(len(classes))
# Subjective level: medium; heapq package makes it an "easy" medium.
class Solution:
    def maxAverageRatio(self, classes: List[List[int]], extraStudents: int) -> float:
        def getIncreasedRatio(passing: int, total: int) -> float:
            """Returns the difference in pass ratio after one brilliant student joins the class."""
            return (passing + 1) / (total + 1) - (passing / total)

        heap = [(-getIncreasedRatio(passing, total), passing, total) for passing, total in classes]
        heapq.heapify(heap)

        for _ in range(extraStudents):
            _, passing, total = heapq.heappop(heap)
            heapq.heappush(heap, (-getIncreasedRatio(passing+1, total+1), passing+1, total+1))

        s = sum(passing / total for _, passing, total in heap)
        return s / len(heap)
