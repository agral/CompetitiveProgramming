import re
from collections import defaultdict

DO_EXAMPLE = True
#DO_EXAMPLE = False
def main():
    A = 0
    B = 0
    Bmap = defaultdict(int)
    left = []
    right = []
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        for line in file:
            left.append(int(line.split("   ")[0]))
            b = int(line.rsplit("   ")[1])
            right.append(b)
            Bmap[b] += 1

        left.sort()
        right.sort()

    for a in range(len(left)):
        A += abs(right[a] - left[a])
        B += left[a] * Bmap[left[a]]

    print(f"Answer A: {A}")


    print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
