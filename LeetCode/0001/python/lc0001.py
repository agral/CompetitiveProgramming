from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy
# Solved on: 2026-02-18
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numToIndex = {}
        for i in range(len(nums)):
            if (target - nums[i]) in numToIndex:
                return [numToIndex[target - nums[i]], i]
            numToIndex[nums[i]] = i

        # Exactly one valid answer is guaranteed to exist, so this won't happen for valid inputs.
        return [-42]
