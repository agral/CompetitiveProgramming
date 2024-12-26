from typing import List

class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        # 2D dynamic programming solution (knapsack problem)
        # Runtime complexity: O(nk); n is length of nums, k is sum of all nums and target.
        # Space complexity: O(k); clearly seen in dp's definition on line #13.
        ans = sum(nums)
        if ans < abs(target) or (ans % 2 != target % 2):
            return 0

        def do_knapsack(val: int) -> int:
            dp = [1] + [0] * ans
            for num in nums:
                for k in range(ans, num - 1, -1): # from ans to num inclusively, decreasing.
                    dp[k] += dp[k - num]
            return dp[val]
        return do_knapsack((ans + target) // 2)


    def findTargetSumWays_recursive_tle(self, nums: List[int], target: int) -> int:
        if len(nums) == 0:
            return 1 if (target == 0) else 0
        else:
            num, new_nums = nums[0], nums[1:]
            return self.findTargetSumWays(new_nums, target-num) + self.findTargetSumWays(new_nums, target+num)

def main():
    s = Solution()
    assert(s.findTargetSumWays([1, 1, 1, 1, 1], 3) == 5)
    assert(s.findTargetSumWays([1], 1) == 1)
    assert(s.findTargetSumWays([1, 0], 1) == 2)
    assert(s.findTargetSumWays([19,32,36,7,37,10,44,21,40,39,39,18,5,34,3,40,33,2,46,46], 29) == 5715)
    assert(s.findTargetSumWays([15,42,44,50,34,46,10,27,27,27,27,40,3,35,4,47,32,45,13,46], 26) == 5332)

    print("All test cases have passed.")

if __name__ == "__main__":
    main()
