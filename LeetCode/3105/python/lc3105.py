from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-11-12
class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        ans = 1
        currDecreasing = 1
        currIncreasing = 1

        for i in range(1, len(nums)):
            if nums[i] > nums[i-1]:
                currIncreasing += 1
                currDecreasing = 1
            elif nums[i] < nums[i-1]:
                currDecreasing += 1
                currIncreasing = 1
            else:
                currDecreasing = 1
                currIncreasing = 1
            ans = max(ans, currIncreasing, currDecreasing)

        return ans
