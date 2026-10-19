from typing import List

# Runtime complexity: O(n) (note: no sorting!)
# Auxiliary space complexity: O(1)
# Subjective level: easy+
# Solved on: 2026-10-19
class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        countLowerThanTarget = 0
        countTarget = 0
        for num in nums:
            if num < target:
                countLowerThanTarget += 1
            elif num == target:
                countTarget += 1

        # Special case #01: no target numbers in nums. Return an empty List.
        if countTarget == 0:
            return []

        # Otherwise after sorting, there will be `countLowerThanTarget` numbers first,
        # and then `countTarget` numbers that equal `target`. So, the answer:
        return list(range(countLowerThanTarget, countLowerThanTarget+countTarget))
