from typing import List

class Solution:
    def maxMatrixSum(self, matrix: List[List[int]]) -> int:
        negativeCount = sum(sum(1 for x in row if x < 0) for row in matrix)
        absoluteMinimum = min(min(abs(x) for x in row) for row in matrix)
        total = sum(sum(abs(x) for x in row) for row in matrix)
        if negativeCount % 2 == 1:
            # odd total number of negatives means that we can not get rid
            # of the last negative number - but this last negative number
            # can be any in the entire matrix; it can even be zero.
            # so, for odd total negatives count, need to subtract the
            # absoluteMinimum twice (once because it was initially added,
            # and once for the actual subtraction):
            total -= 2 * absoluteMinimum

        return total

def main():
    s = Solution()
    assert(s.maxMatrixSum([[1, -1], [-1, 1]]) == 4)
    assert(s.maxMatrixSum([[1, 2, 3], [-1, -2, -3], [1, 2, 3]]) == 16)

if __name__ == "__main__":
    main()
