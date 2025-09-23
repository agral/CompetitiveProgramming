from typing import List

# Runtime complexity: O(n), where n is the length of the version string.
# Auxiliary space complexity: O(n), for storing the converted numbers.
# Subjective level: easy.
class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        nums1 = [int(n) for n in version1.split(".")]
        nums2 = [int(n) for n in version2.split(".")]
        num_entries = max(len(nums1), len(nums2))
        for i in range(num_entries):
            v1 = nums1[i] if i < len(nums1) else 0
            v2 = nums2[i] if i < len(nums2) else 0
            if v1 < v2:
                return -1
            if v1 > v2:
                return 1
        return 0
