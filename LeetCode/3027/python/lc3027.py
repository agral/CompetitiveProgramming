import math
from typing import List

# Runtime complexity: O(n^2), optimal.
# Auxiliary space complexity: O(sort)
# Subjective level: medium. The answer is the same as 3025, with subjective level also medium.
class Solution:
    def numberOfPairs(self, points: List[List[int]]) -> int:
        LEN_POINTS = len(points)
        ans = 0
        points.sort(key=lambda arg: (arg[0], -arg[1]))
        for i, (_, yi) in enumerate(points):
            maxY = -math.inf
            for j in range(i+1, LEN_POINTS):
                (_, yj) = points[j]
                if yi >= yj and yj > maxY:
                    maxY = yj
                    ans += 1
        return ans
