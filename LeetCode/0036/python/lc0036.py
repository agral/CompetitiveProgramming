from typing import List

# Runtime complexity: O(9*9) == O(1)
# Auxiliary space complexity: O(3*9*9) == O(1)
# Subjective level: medium leaning into easy.
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [[False]*9 for row in range(9)]
        cols = [[False]*9 for row in range(9)]
        grids = [[False]*9 for row in range(9)]

        for h in range(9):
            for w in range(9):
                if board[h][w] == '.':
                    continue
                num = int(board[h][w]) - 1

                if rows[h][num]:
                    return False
                rows[h][num] = True

                if cols[w][num]:
                    return False
                cols[w][num] = True

                gridnum = 3 * (h//3) + w//3
                if grids[gridnum][num]:
                    return False
                grids[gridnum][num] = True
        return True
