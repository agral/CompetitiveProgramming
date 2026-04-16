from typing import List

# Runtime complexity: O(2N) == O(N)
# Auxiliary space complexity: O(N + Q)
# Subjective level: medium+
# Solved on: 2026-04-16
class Solution:
    def solveQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        Q = len(queries)
        N = len(nums)

        dist = [N] * N
        lastSeen = {}
        for n in range(2*N):
            i = n % N
            num = nums[i]
            if num in lastSeen:
                lastIndex = lastSeen[num] % N
                d = n - lastIndex
                dist[i] = min(dist[i], d)
                dist[lastIndex] = min(dist[lastIndex], d)
            lastSeen[num] = n

        ans = [-1] * Q
        for q, query in enumerate(queries):
            # If the number's distance to the other with same value is exactly N,
            # then it's an unique entry that can only "map" to itself.
            # When there are at least two entries with same value, the distance
            # will be less than N - exactly (i-j+N)%N.
            if dist[query] < N:
                ans[q] = dist[query]
        return ans
