from sortedcontainers import SortedList
from typing import List

# Runtime complexity: O(n*logn)
# Auxiliary space complexity: O(n)
# Subjective level: hard+.
# Solved on: 2026-02-02
class Solution:
    def minimumCost(self, nums: List[int], k: int, dist: int) -> int:
        windowSum = 0
        for i in range(1, dist+2):
            windowSum += nums[i]
        candidates = SortedList()
        selected = SortedList(nums[i] for i in range(1, dist+2))

        def rebalance() -> int:
            """Updates the `windowSum`. Balances the multiset `selected` so that
            the top (k-1) numbers are kept. Returns the rebalanced `windowSum`.
            """
            nonlocal windowSum
            while len(selected) < (k-1):
                minCandidate = candidates[0]
                windowSum += minCandidate
                selected.add(minCandidate)
                candidates.remove(minCandidate)
            while len(selected) > (k-1):
                maxSelected = selected[-1]
                windowSum -= maxSelected
                selected.remove(maxSelected)
                candidates.add(maxSelected)
            return windowSum

        windowSum = rebalance()
        minWindowSum = windowSum

        # Find the optimal arrangement:
        for i in range(dist+2, len(nums)):
            tooDistant = nums[i - dist - 1]
            if tooDistant in selected:
                windowSum -= tooDistant
                selected.remove(tooDistant)
            else:
                candidates.remove(tooDistant)

            if nums[i] < selected[-1]:
                # nums[i] is immediately a better match.
                windowSum += nums[i]
                selected.add(nums[i])
            else:
                # nums[i] might be reconsidered in the future for rebalancing:
                candidates.add(nums[i])

            windowSum = rebalance()
            minWindowSum = min(minWindowSum, windowSum)

        return minWindowSum + nums[0]
