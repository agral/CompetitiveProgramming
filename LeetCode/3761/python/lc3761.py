from typing import List

# Runtime complexity: O(n*log10(n)) == O(nlogn).
#  this comes from: O(n) linear walk through the nums array, and at each step...
#  ...log10(n): taking the reverse of a number. Other steps are at most O(1) amortized time.
# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-04-17
class Solution:
    def reverse(self, num: int) -> int:
        ans = 0
        while num > 0:
            k = num % 10
            num //= 10
            ans *= 10
            ans += k

        return ans

    def minMirrorPairDistance(self, nums: List[int]) -> int:
        BIG_NUM = 71826576
        ans = BIG_NUM

        rev_seen = {} # maps the reversed number to when it was last seen
        for i, num in enumerate(nums):
            # we want reverse(i) == j.
            if num in rev_seen:
                ans = min(ans, i - rev_seen[num])
                if ans == 1:
                    return 1 # optimization, 1 is the minimum possible answer.

            rev_seen[self.reverse(num)] = i # mark reverse(num) (==r) as last seen here at index i.

        return -1 if ans == BIG_NUM else ans
