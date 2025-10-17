from typing import List
import functools

# Runtime complexity: O(26*n) == O(n)
# Auxiliary space complexity: O(26*n) == O(n)
# Subjective level: hard.
class Solution:
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        def calc(i: int, bitmask: int, orBit: int, nextCanChange: bool) -> int:
            mask = bitmask | orBit
            if mask.bit_count() > k:
                return 1 + dp(i+1, nextCanChange, orBit)
            return dp(i+1, nextCanChange, mask)

        @functools.lru_cache(None)
        def dp(i: int, canChange: bool, bitmask: int) -> int:
            """Returns the count of possible partitions of s[i:n]; where
            `canChange` is true if the letter can still be changed,
            `bitmask` is the bitmask of the letters sen so far."""
            if i == len(s):
                return 0
            # Calculate the initial answer using the unmodified current letter:
            ans = calc(i, bitmask, 1 << (ord(s[i]) - ord('a')), canChange)

            # Change the current letter if possible; update the answer if another letter is strictly better:
            if canChange:
                for offset in range(ord('z')-ord('a')+1):
                    ans = max(ans, calc(i, bitmask, (1 << offset), False))
            return ans

        return 1 + dp(0, True, 0)
