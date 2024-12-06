DO_EXAMPLE = True
DO_EXAMPLE = False

def main():
    DIRECTIONS = [(-1, 0), (0, 1), (1, 0), (0, -1)] # Top, Right, Down, Left
    DIR_REPR =   ["^",     ">",    "v",    "<"    ]
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        map = [line.strip() for line in file]
        H = len(map)
        W = len(map[0])
        gy, gx = (-1, -1)
        for h in range(H):
            for w in range(W):
                if map[h][w] == "^":
                    gy, gx = h, w

        y, x = gy, gx
        dir = 0 # guard goes up, initially.
        print(f"guard at [{y}][{x}]")
        visited = [[False for w in range(W)] for h in range(H)]
        visited[y][x] = True
        num_visited = 1

        guard_in_map_bounds = True
        while guard_in_map_bounds:
            ny, nx = y + DIRECTIONS[dir][0], x + DIRECTIONS[dir][1]
            if ny < 0 or ny >= H or nx < 0 or nx >= W:
                guard_in_map_bounds = False
                break
            if map[ny][nx] == "#":
                dir = (dir + 1) % 4 # turn right
            else:
                y, x = ny, nx
                if not visited[y][x]:
                    visited[y][x] = True
                    num_visited += 1

        A = num_visited
        print(f"Answer A: {A}")

        # now store the level map as 2D array of chars, so that I can mutate them:
        charmap = []
        for h in range(H):
            row = []
            for w in range(W):
                row.append(map[h][w])
            charmap.append(row)

        for oh in range(H):
            # Brute forcing will take a long time. Log progress:
            print(f"calculating: {oh}/{H} ({100*oh/H:.2f}%) done")
            for ow in range(W):
                if map[oh][ow] == "#":
                    # don't put obstacles over already existing obstacles, won't change anything.
                    continue

                # setup new obstacle map and guard's route history:
                omap = [row[:] for row in charmap]
                omap[oh][ow] = "#"
                y, x = gy, gx
                dir = 0
                route = [["#" if omap[h][w] == "#" else "." for w in range(W)] for h in range(H)]
                route[gy][gx] = DIR_REPR[dir]

                # simulate guard's path. At any step:
                #   if out-of-map - there's no loop.
                #   if stepped on already stepped tile + direction matches - it's a loop.
                guard_in_map_bounds = True
                is_looping = False
                while guard_in_map_bounds and not is_looping:
                    ny, nx = y + DIRECTIONS[dir][0], x + DIRECTIONS[dir][1]
                    if ny < 0 or ny >= H or nx < 0 or nx >= W:
                        guard_in_map_bounds = False
                        break
                    if omap[ny][nx] == "#":
                        dir = (dir + 1) % 4 # turn right
                    else:
                        y, x = ny, nx
                        if route[y][x] == ".":
                            route[y][x] = DIR_REPR[dir]
                        else:
                            if route[y][x] == DIR_REPR[dir]:
                                # found a loop!
                                is_looping = True
                                B += 1
                                break

        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
