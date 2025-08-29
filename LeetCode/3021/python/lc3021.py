# Runtime complexity: O(1)
# Auxiliary space complexity: O(1)
# Subjective level: easy.
class Solution:
    def flowerGame(self, n: int, m: int) -> int:
        # x+y has to be even for Alice to win. That's it.
        # when X is even, valid Ys form the sequence 1, 3, 5, ..., m (or m-1, depending on m).
        # when X is odd, valid Ys form the sequence 2, 4, 6, ..., m-1 (or m).

        # Count of even & odd numbers present in the sequence 1, 2, ..., k, for k == either n or m:
        xEven, yEven = n // 2, m // 2
        xOdd, yOdd = (n+1) // 2, (m+1) // 2
        return xEven*yOdd + xOdd*yEven
