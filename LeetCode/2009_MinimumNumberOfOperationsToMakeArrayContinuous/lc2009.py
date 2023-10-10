# Note: this is the most optimal solution, I think.
# It would be O(n), but the array needs to be sorted in the beginning,
# which itself is O(n logn).
class Solution:
    def minOperations(self, nums) -> int:
        initial_length = len(nums)
        nums = sorted(set(nums)) # make nums sorted and with duplicates removed.
        ans = initial_length
        right = 0
        for left in range(len(nums)):
            while right < len(nums) and nums[right] < nums[left] + initial_length:
                right = right + 1
            ans = min(ans, initial_length - (right - left))

        return ans
