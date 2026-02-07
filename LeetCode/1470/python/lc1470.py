from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(2n) == O(n)
# Subjective level: easy
# Solved on: 2026-02-07
class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        ans = [0] * (2 * n)
        for k in range(n):
            ans[2*k] = nums[k]
            ans[2*k+1] = nums[n+k]

        return ans
