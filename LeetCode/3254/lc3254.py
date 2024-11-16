from typing import List

class Solution:
    def resultsArray(self, nums: List[int], k: int) -> List[int]:
        ans = []
        t = 0

        for idx, num in enumerate(nums):
            if idx > 0 and num - nums[idx - 1] != 1:
                t = idx
            if idx + 1 >= k:
                ans.append(num if idx - t >= k - 1 else -1)

        return ans

def main():
    s = Solution()
    assert(s.resultsArray([1, 2, 3, 4, 3, 2, 5], 3) == [3, 4, -1, -1, -1])
    assert(s.resultsArray([2, 2, 2, 2, 2], 4) == [-1, -1])
    assert(s.resultsArray([3, 2, 3, 2, 3, 2], 2) == [-1, 3, -1, 3, -1])
    print("all testcases passed")

if __name__ == "__main__":
    main()
