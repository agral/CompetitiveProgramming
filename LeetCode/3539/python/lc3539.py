import functools
import math
from typing import List

# Runtime complexity: O(m^3 * kn)
# Auxiliary space complexity: O(m^2 * kn)
# Subjective level: very hard.
#   rationale: this is not only a hard dp problem; it involves combinatorics.
#      also all the examples are cases where k == m; so there's some work required
#      to handle cases where k != m...
class Solution:
    def magicalSum(self, m: int, k: int, nums: List[int]) -> int:
        MOD = 1_000_000_007

        @functools.lru_cache(None)
        def dp(m: int, k: int, i: int, carry: int) -> int:
            """Returns the count of magical sequences of length `k`
               that can be formed from the first `i` numbers in `nums`
               using at most `m` elements."""
            if m < 0 or k < 0 or (m + carry.bit_count() < k):
                return 0
            if m == 0:
                return 1 if k == carry.bit_count() else 0
            if i == len(nums):
                return 0

            ans = 0
            for c in range(m + 1):
                delta = math.comb(m, c) * pow(nums[i], c, MOD) % MOD
                nextCarry = carry + c
                ans += dp(m - c, k - (nextCarry % 2), i + 1, nextCarry // 2) * delta
                ans %= MOD
            return ans

        return dp(m, k, 0, 0)
