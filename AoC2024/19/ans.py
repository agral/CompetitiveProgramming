DO_EXAMPLE = False

def can_assemble(target, blocks):
    if len(target) == 0:
        return True
    for block in blocks:
        if target.startswith(block):
            if can_assemble(target[len(block):], blocks):
                return True
    return False

# Quick memoization, only store target as key - blocks never change.
all_assemblies_memo = {}
def all_assemblies(target, blocks):
    if len(target) == 0:
        return 1
    if target in all_assemblies_memo:
        return all_assemblies_memo[target]
    ans = 0
    for block in blocks:
        if target.startswith(block):
            ans += all_assemblies(target[len(block):], blocks)
    all_assemblies_memo[target] = ans
    return ans

def main():
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        lines = [line.strip() for line in file.readlines()]
        blocks = lines[0].split(", ")
        targets = [t for t in lines[2:]]

        a = sum([(1 if can_assemble(t, blocks) else 0) for t in targets])
        b = sum([all_assemblies(t, blocks) for t in targets])
        print(f"Answer A: {a}")
        print(f"Answer B: {b}")


if __name__ == "__main__":
    main()

