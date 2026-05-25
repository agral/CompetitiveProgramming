from typing import List

# Re-solved in Python, based on my Golang solution from 2026-03-26.

# Runtime complexity: O(log(n))
# Auxiliary space complexity: O(1)
# Subjective level: easy+
# Solved on: 2026-05-22
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        L = len(nums)
        left, right = 0, L-1

        while left < right:
            # Divide the range in half; find the part that contains a pivot
            # (pivot present when subrange_left > subrange_right).
            mid = (left+right) // 2
            if nums[left] <= nums[mid]:
                # left part does not contain a pivot point
                if nums[left] <= target and target <= nums[mid]:
                    # the target is contained somewhere within the left part. Continue searching there.
                    right = mid
                else:
                    # the target is contained somewhere within the right partition.
                    # Continue search there.
                    left = mid+1
            else:
                # left part contains a pivot point.
                if nums[mid+1] <= target and target <= nums[right]:
                    left = mid+1
                else:
                    right = mid
        return left if nums[left] == target else -1
