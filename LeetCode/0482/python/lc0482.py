# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy, but pretty boring / not useful.
# Solved on: 2026-02-15
class Solution:
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        backwards_chars = []
        L = 0

        for i in reversed(range(len(s))):
            if s[i] != '-': # ignore the original dashes completely.
                # Insert the dash when appropriate. We're going from the back, so it's all right.
                if L > 0 and L % k == 0:
                    backwards_chars += '-'
                backwards_chars += s[i].upper()
                L += 1
        return "".join(reversed(backwards_chars))
