from typing import List

class Solution:
    def findLengthOfShortestSubarray(self, arr: List[int]) -> int:
        left_end = 0
        right_start = len(arr) - 1

        # find left non-decreasing subarray.
        # It starts with 0th element and ends with `left_end`th: [0 ... left_end]
        while left_end < len(arr) - 1 and (arr[left_end + 1] >= arr[left_end]):
            left_end += 1

        # find right non-decreasing subarray; [right_start ... len(arr)-1].
        while right_start > 0 and (arr[right_start - 1] <= arr[right_start]):
            right_start -= 1

        # Now arr[0 ... left_end] is sorted
        #     arr[left_end+1 ... right_start-1] is NOT sorted,
        #     arr[right_start ... len(arr)-1] is sorted.
        # So to make the array sorted, either arr[left_end+1 ... len(arr)-1] (length=len(arr)-1-left_end),
        # or arr[0, right_start-1] (length=right_start) has to be deleted.
        # The shorter of these two is a candidate answer.
        ans = min(right_start, len(arr) - 1 - left_end)

        # but maybe the answer above can be improved by merging the two sorted arrays.
        # such merge can be done by removing arr[l ... r]; where 0 < l <= left_end,
        # right_start <= r < len(arr) - 1. In order to determine l and r,:
        # * start with l=left_end, r = len(arr)-1
        # * move the indices toward 0 (l) and right_start (r).
        # The removal is done greedily: the largest element gets removed.
        # In case of ties, the r index is moved to the left.
        l = left_end
        r = len(arr) - 1
        while (l >= 0 and r >= right_start and l < r):
            if arr[l] <= arr[r]:
                r -= 1
            else:
                l -= 1
            ans = min(ans, r - l)
        return ans

def main():
    s = Solution()
    # Case 1: removing e.g. [10, 4, 2] makes the remaining array sorted.
    assert(s.findLengthOfShortestSubarray([1, 2, 3, 10, 4, 2, 3, 5]) == 3)

    # Case 2: removing e.g. [5, 4, 3, 2] makes the remaining array sorted.
    assert(s.findLengthOfShortestSubarray([5, 4, 3, 2, 1]) == 4)

    # Case 3: the array is already sorted.
    assert(s.findLengthOfShortestSubarray([1, 2, 3]) == 0)

if __name__ == "__main__":
    main()
