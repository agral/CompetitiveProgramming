import heapq
import itertools
from typing import List

class Solution:
    def putMarbles(self, weights: List[int], k: int) -> int:
        arr = [leftwise + rightwise for leftwise, rightwise in itertools.pairwise(weights)]
        largest = sum(heapq.nlargest(k-1, arr))
        smallest = sum(heapq.nsmallest(k-1, arr))
        return largest - smallest
