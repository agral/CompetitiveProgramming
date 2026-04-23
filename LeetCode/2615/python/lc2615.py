from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium+
# Solved on: 2026-04-23
class Solution:
    def distance(self, nums: List[int]) -> List[int]:
        value_to_index = {}
        # sweep #1: map individual numbers to the indices they appear at in `nums`:
        for idx, value in enumerate(nums):
            if value not in value_to_index:
                value_to_index[value] = [idx]
            else:
                value_to_index[value].append(idx)

        # sweep #2: calculate the actual distance for each entry in `nums`:
        ans = [0] * len(nums)
        for indices in value_to_index.values():
            L = len(indices)
            if L > 1:
                sum_until_this_occurence = sum(indices)
                previous_index = 0
                for idx in range(L):
                    sum_until_this_occurence += (idx-1) * (indices[idx] - previous_index)
                    sum_until_this_occurence -= (L-idx-1) * (indices[idx] - previous_index)
                    previous_index = indices[idx]
                    ans[indices[idx]] = sum_until_this_occurence
        return ans


    # Too complex to pass all the test cases. And not really that elegant. But hey, it works!
    # Runtime complexity: O(n^2)
    # Auxiliary space complexity: O(n)
    # Subjective level: easy
    # Solved on: 2026-04-23
    def distance_naive(self, nums: List[int]) -> List[int]:
        value_to_index = {}
        # sweep #1: map individual numbers to the indices they appear at in `nums`:
        for idx, value in enumerate(nums):
            if value not in value_to_index:
                value_to_index[value] = [idx]
            else:
                value_to_index[value].append(idx)

        # sweep #2: calculate the actual distance for each entry in `nums`:
        ans = [0] * len(nums)
        for idx, value in enumerate(nums):
            for other_idx in value_to_index[value]:
                ans[idx] += abs(idx - other_idx)

        return ans
