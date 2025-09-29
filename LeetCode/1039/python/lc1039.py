from typing import List
import math

# Runtime complexity: O(n^2)
# Auxiliary space complexity: O(n^2)
# Subjective level: medium/hard
class Solution:
    def minScoreTriangulation(self, values: List[int]) -> int:
        LEN_VALUES = len(values)
        dp = [[0] * LEN_VALUES for _ in range(LEN_VALUES)]

        for y in range(2, LEN_VALUES):
            for x in range(y-2, -1, -1):
                dp[y][x] = math.inf
                for z in range(x+1, y):
                    dp[y][x] = min(dp[y][x], dp[z][x] + dp[y][z] + (values[x]*values[y]*values[z]))
        return dp[LEN_VALUES-1][0]
