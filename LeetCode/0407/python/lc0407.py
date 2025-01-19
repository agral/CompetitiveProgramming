import heapq
from typing import List

class Solution:
    def trapRainWater(self, heightMap: List[List[int]]) -> int:
        DIRS = ((1, 0), (0, -1), (-1, 0), (0, 1)) # Top, Left, Bottom, Right
        H = len(heightMap)
        W = len(heightMap[0])
        minHeap = [] # holds tuples: (height, h, w)
        processed = set() # holds tuples: (h, w)
        ans = 0

        # process the border:
        for h in range(H):
            heapq.heappush(minHeap, (heightMap[h][0], h, 0))
            processed.add( (h, 0) )
            heapq.heappush(minHeap, (heightMap[h][W-1], h, W-1))
            processed.add( (h, W-1) )
        for w in range(1, W-1):
            heapq.heappush(minHeap, (heightMap[0][w], 0, w))
            processed.add( (0, w) )
            heapq.heappush(minHeap, (heightMap[H-1][w], H-1, w))
            processed.add( (H-1, w) )

        while minHeap:
            height, h, w = heapq.heappop(minHeap)
            for dh, dw in DIRS:
                hh, ww = h + dh, w + dw
                if hh < 0 or hh >= H or ww < 0 or ww >= W or (hh, ww) in processed:
                    continue

                if heightMap[hh][ww] < height:
                    # this cell will contain trapped water.
                    ans += height - heightMap[hh][ww]
                    heapq.heappush(minHeap, (height, hh, ww))
                else:
                    # this cell will not contain trapped water,
                    # but it might act as a wall for other cells.
                    heapq.heappush(minHeap, (heightMap[hh][ww], hh, ww))
                processed.add( (hh, ww) )

        return ans

def main():
    s = Solution()
    assert(s.trapRainWater([
           [1, 4, 3, 1, 3, 2],
           [3, 2, 1, 3, 2, 4],
           [2, 3, 3, 2, 3, 1],
    ]) == 4)
    assert(s.trapRainWater([
        [3, 3, 3, 3, 3],
        [3, 2, 2, 2, 3],
        [3, 2, 1, 2, 3],
        [3, 2, 2, 2, 3],
        [3, 3, 3, 3, 3],
    ]) == 10)
    print("All testcases successfully passed.")

if __name__ == "__main__":
    main()
