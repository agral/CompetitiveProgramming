from typing import List
import math

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-04-13
class Solution:
    def getMinDistance(self, nums: List[int], target: int, start: int) -> int:
        L = len(nums)
        ans = L
        for i in range(L):
            if nums[i] == target:
                ans = min(ans, abs(start-i))
        return ans
