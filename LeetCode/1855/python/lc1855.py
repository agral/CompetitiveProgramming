from typing import List

# Runtime complexity: O(max(len(nums1), len(nums2))) ~= O(n)
# Auxiliary space complexity: O(1)
# Subjective level: medium-
# Solved on: 2026-04-19
class Solution:
    def maxDistance(self, nums1: List[int], nums2: List[int]) -> int:
        i = 0 # currently processed index of nums1
        j = 0 # currently processed index of nums2
        ans = 0

        while i < len(nums1) and j < len(nums2):
            if nums1[i] <= nums2[j]:
                # this is a valid pair.
                ans = max(ans, j - i)
                j += 1
            else:
                # this is not a valid pair.
                i += 1

        return ans
