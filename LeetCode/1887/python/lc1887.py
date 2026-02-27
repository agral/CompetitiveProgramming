from typing import List

# Runtime complexity: O(sort)
# Auxiliary space complexity: O(sort)
# Subjective level: medium+ (coming up with the actual solution to this was a bit tricky,
#                   as duplicate numbers are allowed in the input data).
# Solved on: 2026-02-27
class Solution:
    def reductionOperations(self, nums: List[int]) -> int:
        L = len(nums)
        nums.sort()
        ans = 0
        for i in range(1, L):
            if nums[i] != nums[i-1]:
                ans += L - i
        return ans
