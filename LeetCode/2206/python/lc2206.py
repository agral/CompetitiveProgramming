from typing import List

class Solution:
    def divideArray(self, nums: List[int]) -> bool:
        seen = {}
        for n in nums:
            if n not in seen:
                seen[n] = 1
            else:
                seen[n] += 1

        for v in seen.values():
            if v % 2 == 1:
                return False

        return True
