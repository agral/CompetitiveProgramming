DO_EXAMPLE = False
#DO_EXAMPLE = True

class Monkey:
    MODP = 16777216
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


def main():
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        codes = [int(line.strip()) for line in file]

        A = 0
        for c in codes:
            m = Monkey(c)
            for _ in range(2000):
                m.generate()
            A += m.secret
        print(f"Answer A: {A}")

if __name__ == "__main__":
    main()

