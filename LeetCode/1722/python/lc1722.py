from typing import List
import collections

class UnionFind:
    def __init__(self, n: int):
        self.id = list(range(n))
        self.rank = [0] * n

    def union_by_rank(self, u: int, v: int) -> None:
        idx_u = self.find(u)
        idx_v = self.find(v)
        if idx_u == idx_v:
            return
        if self.rank[idx_u] < self.rank[idx_v]:
            self.id[idx_u] = idx_v
        elif self.rank[idx_u] > self.rank[idx_v]:
            self.id[idx_v] = idx_u
        else:
            self.id[idx_u] = idx_v
            self.rank[idx_v] += 1

    def find(self, u: int) -> int:
        if self.id[u] != u:
            self.id[u] = self.find(self.id[u])
        return self.id[u]


# Runtime complexity: O(n*logn)
# Auxiliary space complexity: O(n)
# Subjective level: medium+
# Solved on: 2026-04-21
class Solution:
    def minimumHammingDistance(self, source: List[int], target: List[int],
                               allowedSwaps: List[List[int]]) -> int:
        L = len(source)
        counts = [collections.Counter() for _ in range(L)]
        uf = UnionFind(L)
        ans = 0
        for l, r in allowedSwaps:
            uf.union_by_rank(l, r)
        for k in range(L):
            idx = uf.find(k)
            counts[idx][source[k]] += 1

        for k in range(L):
            group = uf.find(k)
            count = counts[group]
            if target[k] not in count:
                ans += 1
            else:
                count[target[k]] -= 1
                if count[target[k]] == 0:
                    del count[target[k]]
        return ans
