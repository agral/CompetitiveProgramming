from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-02-19
class Solution:
    def minSubarray(self, nums: List[int], p: int) -> int:
        L = len(nums)
        S = sum(nums)
        rem = S % p
        if rem == 0:
            return 0

        prefixSumModp = 0
        prefixSumToIndex = {0: -1} # maps prefix sum modulo p to array index.
        ans = L
        for i, num in enumerate(nums):
            prefixSumModp += num
            prefixSumModp %= p
            target = (prefixSumModp + p - rem) % p
            if target in prefixSumToIndex:
                ans = min(ans, i - prefixSumToIndex[target])
            prefixSumToIndex[prefixSumModp] = i

        return -1 if ans == L else ans
