from typing import List

# Runtime complexity: O(n^2)
# Could maybe do better with some priority queue? Right now the solution requires
# recalculating the neighboring sums for all the pairs at every step of the algorithm;
# as long as the list is not strictly nondecreasing. That's O(n) for such a summation,
# done at most n times - O(n^2) in total.
# But this wastes a lot of computation in case of long arrays - maybe there's some more
# clever way to do this. Maybe a hash-map of pair of numbers to their sum?
# On the other hand it's just a simple sum. I've decided to leave this as is.

# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-01-22
class Solution:
    def isNonDecreasing(self, nums: List[int]) -> bool:
        for i in range(1, len(nums)):
            if nums[i-1] > nums[i]:
                return False
        return True

    def minimumPairRemoval(self, nums: List[int]) -> int:
        ans = 0
        while not self.isNonDecreasing(nums):
            neighborSums = [nums[i-1]+nums[i] for i in range(1, len(nums))]
            minSum = min(neighborSums)
            minIndex = neighborSums.index(minSum)

            # Replace the two neighbors of minimal sum found above with one entry - their sum:
            nums.pop(minIndex+1)
            nums[minIndex] = minSum
            ans += 1
        return ans

