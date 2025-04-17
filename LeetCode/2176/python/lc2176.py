from typing import List
import collections
import math

# Runtime complexity: unknown, not larger than O(nk).
# Auxiliary space complexity: O(n).
# n is the length of nums array.
class Solution:
    def countPairs(self, nums: List[int], k: int) -> int:
        ans = 0
        numToIndex = collections.defaultdict(list)
        for idx, num in enumerate(nums):
            numToIndex[num].append(idx)

        for index in numToIndex.values():
            count_gcds = collections.Counter()
            for idx in index:
                gcd_first = math.gcd(idx, k)
                for gcd_second, count in count_gcds.items():
                    if gcd_first * gcd_second % k == 0:
                        ans += count
                count_gcds[gcd_first] += 1
        return ans
