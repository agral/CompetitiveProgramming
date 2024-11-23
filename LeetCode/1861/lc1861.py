from typing import List

class Solution:
    STONE = '#'
    OBSTACLE = '*'
    EMPTY = '.'
    def rotateTheBox(self, box: List[List[str]]) -> List[List[str]]:
        H = len(box)
        W = len(box[0])

        # initially create an array rotated by 90* (note swapped H and W) filled with empty spaces:
        ans = [[self.EMPTY for _ in range(H)] for _ in range(W)]

        # count the number of stones until obstacle / end-of-source-array column.
        # When obstacle/end hit, paste the obstacle, then the stones just above.
        for h in range(H):
            stones = 0
            for w in range(W):
                if box[h][w] == self.STONE:
                    stones += 1
                elif box[h][w] == self.OBSTACLE:
                    ans[w][H-1-h] = self.OBSTACLE
                    for s in range(stones):
                        ans[w-1-s][H-1-h] = self.STONE
                    stones = 0
            # immediately after row is finished, fill it up from the bottom with remaining stones
            # (same as with an obstacle, except no OBSTACLE is spawned):
            for s in range(stones):
                ans[W-1-s][H-1-h] = self.STONE
            # no need to zero-out the used stones; will get overwritten in next iteration anyway.
        return ans

def main():
    s = Solution()
    assert(s.rotateTheBox([['#', '.', '#']]) == [['.'],
                                                 ['#'],
                                                 ['#']])
    assert(s.rotateTheBox([['#', '.', '*', '.'],
                           ['#', '#', '*', '.']]) == [['#', '.'],
                                                      ['#', '#'],
                                                      ['*', '*'],
                                                      ['.', '.']])
    assert(s.rotateTheBox([['#', '#', '*', '.', '*', '.'],
                           ['#', '#', '#', '*', '.', '.'],
                           ['#', '#', '#', '.', '#', '.']]) == [['.', '#', '#'],
                                                                ['.', '#', '#'],
                                                                ['#', '#', '*'],
                                                                ['#', '*', '.'],
                                                                ['#', '.', '*'],
                                                                ['#', '.', '.']])

if __name__ == "__main__":
    main()
