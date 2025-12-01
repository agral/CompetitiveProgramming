from typing import List

# Runtime complexity:
# Auxiliary space complexity:
# Subjective level:
# Solved on: 2025-12-01

class Safe:
    def __init__(self):
        self.dial = 50
        self.numZeroes = 0
        self.numCrosses = 0

    # Just for fun: emulate the entire safe. This is the slowest possible option!
    # But it should still be blazing fast.
    def emulate(self, op):
        delta = -1 if op[0] == 'L' else 1
        count = int(op[1:])
        while count != 0:
            count -= 1
            self.dial += delta
            if self.dial == -1:
                self.dial = 99
            elif self.dial == 100:
                self.dial = 0

            if self.dial == 0:
                self.numCrosses += 1

        # the dial has now completed a full rotation by op.
        if self.dial == 0:
            self.numZeroes += 1

    # That was just for part A. emulate() now solves both A and B.
    # Keeping this just for posterity.
    def rotate(self, op):
        count = int(op[1:])
        if op[0] == 'L':
            count = -count # go left instead of going right.
        self.dial += count
        self.dial %= 100
        if self.dial == 0:
            self.numZeroes += 1
        #print(f"Safe points at: {self.dial}")

def main():
    for inputfile in ["example.txt", "input.txt"]:
        print(inputfile)
        ifile = open(inputfile, "r")

        safe = Safe()
        for line in [l.rstrip("\n") for l in ifile.readlines()]:
            #print(line)
            #safe.rotate(line)
            safe.emulate(line)
        print(f"A: {safe.numZeroes}")
        print(f"B: {safe.numCrosses}")

if __name__ == "__main__":
    main()
