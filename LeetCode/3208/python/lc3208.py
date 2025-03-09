from typing import List

class Solution:
    def numberOfAlternatingGroups(self, colors: List[int], k: int) -> int:
        n = len(colors)
        num_alternating = 1
        ans = 0
        for i in range(n + k - 2):
            if colors[i%n] != colors[(i-1)%n]:
                num_alternating += 1
            else:
                num_alternating = 1
            if num_alternating >= k:
                ans += 1
        return ans
