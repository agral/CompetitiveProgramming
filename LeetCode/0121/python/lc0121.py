from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-04-16
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        minSoFar = prices[0]
        profit = 0
        for i in range(1, len(prices)):
            minSoFar = min(minSoFar, prices[i])
            profit = max(profit, prices[i] - minSoFar)
        return profit
