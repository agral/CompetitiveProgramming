from typing import List
from itertools import pairwise

class Solution:
    def isArraySpecial(self, nums: List[int], queries: List[List[int]]) -> List[bool]:
        ans = []
        id = 0

        # holds IDs enumerating _special_ partitioning of input array.
        # _special_ partitioning ID increases whenever adjacent two numbers share same parity.
        # _special_ partitioning starts with ID: 0 for 0th element.
        specialPartitioning = [0]
        for left, right in pairwise(nums):
            if left % 2 == right % 2:
                id += 1
            specialPartitioning.append(id)

        for q in queries:
            ans.append(specialPartitioning[q[0]] == specialPartitioning[q[1]])

        return ans

def main():
    s = Solution()
    assert(s.isArraySpecial([3, 4, 1, 2, 6], [[0, 4]]) == [False])
    assert(s.isArraySpecial([4, 3, 1, 6], [[0, 2], [2, 3]]) == [False, True])
    print("All tests have passed.")

if __name__ == "__main__":
    main()
