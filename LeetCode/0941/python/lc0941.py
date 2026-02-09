from typing import List

# Runtime complexity: O(L)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-02-09
class Solution:
    def validMountainArray(self, arr: List[int]) -> bool:
        L = len(arr)
        if L < 3:
            return False

        left, right = 0, L-1
        while (left+1 < L) and (arr[left] < arr[left+1]):
            left += 1
        while (right > 0) and (arr[right-1] > arr[right]):
            right -= 1

        return (left > 0) and (right < L-1) and (left == right)
