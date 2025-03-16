import math
from typing import List

# This is very similar to LeetCode #2560, which was daily challenge I solved yesterday.
class Solution:
    def repairCars(self, ranks: List[int], cars: int) -> int:
        def countFixedCars(minutes_available: int) -> int:
            # rank[i] * num_cars_fixed ^ 2 == minutes_available,
            # so num_cars_fixed == (minutes_available / rank[i]) ^ 0.5
            return sum(int(math.sqrt(minutes_available // rank)) for rank in ranks)

        left = 0
        right = min(ranks) * cars * cars
        while left < right:
            mid = (left + right) // 2
            if (countFixedCars(mid) < cars):
                left = mid + 1
            else:
                right = mid
        return left
