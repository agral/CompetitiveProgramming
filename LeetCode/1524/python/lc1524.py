from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy+
# Solved on: 2026-11-15
class Solution:
    def numOfSubarrays(self, arr: List[int]) -> int:
        ans = 0
        dp_even = 0
        dp_odd = 0
        for num in arr:
            if num % 2 == 0:
                dp_even += 1
            else:
                dp_even, dp_odd = dp_odd, dp_even+1
            ans += dp_odd
        return ans % 1_000_000_007
