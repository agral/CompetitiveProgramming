from typing import List

# Runtime complexity: O(H*W) ~= O(n^2)
# Auxiliary space complexity: O(H+W) ~= O(n)
# Subjective level: easy+
# Solved on: 2026-11-08
class Solution:
    def countServers(self, grid: List[List[int]]) -> int:
        H = len(grid)
        W = len(grid[0])
        rowSums = [0] * H # there are H rows
        colSums = [0] * W # and W columns in the grid.
        for h in range(H):
            for w in range(W):
                if grid[h][w] == 1:
                    rowSums[h] += 1
                    colSums[w] += 1

        ans = 0
        for h in range(H):
            for w in range(W):
                if grid[h][w] == 1 and (rowSums[h] > 1 or colSums[w] > 1):
                    ans += 1 # this machine communicates with some other(s).
        return ans
