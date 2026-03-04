from typing import List

# Runtime complexity: O(H*W)
# Auxiliary space complexity: O(H+W)
# Subjective level: easy
# Solved on: 2026-03-04
class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        H = len(mat)
        W = len(mat[0])

        sumRows = [0] * H
        sumCols = [0] * W
        for h in range(H):
            for w in range(W):
                sumRows[h] += mat[h][w]
                sumCols[w] += mat[h][w]

        ans = 0
        for h in range(H):
            for w in range(W):
                if mat[h][w] == 1 and sumRows[h] == 1 and sumCols[w] == 1:
                    ans += 1
        return ans
