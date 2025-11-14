from typing import List

# Runtime complexity: O(n^2)
# Auxiliary space complexity: O(n^2)
# Subjective level: hard.
#  -> but it's quite neat, and once you know how to solve it, it's a mid problem.
#  -> I definitely would not invent this algorithm during a code interview / time pressure.
# Solved on: 2025-11-14
class Solution:
    """A naive solution. It works 100% correctly, but is surely too slow to pass - O(n^3)."""
    def naive_rangeAddQueries(self, n: int, queries: List[List[int]]) -> List[List[int]]:
        ans = [[0] * n for _ in range(n)]
        for query in queries:
            for y in range(query[0], query[2]+1, 1):
                for x in range(query[1], query[3]+1, 1):
                    ans[y][x] += 1
            print(ans)
        print(ans)
        return ans

    """An actual solution, works in O(n^2) for n range updates."""
    def rangeAddQueries(self, n: int, queries: List[List[int]]) -> List[List[int]]:
        ans = [[0] * n for _ in range(n)]
        # prefixSums[y][x] == ans[y][x+1] - ans[y][x]
        prefixSums = [[0] * (n+1) for _ in range(n)]
        for query in queries:
            for y in range(query[0], query[2]+1, 1):
                prefixSums[y][query[1]] += 1
                prefixSums[y][query[3]+1] -= 1

        for y in range(n):
            cumulativeSum = 0
            for x in range(n):
                cumulativeSum += prefixSums[y][x]
                ans[y][x] = cumulativeSum
        return ans
