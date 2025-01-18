from typing import List

class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        memo = {}

        def calc(index: int, num_walls_remaining: int) -> int:
            if (num_walls_remaining <= 0):
                return 0
            if (index == len(time)):
                return 99999999999999999999999 #any big num works
            if (index, num_walls_remaining) in memo:
                return memo[(index, num_walls_remaining)]

            cost_paid = cost[index] + calc(index + 1, num_walls_remaining - time[index] - 1)
            cost_free = calc(index + 1, num_walls_remaining)
            ans = min(cost_paid, cost_free)
            memo[(index, num_walls_remaining)] = ans
            return ans

        return calc(0, len(time))

def main():
    s = Solution()
    assert(s.paintWalls([1, 2, 3, 2], [1, 2, 3, 2]) == 3)
    assert(s.paintWalls([2, 3, 4, 2], [1, 1, 1, 1]) == 4)
    print("All testcases passed successfully")

if __name__ == "__main__":
    main()
