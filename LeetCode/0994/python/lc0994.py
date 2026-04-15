from typing import List
from collections import deque

# Runtime complexity: O(H*W)
# Auxiliary space complexity: O(H*W) in general case; depends on the number of cells enqueued.
# Subjective level: medium.
# Solved on: 2026-04-14
class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        H = len(grid)
        W = len(grid[0])
        FIELD_ORANGE = 1
        FIELD_ROTTEN = 2
        DIRS = ((-1, 0), (0, 1), (1, 0), (0, -1))

        q = deque()
        totalFresh = 0
        for h in range(H):
            for w in range(W):
                if grid[h][w] == FIELD_ORANGE:
                    totalFresh += 1
                elif grid[h][w] == FIELD_ROTTEN:
                    q.append((h, w))

        # handle the corner case where there's no fresh oranges to begin with:
        if totalFresh == 0:
            return 0

        tick = 0
        while q: # in Python this means while !q.empty(). Can use this for any container. Nice!
            tick += 1
            for _ in range(len(q)):
                (h, w) = q.popleft() # note: want a FIFO queue? Need to call popleft(). Using pop() makes it a LIFO queue.
                for dir in DIRS:
                    nh, nw = h+dir[0], w+dir[1] # neighbor's position (nh, nw)
                    if nh < 0 or nh >= H or nw < 0 or nw >= W:
                        continue
                    # all the fresh oranges that neighbor a rotten one start rotting too:
                    if grid[nh][nw] == FIELD_ORANGE:
                        grid[nh][nw] = FIELD_ROTTEN
                        q.append((nh, nw))
                        totalFresh -= 1

        return -1 if totalFresh > 0 else tick-1
