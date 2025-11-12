import math
from typing import List

# Runtime complexity: O(n^2)*O(math.gcd())
# Auxiliary space complexity: O(1)
# Subjective level: medium+ (required some planning with pen & paper)
# Solved on: 2025-11-12
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        n = len(nums)
        onesCount = nums.count(1)
        if onesCount > 0:
            # If at least one "1" is present, then one can easily zero the entire array
            # by using an operation on "1" and other numbers that are not "1" themselves.
            return n - onesCount

        # Assume that the gcd of all the possible pairs of numbers is composite; i.e. >1.
        isComposite = True 
        opsCount = math.inf
        for i in range(n):
            for j in range(i+1, n):
                # WA #01: some numbers are non-coprime, but become coprime after at least one op.
                # gcd = math.gcd(nums[i], nums[j])
                nums[i] = math.gcd(nums[i], nums[j])
                if nums[i] == 1:
                    opsCount = min(opsCount, j-i)
                    isComposite = False
                    break

        # if isComposite:
        if opsCount == math.inf:
            # The gcd of all the possible pairs of numbers was greater than one.
            # (there were no coprime numbers among all the possible pairs)
            return -1

        # After reaching a gcd of 1 with opsCount, still need extra (n-1) operations
        # to make all the other numbers equal to 1.
        # (note: we're not taking the presence of "1"s into account now, as this case has already
        # been handled above; runtime entering here means no raw "1"s in the input data).
        return (n - 1) + int(opsCount)
