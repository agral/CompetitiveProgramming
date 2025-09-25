import math

CH_WALL = "#"
CH_TILE = "."
CH_START = "S"
CH_END = "E"

DIRS = [
    [0, 1], # East
    [-1, 0], # North
    [0, -1],  # West
    [1, 0],  # South
]

DO_EXAMPLE = False
DO_EXAMPLE = True

def main():
    A = 0
    B = 0
    sx, sy, ex, ey = 0, 0, 0, 0
    H, W = 0, 0
    dir = DIRS[0]
    map = []

    def print_map():
        for h in range(H):
            for w in range(W):
                if h == sy and w == sx:
                    print(CH_START, end="")
                elif h == ey and w == ex:
                    print(CH_END, end="")
                else:
                    print(map[h][w], end="")
            print()


    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        for line in file:
            line = line.strip()
            mapline = []
            for i in range(len(line)):
                mapline.append("#" if line[i] == "#" else ".")
                if line[i] == CH_START:
                    sy, sx = H, i
                elif line[i] == CH_END:
                    ey, ex = H, i
            map.append(mapline)
            H += 1
        W = len(map[0])
        print_map()

        best = [[math.inf for ch in line] for line in map]
        extra = [[{"steps": math.inf, "turns": math.inf} for _ in line] for line in map]

        def walk(y, x, dir, cost): # nice try, but hits "RecursionError: max recursion depth exceeded"
            if best[y][x] > cost:
                best[y][x] = cost

                forward_y, forward_x = y + DIRS[dir][0], x + DIRS[dir][1]
                if map[forward_y][forward_x] == CH_TILE:
                    walk(forward_y, forward_x, dir, cost + 1)

                dir_turn_left = (dir - 1) % len(DIRS)
                left_y, left_x = y + DIRS[dir_turn_left][0], x + DIRS[dir_turn_left][1]
                if map[left_y][left_x] == CH_TILE:
                    walk(left_y, left_x, dir_turn_left, cost + 1001)

                dir_turn_right = (dir + 1) % len(DIRS)
                right_y, right_x = y + DIRS[dir_turn_right][0], x + DIRS[dir_turn_right][1]
                print(f"ry={right_y}, rx={right_x}")
                if map[right_y][right_x] == CH_TILE:
                    walk(right_y, right_x, dir_turn_right, cost + 1001)

        # Unrolled the above into an iterated version:
        # Note to self: there's this trick in Python:
        # sys.setrecursionlimit(25000)
        # The limit is 1000 by default. The map is 150x150 tiles, so 25000 is more than enough.
        steps_to_walk = [{"y": sy, "x": sx, "dir": 1, "cost": 1000, "steps": 0, "turns": 0}]
        while len(steps_to_walk) > 0:
            s = steps_to_walk[0]
            steps_to_walk = steps_to_walk[1:]
            y, x, dir, cost, steps, turns = s["y"], s["x"], s["dir"], s["cost"], s["steps"], s["turns"]
            if best[y][x] > cost:
                best[y][x] = cost
                extra[y][x]["steps"] = steps
                extra[y][x]["turns"] = turns

                if y == ey and x == ex:
                    continue # when at exit, do not seek further.

                forward_y, forward_x = y + DIRS[dir][0], x + DIRS[dir][1]
                if map[forward_y][forward_x] == CH_TILE:
                    steps_to_walk.append({"y": forward_y, "x": forward_x, "dir": dir, "cost": cost + 1,
                                          "steps": steps + 1, "turns": turns})

                dir_turn_left = (dir - 1) % len(DIRS)
                left_y, left_x = y + DIRS[dir_turn_left][0], x + DIRS[dir_turn_left][1]
                if map[left_y][left_x] == CH_TILE:
                    steps_to_walk.append({"y": left_y, "x": left_x, "dir": dir_turn_left, "cost": cost + 1001,
                                          "steps": steps + 1, "turns": turns + 1})

                dir_turn_right = (dir + 1) % len(DIRS)
                right_y, right_x = y + DIRS[dir_turn_right][0], x + DIRS[dir_turn_right][1]
                if map[right_y][right_x] == CH_TILE:
                    steps_to_walk.append({"y": right_y, "x": right_x, "dir": dir_turn_right, "cost": cost + 1001,
                                          "steps": steps + 1, "turns": turns + 1})

        walk(sy, sx, 0, 0)
        for h in range(H):
            for w in range(W):
                print(f"{best[h][w]}\t", end="")
            print()

        A = best[ey][ex]
        print(f"Answer A: {A}")

        is_path = [[False for _ in line] for line in map]

        steps_b = [{"y": ey, "x": ex, "cost": math.inf}]
        while len(steps_b) > 0:
            s = steps_b[0]
            steps_b = steps_b[1:]
            y, x, cost = s["y"], s["x"], s["cost"]
            if best[y][x] < math.inf and best[y][x] > A:
                continue
            is_path[y][x] = True
            if not is_path[y-1][x] and (best[y-1][x] < best[y][x] or best[y-1][x] == best[y][x] + 999):
                steps_b.append({"y": y-1, "x": x, "cost": best[y][x]})
            if not is_path[y+1][x] and (best[y+1][x] < best[y][x] or best[y+1][x] == best[y][x] + 999):
                steps_b.append({"y": y+1, "x": x, "cost": best[y][x]})
            if not is_path[y][x-1] and (best[y][x-1] < best[y][x] or best[y][x-1] == best[y][x] + 999):
                steps_b.append({"y": y, "x": x-1, "cost": best[y][x]})
            if not is_path[y][x+1] and (best[y][x+1] < best[y][x] or best[y][x+1] == best[y][x] + 999):
                steps_b.append({"y": y, "x": x+1, "cost": best[y][x]})

        for h in range(H):
            for w in range(W):
                print("O" if is_path[h][w] else map[h][w], end="")
                if is_path[h][w]:
                    B += 1
            print()
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
