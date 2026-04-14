from typing import List

# Runtime complexity: O(H*W)
# Auxiliary space complexity: O(H*W)
# Subjective level: medium+
# Solved on: 2026-04-14
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        DIRS = ((-1, 0), (0, 1), (1, 0), (0, -1))
        LAND = "1"
        PROCESSED_LAND = "G"
        H = len(grid)
        W = len(grid[0])

        def dfs(h: int, w: int) -> None:
            if h < 0 or h >= H or w < 0 or w >= W or grid[h][w] != LAND:
                return
            grid[h][w] = PROCESSED_LAND
            for (dh, dw) in DIRS:
                dfs(h+dh, w+dw)

        ans = 0
        for h in range(H):
            for w in range(W):
                if grid[h][w] == LAND:
                    ans += 1
                    dfs(h, w)
        return ans
