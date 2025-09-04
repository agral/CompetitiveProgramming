from typing import List

# Runtime complexity: O(1)
# Auxiliary space complexity: O(1)
# Subjective level: easy
class Solution:
    def findClosest(self, x: int, y: int, z: int ) -> int:
        xz = abs(x - z)
        yz = abs(y - z)
        if (xz < yz):
            return 1
        if (yz < xz):
            return 2
        return 0
