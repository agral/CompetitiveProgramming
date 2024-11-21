from typing import List

# NOTE:
# m corresponds to ROWS (y)
# n corresponds to COLUMNS (x)

class Solution:
    EMPTY = 0
    GUARDED = 1
    GUARD = 2
    WALL = 3
    def countUnguarded(self, m: int, n: int, guards: List[List[int]],
                       walls: List[List[int]]) -> int:
        level = [[0 for _ in range(n)] for _ in range(m)]

        for wall in walls:
            level[wall[0]][wall[1]] = self.WALL
        for guard in guards:
            level[guard[0]][guard[1]] = self.GUARD

            for c in range(guard[1] - 1, -1, -1): # guard to the left:
                if level[guard[0]][c] == self.EMPTY:
                    level[guard[0]][c] = self.GUARDED
                elif level[guard[0]][c] == self.GUARD or level[guard[0]][c] == self.WALL:
                    break
                # else do nothing, but carry on - the GUARDED field may be caused by a guard not on this line.
            for c in range(guard[1] + 1, n): # guard to the right:
                if level[guard[0]][c] == self.EMPTY:
                    level[guard[0]][c] = self.GUARDED
                elif level[guard[0]][c] == self.GUARD or level[guard[0]][c] == self.WALL:
                    break
                # else do nothing, but carry on - the GUARDED field may be caused by a guard not on this line.
            for r in range(guard[0] - 1, -1, -1): # guard to the top:
                if level[r][guard[1]] == self.EMPTY:
                    level[r][guard[1]] = self.GUARDED
                elif level[r][guard[1]] == self.GUARD or level[r][guard[1]] == self.WALL:
                    break
                # else do nothing, but carry on - the GUARDED field may be caused by a guard not on this line.
            for r in range(guard[0] + 1, m): # guard to the bottom:
                if level[r][guard[1]] == self.EMPTY:
                    level[r][guard[1]] = self.GUARDED
                elif level[r][guard[1]] == self.GUARD or level[r][guard[1]] == self.WALL:
                    break
                # else do nothing, but carry on - the GUARDED field may be caused by a guard not on this line.

        unguarded = 0
        for row in range(m):
            for col in range(n):
                if level[row][col] == self.EMPTY:
                    unguarded += 1
        return unguarded

def main():
    s = Solution()
    assert(s.countUnguarded(4, 6, [[0, 0], [1, 1], [2, 3]], [[0, 1], [2, 2], [1,4]]) == 7)
    assert(s.countUnguarded(3, 3, [[1, 1]], [[0, 1], [1, 0], [2, 1], [1, 2]]) == 4)
    assert(s.countUnguarded(2, 7, [[1, 5], [1, 1], [1, 6], [0, 2]], [[0, 6], [0, 3], [0, 5]]) == 1)

if __name__ == "__main__":
    main()
