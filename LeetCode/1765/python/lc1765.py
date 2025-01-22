from typing import List
import collections

class Solution:
    def highestPeak(self, isWater: List[List[int]]) -> List[List[int]]:
        H = len(isWater)
        W = len(isWater[0])
        DIRS = ((1, 0), (0, 1), (-1, 0), (0, -1))
        q = collections.deque()
        ans = [[-1] * W for _ in range(H)]

        # Seed the initial contents of the queue - all the spaces with water;
        # these are guaranteed to have elevation of 0.
        for h in range(H):
            for w in range(W):
                if isWater[h][w] == 1:
                    ans[h][w] = 0
                    q.append( (h, w) )

        while q:
            h, w = q.popleft()
            for (dh, dw) in DIRS:
                y, x = h + dh, w + dw
                if y < 0 or y >= H or x < 0 or x >= W or ans[y][x] != -1:
                    continue
                ans[y][x] = ans[h][w] + 1
                q.append((y, x))

        return ans

def main():
    s = Solution()
    assert(s.highestPeak([[0, 1], [0, 0]]) == [[1, 0], [2, 1]])
    assert(s.highestPeak([[0, 0, 1], [1, 0, 0], [0, 0, 0]]) == [[1, 1, 0], [0, 1, 1], [1, 2, 2]])
    print("All testcases passed successfully.")

if __name__ == "__main__":
    main()
