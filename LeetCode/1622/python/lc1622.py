# Subjective level: hard. Requires at least passing familiarity with Fermat's work.
# Solved on: 2026-03-15
class Fancy:
    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def __init__(self) -> None:
        self.MOD = 1_000_000_007
        self.values = []
        # Note: each value in self.values actually represents a * v + b.
        # This is required to get actual values in O(1) time after many adds/muls.
        self.a = 1
        self.b = 0

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(v)
    def append(self, val: int) -> None:
        # `(val - b) // a` is calculated & stored.
        # Using Fermat's little theorem,
        #    a^(p-1) == 1 (mod p)
        #    a^(p-2) == a^(-1) (mod p)
        #    -> (val - b) // a == (val - b) * (a ^ (p-2)) (mod p)
        z = ((val - self.b) + self.MOD) % self.MOD

        # note: it's built-in `pow`, not math.pow.
        # https://docs.python.org/3/library/functions.html#pow
        # (the third argument is a modulus, computed much more efficiently than a^b % mod).
        # - so I don't have to write own modPow function, which is nice.
        self.values.append(z * pow(self.a, self.MOD - 2, self.MOD)) 


    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def addAll(self, inc: int) -> None:
        self.b += inc
        self.b %= self.MOD

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def multAll(self, m: int) -> None:
        self.a *= m
        self.a %= self.MOD
        self.b *= m
        self.b %= self.MOD

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(|append()|)
    def getIndex(self, idx: int) -> int:
        if idx >= len(self.values):
            return -1
        return (self.a * self.values[idx] + self.b) % self.MOD
