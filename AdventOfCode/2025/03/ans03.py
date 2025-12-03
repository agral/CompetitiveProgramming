from typing import List

# Runtime complexity: O(n^2)
# Aux space complexity: O(1)
# Solved on: 2025-12-03.

def getBestJoltage(bank: List[int], numOn: int) -> int:
    lastUsedIdx = -1
    lenBank = len(bank)
    ans = 0 # the resulting joltage from picking the batteries.
    for numBattery in range(numOn):
        # ith battery: pick the first largest one from (lastUsedIdx+1) to (len(bank)-(numOn-numBattery)+1).
        battIdx = lastUsedIdx+1
        for i in range(battIdx+1, lenBank-(numOn-numBattery-1)):
            if bank[i] > bank[battIdx]:
                battIdx = i
        # the best battery has been selected, update the resulting joltage and move on to the next battery:
        ans = 10 * ans + bank[battIdx]
        lastUsedIdx = battIdx
    return ans


def main():
    for inputfile in ["example.txt", "input.txt"]:
        print(inputfile)
        ifile = open(inputfile, "r")

        ansA = 0
        ansB = 0
        for line in [l.rstrip("\n") for l in ifile.readlines()]:
            # part A:
            joltages = [int(ch) for ch in line]
            # first digit: the first largest digit that is not the last one.
            idxFirst = 0
            for i in range(1, len(joltages) - 1):
                if joltages[i] > joltages[idxFirst]:
                    idxFirst = i
            # second digit: whatever is the largest digit from idxFirst+1 to the end of the battery.
            idxSecond = idxFirst+1
            for i in range(idxSecond+1, len(joltages)):
                if joltages[i] > joltages[idxSecond]:
                    idxSecond = i
            joltage = 10 * joltages[idxFirst] + joltages[idxSecond]
            ansA += joltage

            ansB += getBestJoltage(joltages, 12)

        print(f"A: {ansA}")
        print(f"B: {ansB}")
        print()

if __name__ == "__main__":
    main()
