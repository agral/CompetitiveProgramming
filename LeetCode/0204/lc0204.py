from typing import List

class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 2:
            return 0
        sieve = [True] * n
        sieve[0] = False
        sieve[1] = False
        limit = int(n ** .5)
        for num in range(2, limit + 1):
            if sieve[num]:
                for composite in range(2 * num, n, num):
                    sieve[composite] = False
        return sum(sieve)

def main():
    s = Solution()
    assert(s.countPrimes(10) == 4)
    assert(s.countPrimes(0) == 0)
    assert(s.countPrimes(1) == 0)
    print("All testcases passed.")

if __name__ == "__main__":
    main()

