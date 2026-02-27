from typing import List

# Runtime complexity: O(sort)
# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-02-27
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort()
        ans = [intervals[0]]
        for (start, stop) in intervals[1:]:
            if ans[-1][1] < start:
                ans.append([start, stop])
            else:
                ans[-1][1] = max(ans[-1][1], stop)

        return ans
