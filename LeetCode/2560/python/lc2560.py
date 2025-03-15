from typing import List

class Solution:
    def minCapability(self, nums: List[int], k: int) -> int:
        def countRobbedHouses(robbingCapacity: int) -> int:
            ans = 0
            # for i in range(len(nums)): #won't work.
            i = 0
            SZ = len(nums)
            while i < SZ:
                if nums[i] <= robbingCapacity:
                    ans += 1
                    i += 1
                i += 1
            return ans

        left = min(nums)
        right = max(nums)
        while left < right:
            mid = (left + right) // 2
            if (countRobbedHouses(mid) < k):
                left = mid+1
            else:
                right = mid
        return left
