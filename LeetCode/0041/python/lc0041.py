from typing import List

# Runtime complexity: O(2n) == O(n)
# Auxiliary space complexity: O(1) (but the passed `nums` is modified! That's ok though.)
# Subjective level: hard
# Solved on: 2026-02-18
class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        L = len(nums)
        for i in range(L):
            # swap nums[nums[i]-1] with nums[i] for each mismatching number:
            while nums[i] > 0 and nums[i] <= L and nums[nums[i]-1] != nums[i]:
                nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]

        for i in range(L):
            if nums[i] != i+1:
                return i+1
        return L+1
