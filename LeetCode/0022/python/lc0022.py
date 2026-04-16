from typing import List

# Runtime complexity: O(2^(2n)), probably. Otherwise 2^n. As n <= 8, it's O(1) anyway ;-)
# Auxiliary space complexity: O(2n) (for the stack)
# Subjective level: medium
# Solved on: 2026-04-16
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        ans = []
        def generate(remainingLeft: int, remainingRight: int, stack: list[str]) -> None:
            if remainingLeft == 0 and remainingRight == 0:
                ans.append("".join(stack))
                return

            if remainingLeft > 0:
                stack.append('(')
                generate(remainingLeft-1, remainingRight, stack)
                stack.pop()
            if remainingLeft < remainingRight:
                stack.append(')')
                generate(remainingLeft, remainingRight-1, stack)
                stack.pop()

        generate(n, n, [])
        return ans


    # Runtime complexity: O(3^(n-1)); but as n <= 8, that's O(1) ;-)
    # Auxiliary space complexity: O(1)
    # However, does not work for n >= 4. Won't work, the approach is wrong, though it may look promising.
    def generateParenthesisNaive_WontWork(self, n: int) -> List[str]:
        ans = []
        def generate(prefix: str, insertAt: int, remaining: int) -> None:
            if remaining == 0:
                ans.append(prefix)
                return
            # insert one pair of () at insertAt, then call generate() twice, with insertAt:
            #  - before the just generated pair, i.e. I()
            #  - inside the just generated pair, i.e. (I)
            #  - to the right of the just generated pair, i.e. ()I
            # and with `remaining` decreased by one.
            newPrefix = prefix[:insertAt] + "()" + prefix[insertAt:]
            generate(newPrefix, insertAt, remaining-1)
            generate(newPrefix, insertAt+1, remaining-1)
            generate(newPrefix, insertAt+2, remaining-1)

        generate("", 0, n)

        # now remove the duplicates by creating a set and converting it back to a list.
        return list(set(ans))
