from typing import List

# Runtime complexity: O(sort) == O(nlogn)
# Auxiliary space complexity: O(1)
# Subjective level: hard
# Solved on: 2025-12-01
class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        totalCharge = sum(batteries)
        batteries.sort()

        numBatteries = n
        while batteries[-1] > totalCharge // numBatteries:
            totalCharge -= batteries.pop(-1)
            numBatteries -= 1

        # now all the batteries have charge <= average,
        # so all will be fully used - no waste.
        return totalCharge // numBatteries
