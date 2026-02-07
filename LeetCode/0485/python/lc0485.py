from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-02-07
class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        ans,  curr = 0, 0
        for num in nums:
            if num == 0:
                curr = 0
            else:
                curr += 1
                ans = max(ans, curr)
        return ans
