from typing import List
import heapq

# Runtime complexity: O(n*logn)
# Auxiliary space complexity: O(n)
# Subjective level: medium+
# Solved on: 2026-02-16
class Solution:
    def eatenApples(self, apples: List[int], days: List[int]) -> int:
        ans = 0
        L = len(apples) # == len(days).
        minHeap = []

        day = 0
        while day < L or minHeap:
            # Remove all the rotten apples:
            while minHeap and minHeap[0][0] <= day:
                heapq.heappop(minHeap)

            # Add the apples that grew today - and will expire at day + days[day]:
            if day < L and apples[day] > 0:
                heapq.heappush(minHeap, (day + days[day], apples[day]))

            # Eat one apple, if currently available:
            if minHeap:
                rotsOn, numApples = heapq.heappop(minHeap)
                ans += 1
                if numApples >= 2:
                    heapq.heappush(minHeap, (rotsOn, numApples-1))

            # Advance one day:
            day += 1
        return ans
