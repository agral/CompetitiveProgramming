from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium+
# Solved on: 2026-02-19
class Solution:
    def waysToMakeFair(self, nums: List[int]) -> int:
        L = len(nums)
        sumsEven = [0 for _ in range(L+1)] # holds the sum of even-indexed values in nums<0..i)
        sumsOdd = [0 for _ in range(L+1)] # -||-, odd-indexed values in nums<0..i)
        ans = 0

        for i in range(1, L+1):
            sumsEven[i] = sumsEven[i-1]
            sumsOdd[i] = sumsOdd[i-1]
            if i % 2 == 0:
                sumsEven[i] += nums[i-1]
            else:
                sumsOdd[i] += nums[i-1]

        total = sumsEven[L] + sumsOdd[L] # == sum(nums), but that's already calculated.
        for i in range(L):
            # When removing ith element, all elements after i flip their evenness.
            # so sumsEven[i+1, ..., L-1] become odd, and sumsOdd[i+1, ..., L-1] become even.
            newEvenSum = sumsEven[i] + sumsOdd[L] - sumsOdd[i+1]
            newOddSum = total - (newEvenSum + nums[i])
            if newEvenSum == newOddSum:
                ans += 1
        return ans
