from typing import List

# Runtime complexity: O(n^2) (optimal).
# Auxiliary space complexity: 4 * O(mid+1) which is O(r-l+1), generally O(n).
# Subjective level: very hard.
# Solved on: 2026-06-23

# Note: It still *passes* with a runtime of full 12 seconds. What a challenging question!

class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1_000_000_007

        # observation 1: the number of zigzag arrays of length >= 3
        # is 0 when l >= r.

        # observation 2: the number of zigzag arrays explodes somewhat akin
        # to a pascal's triangle (of sorts; as the range gets big, the count
        # of these arrays gets *big* fast).

        mid = r - l + 1
        alpha = [0] * (mid+1)
        beta = [0] * (mid+1)
        alpha_next = [0] * (mid+1)
        beta_next = [0] * (mid+1)

        for i in range(1, mid+1):
            alpha[i] = i - 1
            beta[i] = mid - i

        for i in range(2, n): # so from 2 to n-1 inclusively; or in C: `for int i = 2; i < n; i++`.
            prefixSum = [0] * (mid+2)
            for p in range(1, mid+1):
                prefixSum[p+1] = (prefixSum[p] + beta[p]) % MOD

            suffixSum = [0] * (mid+3)
            for s in range(mid, 0, -1): # C: `for int s = mid; s > 0; s--`.
                suffixSum[s] = (suffixSum[s+1] + alpha[s]) % MOD

            for z in range(1, mid+1):
                alpha_next[z] = prefixSum[z]
                beta_next[z] = suffixSum[z+1]

            alpha, alpha_next = alpha_next, alpha # is this the correct way to swap two lists?
            beta, beta_next = beta_next, beta # if the tests pass, I'll assume I did right here.

        ans = 0
        for i in range(1, mid+1):
            ans = (ans + alpha[i] + beta[i]) % MOD
        return ans
