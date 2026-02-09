from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: medium-
# Solved on: 2026-02-09
class Solution:
    def timeRequiredToBuy(self, tickets: List[int], k: int) -> int:
        ans = 0
        for i in range(len(tickets)):
            kFactor = tickets[k] - 1 if i > k else tickets[k]
            ans += min(tickets[i], kFactor)
        return ans
