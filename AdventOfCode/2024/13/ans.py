import math

#DO_EXAMPLE = True
DO_EXAMPLE = False

def main():
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        lines = [line.strip() for line in file]
        machines = []
        i = 0
        while i < len(lines):
            a = lines[i].split(": ")[1].split(", ")
            ax, ay = int(a[0][1:]), int(a[1][1:])
            b = lines[i+1].split(": ")[1].split(", ")
            bx, by = int(b[0][1:]), int(b[1][1:])
            p = lines[i+2].split(": ")[1].split(", ")
            px, py = int(p[0][2:]), int(p[1][2:])
            machines.append({"ax": ax, "ay": ay, "bx": bx, "by": by, "px": px, "py": py})
            i += 4

            # part A can be brute forced.
            cost = math.inf
            for btna in range(100):
                for btnb in range(100):
                    if btna * ax + btnb * bx == px and btna * ay + btnb * by == py:
                        cost = min(cost, 3 * btna + btnb)
            if cost < math.inf:
                A += cost

            # part B should probably be solved as set of linear equations.
            px += 10000000000000
            py += 10000000000000
            # A * ax + B * bx = px  --\  B = (py - A * ay) / by
            # A * ay + B * by = py  --/  A * ax + ((py - A * ay) / by) * bx = px
            # solving this results with the following equation for A:
            #numA = (px - (bx * py / by)) / (ax - (ay * bx / by))
            # or simplifying a bit:
            x = ax * (bx * py - by * px) / (ay * bx - by * ax)

            # now, only accept solutions that use natural numbers of button presses
            # (e.g. don't allow a solution calling for 2.5 presses of button A):
            numA = int(x / ax)
            numB = int((px - x) / bx)
            if numA * ax + numB * bx == px and numA * ay + numB * by == py:
                B += 3 * numA + numB
        print(f"Answer A: {A}")

        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
