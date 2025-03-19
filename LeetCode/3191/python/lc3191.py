from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        SZ = len(nums)
        ans = 0
        for i in range(SZ - 2):
            if nums[i] == 0:
                ans += 1
                nums[i+1] ^= 1
                nums[i+2] ^= 1
        if nums[-2] == 0 or nums[-1] == 0:
            return -1
        return ans
