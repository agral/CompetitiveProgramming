DO_EXAMPLE = True
DO_EXAMPLE = False

def visualize(robots, H, W):
    counts = [[0 for w in range(W)] for h in range(H)]
    for r in robots:
        counts[r["y"]][r["x"]] += 1
    for h in range(H):
        for w in range(W):
            print(counts[h][w] if counts[h][w] > 0 else ".", end="")
        print()
    print()


def main():
    A = 0
    B = 0

    # Wow, spent 15 minutes debugging this while it turns out I just can't read comprehensively.
    #H = 7 if DO_EXAMPLE else 101
    #W = 11 if DO_EXAMPLE else 103
    H = 7 if DO_EXAMPLE else 103
    W = 11 if DO_EXAMPLE else 101

    MH, MW = (H // 2, W // 2)
    robots = []
    quadrants = [0, 0, 0, 0]
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        for line in file:
            (ps, vs) = [s[2:] for s in line.strip().split(" ")]
            x, y = [int(s) for s in ps.split(",")]
            vx, vy = [int(s) for s in vs.split(",")]
            # print(f"Robot at ({x}, {y}) v=({vx}, {vy})")

            # part A: where will the robots be after 100 seconds?
            x100, y100 = (x + 100 * vx) % W, (y + 100 * vy) % H
            robots.append({"x":x, "y":y, "vx":vx, "vy":vy, "x100":x100, "y100":y100})
            if x100 > MW and y100 < MH:
                quadrants[0] += 1
            elif x100 < MW and y100 < MH:
                quadrants[1] += 1
            elif x100 < MW and y100 > MH:
                quadrants[2] += 1
            elif x100 > MW and y100 > MH:
                quadrants[3] += 1
    A = quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
    print(quadrants)
    print(f"Answer A: {A}")

    # Part B: iterate until you see a christmas tree.
    # let's say when at least 70% of the robots have 2 or more neighbors, this might be a tree.
    # I expect to see a long continuous "string" of robots arranged into a tree-shape.
    neighbor_score = 0
    B = 0
    while neighbor_score < 0.7:
        # move the robots
        B += 1
        for r in robots:
            r["x"] = (r["x"] + r["vx"]) % W
            r["y"] = (r["y"] + r["vy"]) % H

        # map their positions to a 2D array
        counts = [[0 for _ in range(W)] for _ in range(H)]
        for r in robots:
            counts[r["y"]][r["x"]] += 1

        # calculate neighbor score:
        scoring = 0
        for r in robots:
            n = 0 # initially no neighbors
            if r["y"] > 0 and r["x"] > 0 and counts[r["y"]-1][r["x"]-1] > 0:
                n += 1
            if r["y"] > 0 and counts[r["y"]-1][r["x"]] > 0:
                n += 1
            if r["y"] > 0 and r["x"] < W-1 and counts[r["y"]-1][r["x"]+1] > 0:
                n += 1
            if r["x"] > 0 and counts[r["y"]][r["x"]-1] > 0:
                n += 1
            if r["x"] < W-1 and counts[r["y"]][r["x"]+1] > 0:
                n += 1
            if r["y"] < H-1 and r["x"] > 0 and counts[r["y"]+1][r["x"]-1] > 0:
                n += 1
            if r["y"] < H-1 and counts[r["y"]+1][r["x"]] > 0:
                n += 1
            if r["y"] < H-1 and r["x"] < W-1 and counts[r["y"]+1][r["x"]+1] > 0:
                n += 1
            if n >= 2:
                scoring += 1
        neighbor_score = scoring / len(robots)
        print(f"Turn: {B}, score: {neighbor_score:.2f}")

    print(f"Answer B: {B}")
    visualize(robots, H, W) # You have to :-)

if __name__ == "__main__":
    main()
