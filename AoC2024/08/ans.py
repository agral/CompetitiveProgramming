DO_EXAMPLE = True
DO_EXAMPLE = False
def main():
    A = 0
    B = 0
    antennas = {}
    amap = []
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        amap = [line.strip() for line in file]
    H = len(amap)
    W = len(amap[0])

    for h in range(H):
        for w in range(W):
            if amap[h][w] != ".":
                if amap[h][w] not in antennas.keys():
                    antennas[amap[h][w]] = []
                antennas[amap[h][w]].append((h, w))

    antinode_locations = set()
    for (ak, av) in antennas.items():
        for first in range(len(av) - 1):
            for second in range(first + 1, len(av)):
                # get line equation, but we seemingly do not care about
                # the inner antinodes (between two antennas) - only external ones...?
                f = av[first]
                s = av[second]
                dy, dx = s[0] - f[0], s[1] - f[1]
                a = dy / dx
                b = f[0] - (a * f[1])
                # so yeah don't use the line equation for now.

                # first antinode: -dy, -dx from f; second antinote: dy, dx from s.
                a1 = (f[0]-dy, f[1]-dx)
                a2 = (s[0]+dy, s[1]+dx)
                for an in [a1, a2]:
                    if an[0] >= 0 and an[0] < H and an[1] >= 0 and an[1] < W:
                        antinode_locations.add(an)

    b_locations = set()
    for (ak, av) in antennas.items():
        for first in range(len(av) - 1):
            for second in range(first + 1, len(av)):
                f = av[first]
                s = av[second]
                delta = (s[0] - f[0], s[1] - f[1])
                b_locations.add(f)
                b_locations.add(s)
                # generate antinodes from the first one, downwards:
                an = [f[0], f[1]]
                while True:
                    an[0] -= delta[0]
                    an[1] -= delta[1]
                    if an[0] < 0 or an[0] >= H or an[1] < 0 or an[1] >= W:
                        break
                    b_locations.add((an[0], an[1]))
                an = [s[0], s[1]]
                while True:
                    an[0] += delta[0]
                    an[1] += delta[1]
                    if an[0] < 0 or an[0] >= H or an[1] < 0 or an[1] >= W:
                        break
                    b_locations.add((an[0], an[1]))

    A = len(antinode_locations)
    print(f"Answer A: {A}")

    B = len(b_locations)
    print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
