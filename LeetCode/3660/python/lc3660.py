from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium+ / hard-
# Solved on: 2026-05-07 (started); finished on 2026-11-22
class Solution:
    def maxValue(self, nums: List[int]) -> List[int]:
        L = len(nums)

        # Calculate the maximum prefix array:
        prefixMax = [0] * L
        prefixMax[0] = nums[0]
        for i in range(1, L):
            prefixMax[i] = max(prefixMax[i-1], nums[i])

        # Calculate the answer, keep track of the minimum suffix seen so far.
        # This can be done in a single backwards sweep.
        ans = [0] * L
        suffixMin = float('inf')
        for i in range(L-1, -1, -1): # that's Python for golang's: `for i := L-1; i >= 0; i--`
            if i < L-1: # if there are elements after this element; i.e. it's not the last one
                if prefixMax[i] > suffixMin:
                    ans[i] = ans[i+1] # then use the value from the right (if it's possible to jump forward);
                else:
                    ans[i] = prefixMax[i] # otherwise use the local prefix-max.
            else: # handle the case of this being the last element in the array; so can't jump forward.
                if prefixMax[i] <= suffixMin:
                    ans[i] = prefixMax[i] # a case where backwards jump was possible from the last element.
                else:
                    # couldn't jump backwards; so the answer here is 0. Don't even need to change it,
                    # it's already zero from the beginning.
                    # ans[i] = 0
                    pass

            # After handling ith answer, update suffixMin - include the current value.
            suffixMin = min(suffixMin, nums[i])

        return ans
