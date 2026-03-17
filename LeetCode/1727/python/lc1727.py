from typing import List

# Runtime complexity: O(W*H) + O(sort) ~= O(W*H)
# Auxiliary space complexity: O(W)
# Subjective level: medium-
# Solved on: 2026-03-17
class Solution:
    def largestSubmatrix(self, matrix: List[List[int]]) -> int:
        W = len(matrix[0])
        histogram = [0] * W
        ans = 0
        for row in matrix:
            for i in range(W):
                histogram[i] = 0 if row[i] == 0 else histogram[i] + 1

            sortedHistogram = sorted(histogram)
            # histogram.sort()

            for i in range(W):
                ans = max(ans, sortedHistogram[i] * (W-i))
        return ans
