from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy+
# Solved on: 2026-04-15
class Solution:
    def closestTarget(self, words: List[str], target: str, startIndex: int) -> int:
        L = len(words)
        BIG_INT = 71826576 # more than the total possible length of `words`, which is 100.
        ans = BIG_INT
        for i, s in enumerate(words):
            if s == target:
                distForwards = (startIndex - i + L) % L
                distBackwards = (i - startIndex + L) % L
                ans = min(ans, distForwards, distBackwards)
        if ans == BIG_INT:
            return -1
        return ans
