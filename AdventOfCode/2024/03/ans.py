import re
DO_EXAMPLE = False
def main():
    A = 0
    B = 0
    #with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
    with open("example2.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        memory = file.read().replace("\n", " ") # it is not guaranteed that the memory file is one line.

        # search for pattern: "mul([0-9]+,[0-9]+)"
        muls = re.findall(r"mul\(\d+,\d+\)", memory)
        for instruction in muls:
            (lhs,rhs) = [int(x) for x in re.findall(r"\d+", instruction)]
            A += lhs * rhs
        print(f"Answer A: {A}")

        # part B is going to be a big ugly hack :-)
        # store positions of all dos & donts in separate arrays:
        dos = [0] + [m.start() for m in re.finditer(r"do\(\)", memory)]
        donts = [m.start() for m in re.finditer(r"don't\(\)", memory)]

        # then make another array with both dos & donts combined:
        acts = [{"val": x, "op": "do"} for x in dos]
        for y in donts:
            acts.append({"val": y, "op": "dont"})

        # sort all entries so that they are in the order they appear in memory:
        acts.sort(key=lambda e: e["val"])

        # remove duplicates as they have no effect. As a result, dos and donts will appear alternately.
        ops = [acts[0]]
        for i in range(1, len(acts)):
            if acts[i]["op"] != acts[i-1]["op"]:
                ops.append(acts[i])

        # these do-dont pairs map directly to ranges of uncorrupted/usable memory:
        do_memory = ""
        for i in range(0, len(ops) // 2):
            do_memory += memory[ops[2 * i]["val"]:ops[2 * i + 1]["val"]]
        # also if there's a "do" without a terminating "don't", include the memory
        # from the last "do" until the end:
        if len(ops) % 2 == 1 and ops[-1]["op"] == "do":
            do_memory += memory[ops[-1]["val"]:]
        print(memory)
        print(do_memory)

        # now basically repeat what was done in part A, but use do_memory instead of memory.
        muls = re.findall(r"mul\(\d+,\d+\)", do_memory)
        for instruction in muls:
            (lhs,rhs) = [int(x) for x in re.findall(r"\d+", instruction)]
            B += lhs * rhs
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
