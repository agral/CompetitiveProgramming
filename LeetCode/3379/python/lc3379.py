from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy
# Solved on: 2026-02-05
class Solution:
    def constructTransformedArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n
        for i in range(n):
            ans[i] = nums[(i + nums[i] + (1000*n)) % n]

        return ans
