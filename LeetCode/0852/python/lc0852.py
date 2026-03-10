from typing import List

# Runtime complexity: O(log(n))
# Auxiliary space complexity: O(1)
# Subjective level: easy+
# Solved on: 2026-03-10
class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        L = len(arr)
        left = 0
        right = L-1
        while left < right:
            mid = (left + right) // 2
            if arr[mid] > arr[mid+1]:
                # mid-to-right contains the descending slope of the mountain.
                # the peak must be somewhere in the left-to-mid part.
                right = mid
            else: # otherwise, arr[mid] < arr[mid+1]. Thus the peak
                # must be somewhere in the mid-to-right part.
                left = mid+1
        return left
