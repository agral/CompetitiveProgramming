from typing import List
import collections

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy
class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        count = collections.Counter(nums).values()
        max_count = max(count)
        max_count_elems = sum(c == max_count for c in count)
        return max_count_elems * max_count
