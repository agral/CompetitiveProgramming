import math
from typing import List

# Runtime complexity: O(nlogn)
# Auxiliary space complexity: O(n)
# Subjective level: medium
class Solution:
    def replaceNonCoprimes(self, nums: List[int]) -> List[int]:
        ans = []
        for num in nums:
            while ans and math.gcd(ans[-1], num) > 1:
                num = math.lcm(num, ans.pop())
            ans.append(num)
        return ans
