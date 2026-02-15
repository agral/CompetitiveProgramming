from typing import List
import heapq

# Runtime complexity: O(n*logn + k*logn), where k is proportional to the biggest target value (I think..?)
# Auxiliary space complexity: O(n)
# Subjective level: hard.
# Solved on: 2026-02-10 to 2026-02-15
class Solution:
    def isPossible(self, target: List[int]) -> bool:
        T = len(target)

        # Handle the corner case where there's just one number in the target array.
        if T == 1:
            return target[0] == 1

        # note to self: with max-heap/min-heap, there's no need to sort.
        # target.sort()
        total = sum(target)
        heapq.heapify_max(target)

        while target[0] > 1:
            curr = heapq.heappop_max(target)
            total_minus_curr = total - curr

            # another corner case; occurs only for target of length 2:
            if total_minus_curr == 1:
                return True

            modulo = curr % total_minus_curr
            if modulo == 0 or modulo == curr:
                # When curr divides total-minus-curr evenly (modulo == 0),
                # or when there's no change (modulo == curr), no progress can be done.
                return False

            heapq.heappush_max(target, modulo)
            total = total - curr + modulo

        return True
