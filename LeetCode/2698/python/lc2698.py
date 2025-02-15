class Solution:
    def punishmentNumber(self, n: int) -> int:
        def isGoodNumber(acc: int, runningSum: int, numChars: str, idx: int, target: int) -> bool:
            if idx == len(numChars):
                return acc + runningSum == target
            digit = int(numChars[idx])
            return (
                isGoodNumber(acc, 10 * runningSum + digit, numChars, idx+1, target) or
                isGoodNumber(acc + runningSum, digit, numChars, idx+1, target)
            )

        ans = 0
        for num in range(1, n+1):
            if isGoodNumber(0, 0, str(num*num), 0, num):
                ans += (num * num)
        return ans
