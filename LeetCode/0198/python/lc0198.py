from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium-
# Solved on: 2026-04-10
class Solution:
    def rob(self, houses: List[int]) -> int:
        NUM_HOUSES = len(houses)

        # update: let's just handle the corner case with an if.
        if NUM_HOUSES == 1:
            return houses[0]

        # dp[i] holds the maximum possible gain from breaking into houses [0..i).
        # This extra zeroth entry is necessary only to make this algo work around
        # the case of only one house being present on the street.
        # otherwise we could go with dp[i] == gain(houses[0..i]), and len(dp) == len(houses).
        dp = [0] * NUM_HOUSES
        dp[0] = houses[0]
        dp[1] = max(houses[0], houses[1])
        for i in range(2, NUM_HOUSES):
            dp[i] = max(dp[i-2] + houses[i], dp[i-1])

        return dp[-1]
