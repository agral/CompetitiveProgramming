from typing import List

# Runtime complexity: O(sort) == O(nlogn)
# Auxiliary space complexity: O(n)
# Subjective level: easy+
# Solved on: 2026-02-19
class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        # Sort the array. It is guaranteed that at least two elements are present in the array,
        # so for the moment assume that the first two are part of the answer. Calculate their difference.
        arr.sort()
        ans = [[arr[0], arr[1]]]
        minDiff = arr[1] - arr[0]

        # Iterate over the rest of neighboring pairs, calculating their difference.
        for i in range(2, len(arr)):
            diff = arr[i] - arr[i-1]
            # If this pair's difference is lower than up-to-now known minDiff,
            # replace the stored answer with this pair, and update the minDiff
            # to reflect the just found lowest-so-far diff:
            if diff < minDiff:
                ans = [[arr[i-1], arr[i]]]
                minDiff = diff

            # And if this pair's difference equals the currently known minDiff,
            # Add this pair to the answer list.
            elif diff == minDiff:
                ans.append([arr[i-1], arr[i]])
        return ans
