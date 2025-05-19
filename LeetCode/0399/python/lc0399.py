from typing import List
import collections

# Runtime complexity: 
# Auxiliary space complexity: 
class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
        ans = []

        # divisions[i][j] holds the result of dividing i by j
        divisions = collections.defaultdict(dict)
        for (i, j), value in zip(equations, values):
            divisions[i][j] = value
            divisions[j][i] = 1.0 / value

        def perform_division(lhs: str, rhs: str, known: set[str]) -> float:
            if lhs == rhs:
                return 1.0

            known.add(lhs)
            for divisor, value in divisions[lhs].items():
                if divisor in known:
                    continue
                result = perform_division(divisor, rhs, known)
                if result > 0.0:
                    return value * result # since (lhs / divisor) * (divisor / rhs) === (lhs / rhs)
            return -1.0 # denotes an invalid result.

        for lhs, rhs in queries:
            if lhs not in divisions or rhs not in divisions:
                ans.append(-1.0)
            else:
                ans.append(perform_division(lhs, rhs, set()))
        return ans
