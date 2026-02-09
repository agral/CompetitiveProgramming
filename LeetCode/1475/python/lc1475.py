from typing import List

# Runtime complexity: O(n^2)
# Auxiliary space complexity: O(n) (for storing the answer)
# Subjective level: easy
# Solved on: 2026-02-09
class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        ans = prices[:]
        for idx in range(len(prices)):
            for idxFollowing in range(idx+1, len(prices)):
                if prices[idxFollowing] <= prices[idx]:
                    ans[idx] -= prices[idxFollowing]
                    break # the inner for idxFollowing... loop
        return ans
