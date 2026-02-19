from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-02-19
class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        L = len(nums)
        prefixSumModKToIndex = {0: -1}
        prefixSumModK = 0

        for idx in range(L):
            prefixSumModK += nums[idx]
            prefixSumModK %= k
            if prefixSumModK in prefixSumModKToIndex:
                if idx - prefixSumModKToIndex[prefixSumModK] >= 2:
                    return True
            else:
                prefixSumModKToIndex[prefixSumModK] = idx

        return False
