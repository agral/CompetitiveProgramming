import time
from collections import defaultdict

DO_EXAMPLE = False

# For part B; this has to be global value or otherwise would need to be passed
# to each generate_price_deltas() invocation. Holds total profits for each 4-long delta price sequence.
gains = defaultdict(int)

class Monkey:
    MODP = 16777216 # 2 ** 24
    ANDP = 0b111111111111111111111111
    def __init__(self, secret):
        self.secret = secret

    def generate(self):
        m64 = self.secret * 64 # shl 6
        self.secret ^= m64
        self.secret %= self.MODP
        d32 = self.secret // 32 # shr 5
        self.secret ^= d32
        self.secret %= self.MODP
        m2048 = self.secret * 2048 # shl 11
        self.secret ^= m2048
        self.secret %= self.MODP
        return self.secret

    def generate_fast(self): # turns out it's somewhat slower than generate().
        shl6 = self.secret << 6
        self.secret ^= shl6
        self.secret &= self.ANDP
        shr5 = self.secret >> 5
        self.secret ^= shr5
        self.secret &= self.ANDP
        shl11 = self.secret << 11
        self.secret ^= shl11
        self.secret &= self.ANDP

    def generate_price_deltas(self):
        already_occurred = {}
        prices = [self.generate() % 10 for _ in range(2000)]
        price_deltas = [prices[i+1] - prices[i] for i in range(len(prices)-1)]
        # need to check every consecutive subsequence of deltas of length exactly 4:
        for cssq in range(len(price_deltas) - 3):
            # also can not use array as dict key, but tuples are fine.
            t = tuple(price_deltas[cssq:cssq+4])
            if t not in already_occurred:
                # bought for 0, sell for price at the end of the sequence.
                # also note the extra +1; this is due to price_deltas starting at second price
                # (there is no delta for the first number in the pseudorandom sequence).
                already_occurred[t] = prices[cssq+3+1] 
        for t, gain in already_occurred.items():
            gains[t] += gain

def perf_test(codes):
    v1_start = time.time()
    for c in codes:
        m = Monkey(c)
        for _ in range(2000):
            m.generate()
    v1_end = time.time()
    print(f"generate(): {(v1_end-v1_start) * 1000} ms")

    v2_start = time.time()
    for c in codes:
        m = Monkey(c)
        for _ in range(2000):
            m.generate_fast()
    v2_end = time.time()
    print(f"generate_fast(): {(v2_end-v2_start) * 1000} ms")

def look_for_cycles():
    # There are no cycles during 2000 iterations; let's try bruteforcing part B.
    m = Monkey(123)
    known = set()
    for i in range(1, 2001):
        c = m.generate()
        if c in known:
            print(f"Code {c} repeated at i={i}")
        else:
            known.add(c)
    print(len(known))


def main():
    # Of course the example has to randomly change between parts A and B...
    #with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
    with open("example2.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        codes = [int(line.strip()) for line in file]

        A = 0
        for i, c in enumerate(codes):
            print(f"\rAnalyzing monkey #{i+1}/{len(codes)}...", end="", flush=True)
            m = Monkey(c)
            for _ in range(2000):
                m.generate()
            A += m.secret

            m2 = Monkey(c)
            m2.generate_price_deltas()

        B = max(gains.values())
        print(f"\nAnswer A: {A}")
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
