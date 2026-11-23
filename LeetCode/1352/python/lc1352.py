# Runtime complexity: O(1)
# Auxiliary space complexity: O(n) [O(k) to be precise]
# Subjective level: medium
# Solved on: 2026-11-23
class ProductOfNumbers:
    # runtime complexity: O(1), space complexity: O(1)
    def __init__(self):
        self.prefix = [1]

    # runtime complexity: O(1), space complexity: O(n) (each invocation extends `prefix` by 1)
    def add(self, num: int) -> None:
        if num == 0:
            self.prefix = [1]
        else:
            self.prefix.append(self.prefix[-1] * num)

    # runtime complexity: O(1), space complexity: O(2) == O(1) (just two entries accessed)
    def getProduct(self, k: int) -> int:
        L = len(self.prefix)
        if L <= k:
            # there's not enough new values yet after the last "0"; the product is 0.
            return 0
        else:
            # return a prefix product of the last k numbers:
            return self.prefix[L-1] // self.prefix[L-1-k]
