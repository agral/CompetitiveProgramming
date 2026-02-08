from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy+
# Solved on: 2026-02-08
class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        token_to_operation = {
            "+": lambda lhs, rhs: lhs+rhs,
            "-": lambda lhs, rhs: lhs-rhs,
            "*": lambda lhs, rhs: lhs*rhs,
            # Note: not //, as this behaves wrong with negative numbers (6//-123 == -1).
            # "/": lambda lhs, rhs: lhs//rhs,
            "/": lambda lhs, rhs: int(lhs/rhs),
        }
        stack = []
        for token in tokens:
            # token is either an operator {"+", "-", "*", "/"}, or an integer in the range [-200, 200].
            if token in token_to_operation:
                rhs = stack.pop()
                lhs = stack.pop()
                stack.append(token_to_operation[token](lhs, rhs))
                # print(f"lhs={lhs}, rhs={rhs}, result={stack[-1]}")
            else:
                stack.append(int(token))
        return stack.pop()
