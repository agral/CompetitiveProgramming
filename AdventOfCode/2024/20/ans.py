import math

# The wording in this puzzle was SO CONFUSING.

DIRS = [[-1, 0], [0, -1], [1, 0], [0, 1]]
DO_EXAMPLE = False

def manhattan(y1, x1, y2, x2):
    return abs(y1 - y2) + abs(x1 - x2)

def print_track(track):
    for y in range(track["h"]):
        for x in range(track["w"]):
            print(track["m"][y][x], end="")
        print()

def print_standard_path(track):
    for y in range(track["h"]):
        for x in range(track["w"]):
            print(f"{track["p"][y][x]}\t", end="")
        print()

def find_standard_path(track):
    num_turn = 0
    y, x = track["s"][0], track["s"][1]
    path = [[math.inf for _ in row] for row in track["m"]]
    path[y][x] = 0
    while track["m"][y][x] != "E":
        newy, newx = y, x
        for dir in DIRS:
            yy, xx = y + dir[0], x + dir[1]
            if track["m"][yy][xx] != "#" and num_turn < path[yy][xx]:
                newy, newx = yy, xx

        num_turn += 1
        y, x = newy, newx
        path[y][x] = num_turn
    track["p"] = path

def main():
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        map = [[ch for ch in line.strip()] for line in file]
        track = {
            "m": map,
            "h": len(map),
            "w": len(map[0]),
        }
        for y in range(track["h"]):
            for x in range(track["w"]):
                if map[y][x] == "S":
                    track["s"] = [y, x]
                elif map[y][x] == "E":
                    track["e"] = [y, x]
        find_standard_path(track)

        A, B = 0, 0
        for sy in range(track["h"]):
            print(f"\rCalculating answer for row={sy+1}/{track['h']}", end="", flush=True)
            for sx in range(track["w"]):
                if track["p"][sy][sx] < math.inf:
                    for ey in range(track["h"]):
                        for ex in range(track["w"]):
                            if track["p"][ey][ex] < math.inf:
                                dist = manhattan(sy, sx, ey, ex)
                                if dist == 2 and track["p"][ey][ex] - track["p"][sy][sx] - 2 >= 100:
                                    A += 1
                                if dist <= 20 and track["p"][ey][ex] - track["p"][sy][sx] - dist >= 100:
                                    B += 1
        print(f"\nAnswer A: {A}")
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
