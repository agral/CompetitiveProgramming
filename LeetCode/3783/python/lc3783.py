from typing import List

# Runtime complexity: O(log10(n))
# Auxiliary space complexity: O(1)
# Subjective level: very easy
# Solved on: 2026-04-18
class Solution:
    def reverse(self, num: int) -> int:
        ans = 0
        while num > 0:
            k = num % 10
            num //= 10
            ans *= 10
            ans += k
        return ans

    def mirrorDistance(self, n: int) -> int:
        return abs(n - self.reverse(n))
