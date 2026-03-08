from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n) # for holding the answer string. Without that, O(1).
# Subjective level: easy
# Solved on: 2026-03-08
class Solution:
    def findDifferentBinaryString(self, nums: List[str]) -> str:
        L = len(nums)
        ans = ['0'] * L
        for i in range(L):
            if nums[i][i] == '0':
                ans[i] = '1'
            # otherwise it remains a '0'.
        # Now that ans has different ith "bit" than each of its sources in `nums`,
        # it is guaranteed to be different than each of the numbers in `nums`.
        return "".join(ans)
