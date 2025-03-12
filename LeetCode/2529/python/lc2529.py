from typing import List

class Solution:
    # An O(n) standard search algorithm.
    def maximumCount(self, nums: List[int]) -> int:
        return max(sum(n < 0 for n in nums), sum(n > 0 for n in nums))
