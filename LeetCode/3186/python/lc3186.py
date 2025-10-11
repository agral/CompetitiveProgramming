import collections
from typing import List

# Runtime complexity: O(sort) + O(n) == O(sort)
# Auxiliary space complexity: O(2n) == O(n)
# Subjective level: medium/hard.
class Solution:
    def maximumTotalDamage(self, power: List[int]) -> int:
        c = collections.Counter(power)
        damages = sorted(c.keys())

        # dp[i][isUsed]: maximum damage using damages[0..i]; isUsed is either
        #    0: the last (ith) damage has not been used; or 1: has been used.
        dp = [[0, 0] for _ in range(len(damages))]

        dp[0] = [0, damages[0] * c[damages[0]]]
        for n, damage in enumerate(damages):
            if n == 0:
                continue
            dp[n][0] = max(dp[n-1][0], dp[n-1][1])
            dp[n][1] = damage * c[damage]
            if n >= 1 and (damages[n-1] not in (damage-1, damage-2)):
                dp[n][1] += max(dp[n-1])
            elif n >= 2 and (damages[n-2] != damage - 2):
                dp[n][1] += max(dp[n-2])
            elif n >= 3:
                dp[n][1] += max(dp[n-3])

        return max(dp[-1]) # python's *so* concise w.r.t. golang, gotta love it.
