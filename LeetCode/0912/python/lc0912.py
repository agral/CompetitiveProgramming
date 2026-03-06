from typing import List
import random

# Subjective level: medium+ / hard-
# Solved on: 2026-02-28
class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        # quicksort seems to work badly for the annoying testcase consisting of some thousands of '2's.
        # self.qsort(nums, 0, len(nums)-1)
        self.mergesort(nums, 0, len(nums)-1)
        return nums

    # Runtime complexity: O(nlogn)
    # Auxiliary space complexity: O(n)
    def mergesort(self, nums: List[int], l: int, r: int) -> None:
        if l >= r:
            return
        def do_merge(nums: List[int], l: int, mid: int, r: int) -> None:
            L = r-l+1 # the length of the sorted part
            sorted = [0] * L
            curr = 0 # index currently being sorted
            ml = l # merge range's left index
            mr = mid+1 # merge range's right index
            while ml <= mid and mr <= r:
                if nums[ml] < nums[mr]:
                    sorted[curr] = nums[ml]
                    ml += 1
                else:
                    sorted[curr] = nums[mr]
                    mr += 1
                curr += 1

            # There might be remaining left and right parts; put both into the sorted array:
            while ml <= mid:
                sorted[curr] = nums[ml]
                ml += 1
                curr += 1
            while mr <= r:
                sorted[curr] = nums[mr]
                mr += 1
                curr += 1

            nums[l:l+L] = sorted

        mid = (l+r) // 2
        self.mergesort(nums, l, mid)
        self.mergesort(nums, mid+1, r)
        do_merge(nums, l, mid, r)

    # Runtime complexity: O(nlogn) [optimal, usual, but up to O(n^2) for specific inputs/bad pivot choices]
    # Auxiliary space complexity: O(n)
    def qsort(self, nums: List[int], l: int, r: int) -> None:
        if l >= r:
            return

        def partition(l: int, r: int) -> int:
            pivotIdx = l + random.randint(0, r-l)
            pivot = nums[pivotIdx]
            nums[pivotIdx], nums[r] = nums[r], nums[pivotIdx]
            swapIdx = l
            for i in range(l, r):
                if nums[i] <= pivot:
                    nums[swapIdx], nums[i] = nums[i], nums[swapIdx]
                    swapIdx += 1
            nums[swapIdx], nums[r] = nums[r], nums[swapIdx]
            return swapIdx

        mid = partition(l, r)
        self.qsort(nums, l, mid-1)
        self.qsort(nums, mid+1, r)
