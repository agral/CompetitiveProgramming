from typing import List

# Runtime complexity: O(sort) == O(n*logn)
# Auxiliary space complexity: O(1)
class Solution:
    def countDays(self, days: int, meetings: List[List[int]]) -> int:
        freeDays, previousEnd = 0, 0
        meetings.sort()
        for start, end in meetings:
            if start > previousEnd:
                freeDays += start - previousEnd - 1
            previousEnd = max(previousEnd, end)

        return freeDays + max(days - previousEnd, 0)
