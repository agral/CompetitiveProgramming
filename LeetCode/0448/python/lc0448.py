from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy
# Solved on: 2026-02-08
class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums)
        seen = [False] * (n+1)
        for num in nums:
            seen[num] = True
        ans = []
        for i in range(1, n+1):
            if not seen[i]:
                ans.append(i)

        return ans
