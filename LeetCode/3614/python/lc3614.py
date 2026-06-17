# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: hard.
# Solved on: 2026-06-17
class Solution:
    def processStr(self, s: str, k: int) -> str:
        size = 0
        for ch in s:
            if ch == '*':
                size -= 1
                if size == -1:
                    size = 0
            elif ch == '#':
                size *= 2
            elif ch == '%':
                # not caring about the actual contents of the string yet,
                # just calculating the size. So ignore the '%' character,
                # but don't increase the size by 1 (as done for other chars).
                pass
            else:
                size += 1

        if k >= size:
            return '.'

        reversed_s = reversed(s)
        for ch in reversed_s:
            if ch == '*':
                size += 1
            elif ch == '#':
                size //= 2
                if k >= size:
                    k -= size
            elif ch == '%':
                k = size-k-1 # current string gets reversed, so reverse the k pointer.
            else:
                size -= 1
                if size == k:
                    return ch
        return "ERROR"
