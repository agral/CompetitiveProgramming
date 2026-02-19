from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: very easy
# Solved on: 2026-02-19
class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        altitude = 0
        ans = 0
        for g in gain:
            altitude += g
            ans = max(ans, altitude)
        return ans
