from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-02-09
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = []
        ans = 0

        for i in range(len(heights) + 1):
            while stack and (i == len(heights) or heights[i] < heights[stack[-1]]):
                height = heights[stack.pop()]
                width = i
                if stack:
                    width = i - 1 - stack[-1]
                ans = max(ans, width * height)
            stack.append(i)

        return ans
