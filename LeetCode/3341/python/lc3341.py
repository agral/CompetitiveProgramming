from typing import List
from typing import Tuple
import heapq
import math

# Runtime complexity: O(XY*log(XY)), where X, Y are the appropriate dimensions of moveTime array.
# Auxiliary space complexity: O(XY)
class Solution:
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        Y, X = len(moveTime) - 1, len(moveTime[0]) - 1
        return self._shortestPath(moveTime, (0, 0), (Y, X))

    def _shortestPath(self, moveTime: List[List[int]],
                      source: Tuple[int, int], destination: Tuple[int, int]) -> int:
        DIRS = ((1, 0), (0, 1), (-1, 0), (0, -1))
        Y, X = len(moveTime), len(moveTime[0])
        distance = [[math.inf] * X for _ in range(Y)]
        distance[0][0] = 0

        minHeap = [(0, source)]
        while minHeap:
            d, location = heapq.heappop(minHeap)
            if location == destination:
                return d
            y, x = location
            if d > distance[y][x]:
                continue
            for dy, dx in DIRS:
                yy, xx = y + dy, x + dx
                if yy < 0 or yy >= Y or xx < 0 or xx >= X:
                    continue
                updatedDist = max(moveTime[yy][xx], d) + 1
                if updatedDist < distance[yy][xx]:
                    distance[yy][xx] = updatedDist
                    heapq.heappush(minHeap, (updatedDist, (yy, xx)))

        raise RuntimeError("Exit has not been found!")
