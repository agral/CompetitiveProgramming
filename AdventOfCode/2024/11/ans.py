from typing import List
from collections import defaultdict
DO_EXAMPLE = True
DO_EXAMPLE = False

def evolve_a(stones: List[int]) -> List[int]:
    ans = []
    for s in stones:
        if s == 0:
            ans.append(1)
        elif len(str(s)) % 2 == 0:
            tmp = str(s)
            tmpl = len(tmp)
            s1 = tmp[:tmpl//2]
            s2 = tmp[tmpl//2:]
            ans.append(int(s1))
            ans.append(int(s2))
        else:
            ans.append(s * 2024)
    return ans

cycles = {}
def find_cycle(stone: int) -> List:
    steps = []
    for step_num in range(75):
        steps.append([])
        steps[step_num] = evolve_a(steps[step_num-1])
        if stone in steps[step_num]:
            break
    print(f"{stone}: cycle found in {len(steps)} steps")
    cycles[stone] = {"period": len(steps), "steps": steps}
    return steps


def evolve_b(stones: List[int], num_blinks: int) -> int:
    counts = {s: 1 for s in stones}

    def blink(counts):
        acc = defaultdict(int)
        for i, count in counts.items():
            if i == 0:
                acc[1] += count
            else:
                tmp = str(i)
                tmpl = len(tmp)
                if tmpl % 2 == 0:
                    s1, s2 = int(tmp[:tmpl//2]), int(tmp[tmpl//2:])
                    acc[s1] += count
                    acc[s2] += count
                else:
                    acc[2024 * i] += count
        return acc

    for _ in range(num_blinks):
        counts = blink(counts)

    return sum(counts.values())


def main():
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        stones = []
        for line in file:
            stones = [int(x) for x in line.strip().split(" ")]
        bstones = stones[:]

        for _ in range(25):
            stones = evolve_a(stones)
        A = len(stones)
        print(f"Answer A: {A}")

        # just for verification that the iterative solution works for part A too:
        verify = evolve_b(bstones[:], 25)
        if verify == A:
            print("Verification complete.")
        else:
            print("Iterative solution gives WA for part A")
            return 1

        B = evolve_b(bstones, 75)
        print(f"Answer B: {B}")
        return 0


if __name__ == "__main__":
    main()
