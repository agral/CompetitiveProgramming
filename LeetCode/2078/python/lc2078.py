from typing import List

# Runtime complexity: O(2*L) == O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-04-20
class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        L = len(colors)
        # The max difference will be either including the house to the very left,
        diffL, diffR = 0, 0
        for i in range(L-1, 0, -1):
            if colors[0] != colors[i]:
                diffL = i
                break
        # or to the house to the very right.
        for i in range(L-1):
            if colors[i] != colors[L-1]:
                diffR = L - 1 - i
                break
        #print(f"colors={colors}, diffL={diffL}, diffR={diffR}")
        return max(diffL, diffR)
