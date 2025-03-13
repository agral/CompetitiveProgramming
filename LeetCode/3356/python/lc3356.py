from typing import List

# Time complexity: O(n)
# Aux space complexity: O(n)

class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        q = 0
        line = [0] * (len(nums) + 1)
        decrement = 0
        for i, num in enumerate(nums):
            while decrement + line[i] < num:
                if q == len(queries):
                    return -1
                li, ri, vali = queries[q]
                q += 1
                if ri >= i:
                    line[max(li, i)] += vali
                    line[ri + 1] -= vali
            decrement += line[i]
        return q
