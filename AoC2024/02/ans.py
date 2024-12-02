from typing import List
DO_EXAMPLE = True
#DO_EXAMPLE = False

def check_valid(l: List) -> bool:
    if not(all(l[k] <= l[k-1] for k in range(1, len(l))) or all(l[k] >= l[k-1] for k in range(1, len(l)))):
        return False
    if any(abs(l[k] - l[k-1]) < 1 for k in range(1, len(l))) or any(abs(l[k] - l[k-1]) > 3 for k in range(1, len(l))):
        return False
    return True

def main():
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        for line in file:
            levels = [int(s.strip()) for s in line.split(" ")]
            rev_levels = levels[:]
            rev_levels.reverse()
            is_good = sorted(levels) == levels or sorted(rev_levels) == rev_levels
            for k in range(1, len(levels)):
                diff = abs(levels[k] - levels[k-1])
                is_good = is_good and diff >= 1 and diff <= 3
            if is_good:
                A += 1

            is_b_good = is_good # if it's good for A, it's surely good with relaxed conditions for B.
            for b in range(len(levels)):
                blevels = levels[:b] + levels[(b+1):]
                is_b_good = is_b_good or check_valid(blevels)
            if is_b_good:
                B += 1


        print(f"Answer A: {A}")

        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
