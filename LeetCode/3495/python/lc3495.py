from typing import List

# Runtime complexity: O(len(queries)), assuming bit operations are O(1).
# Auxiliary space complexity: O(1)
# Subjective level: hard.
class Solution:
    def minOperations(self, queries: List[List[int]]) -> int:
        def calcNumOps(upper: int) -> int:
            """Returns the count of operations required for [1, upper] range."""
            rv, numOp, powerOfFour = 0, 0, 1;
            while powerOfFour <= upper:
                l, r = powerOfFour, min(upper, 4*powerOfFour-1)
                numOp += 1
                rv += (r-l+1) * numOp
                powerOfFour *= 4
            return rv

        ans = 0
        for (l, r) in queries:
            ans += ((1 + calcNumOps(r) - calcNumOps(l-1)) // 2)
        return ans
