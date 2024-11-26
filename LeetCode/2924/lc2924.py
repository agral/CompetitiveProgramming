from typing import List

class Solution:
    def findChampion(self, n: int, edges: List[List[int]]) -> int:
        in_degrees = [0] * n
        for _, v in edges:
            in_degrees[v] += 1

        if in_degrees.count(0) > 1:
            return -1
        return in_degrees.index(0) # guaranteed to exist.

def main():
    s = Solution()
    assert(s.findChampion(3, [[0, 1], [1, 2]]) == 0)
    assert(s.findChampion(4, [[0, 2], [1, 3], [1, 2]]) == -1)

if __name__ == "__main__":
    main()
