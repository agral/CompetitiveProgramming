from typing import List
import collections

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy+
# Solved on: 2026-02-08
class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        MAX_NUM = 101 # 0 <= nums[i] <= 100; that's 101 total possible numbers

        # 1. Count how many times each individual number appears in `nums`:
        # In the end, after step #02, counts_lte[i] holds how many numbers
        # less than OR EQUAL `i` exist in total in `nums`.
        counts_lte = [0] * MAX_NUM 
        for num in nums:
            counts_lte[num] += 1

        # 2. Now that the individual numbers are counted, include numbers
        # less than `i` too in `counts_lte[i]`:
        for num in range(1, MAX_NUM):
            counts_lte[num] += counts_lte[num-1]

        # 3. Construct and return the answer array:
        return [0 if num==0 else counts_lte[num-1] for num in nums]

    # Runtime complexity: O(n^2)
    # Auxiliary space complexity: O(n)
    # Subjective level: easy
    # Solved on: 2026-02-08
    def smallerNumbersThanCurrent_naive(self, nums: List[int]) -> List[int]:
        ans = []
        for num in nums:
            summ = 0
            for n in nums:
                if n < num:
                    summ += 1
            ans.append(summ)
        return ans
