import collections
from typing import List

class Solution:
    def minCost(self, grid: List[List[int]]) -> int:
        DIRS = ((0, 1), (0, -1), (1, 0), (-1, 0))
        H = len(grid)
        W = len(grid[0])
        dp = [[-1] * W for _ in range(H)]
        q = collections.deque()

        def dfs(h: int, w: int, cost: int) -> None:
            if h < 0 or h >= H or w < 0 or w >= W or dp[h][w] != -1:
                return
            dp[h][w] = cost
            dh, dw = DIRS[grid[h][w] - 1]
            q.append( (h, w) )
            dfs(h + dh, w + dw, cost)

        dfs(0, 0, 0)

        total_cost = 0
        while q:
            total_cost += 1
            for _ in range(len(q)):
                h, w = q.popleft()
                for dh, dw in DIRS:
                    dfs(h + dh, w + dw, total_cost)

        return dp[H-1][W-1]

def main():
    s = Solution()
    assert(s.minCost([[1, 1, 1, 1], [2, 2, 2, 2], [1, 1, 1, 1], [2, 2, 2, 2]]) == 3)
    assert(s.minCost([[1, 1, 3], [3, 2, 2], [1, 1, 4]]) == 0)
    assert(s.minCost([[1, 2], [3, 4]]) == 1)
    print("All testcases passed.")

if __name__ == "__main__":
    main()
