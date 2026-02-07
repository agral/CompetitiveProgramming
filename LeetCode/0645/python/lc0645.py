from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy
# Solved on: 2026-02-07
class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        n = len(nums)
        seen = [0] * n
        for i in range(n):
            seen[nums[i] - 1] += 1

        duplicate, absent = -1, -1
        for i in range(n):
            if seen[i] == 0:
                absent = i+1
            elif seen[i] == 2:
                duplicate = i+1
        return [duplicate, absent]
