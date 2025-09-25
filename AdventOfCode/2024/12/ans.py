DO_EXAMPLE = False
DIRS = [[-1, 0], [0, 1], [1, 0], [0, -1]] # Top, right, bottom, left

def main():
    A = 0
    B = 0
    land = []
    aland = []
    H, W = 0, 0

    def get_fencing_cost(y, x) -> int:
        if y < 0 or y >= H or x < 0 or x >= W or aland[y][x]["used"]:
            return 0

        area = 0
        perimeter = 0
        bulk_perimeter = 0

        # OK, (y, x) has not been processed yet and is in bounds.
        # note: to_visit will only be extended by (y,x) coords in plot bounds.
        to_visit = [[y, x]]

        while len(to_visit) > 0:
            vy, vx = to_visit[0][0], to_visit[0][1]
            to_visit = to_visit[1:]
            if aland[y][x]["plant"] == aland[vy][vx]["plant"]:
                if not aland[vy][vx]["used"]:
                    aland[vy][vx]["used"] = True
                    area += 1
                    delta_per = 4
                    delta_bulk = 0
                    for d in range(len(DIRS)):
                        aland[vy][vx]["f"][d] = True
                        ny, nx = vy + DIRS[d][0], vx + DIRS[d][1]
                        if ny >= 0 and ny < H and nx >= 0 and nx < W:
                            if aland[ny][nx]["plant"] == aland[vy][vx]["plant"]:
                                # no fence between same plants
                                delta_per -= 1
                                aland[vy][vx]["f"][d] = False
                            if not aland[ny][nx]["used"]:
                                to_visit.append([ny, nx])

                    # top & bottom fences bulk: look at cell to the left.
                    if vx > 0 and aland[vy][vx]["plant"] == aland[vy][vx-1]["plant"]:
                        if aland[vy][vx]["f"][0] and aland[vy][vx-1]["f"][0]:
                            delta_bulk -= 1
                        if aland[vy][vx]["f"][2] and aland[vy][vx-1]["f"][2]:
                            delta_bulk -= 1
                    # left & right fences bulk: look at cell to the top.
                    if vy > 0 and aland[vy][vx]["plant"] == aland[vy-1][vx]["plant"]:
                        if aland[vy][vx]["f"][1] and aland[vy-1][vx]["f"][1]:
                            delta_bulk -= 1
                        if aland[vy][vx]["f"][3] and aland[vy-1][vx]["f"][3]:
                            delta_bulk -= 1
                    perimeter += delta_per
                    bulk_perimeter += (delta_per + delta_bulk)

        return area * perimeter

    def get_bulk_fencing_cost(y, x) -> int:
        if y < 0 or y >= H or x < 0 or x >= W or aland[y][x]["used"]:
            return 0

        area = 0
        perimeter = 0
        bulk_perimeter = 0

        # OK, (y, x) has not been processed yet and is in bounds.
        # note: to_visit will only be extended by (y,x) coords in plot bounds.
        to_visit = [[y, x]]

        while len(to_visit) > 0:
            vy, vx = to_visit[0][0], to_visit[0][1]
            to_visit = to_visit[1:]
            if aland[y][x]["plant"] == aland[vy][vx]["plant"]:
                if not aland[vy][vx]["used"]:
                    aland[vy][vx]["used"] = True
                    area += 1
                    delta_per = 4
                    delta_bulk = 0
                    for d in range(len(DIRS)):
                        aland[vy][vx]["f"][d] = True
                        ny, nx = vy + DIRS[d][0], vx + DIRS[d][1]
                        if ny >= 0 and ny < H and nx >= 0 and nx < W:
                            if aland[ny][nx]["plant"] == aland[vy][vx]["plant"]:
                                # no fence between same plants
                                delta_per -= 1
                                aland[vy][vx]["f"][d] = False
                            if not aland[ny][nx]["used"]:
                                to_visit.append([ny, nx])

                    # top & bottom fences bulk: look at cell to the left.
                    if vx > 0 and aland[vy][vx]["plant"] == aland[vy][vx-1]["plant"]:
                        if aland[vy][vx]["f"][0] and aland[vy][vx-1]["f"][0]:
                            delta_bulk -= 1
                        if aland[vy][vx]["f"][2] and aland[vy][vx-1]["f"][2]:
                            delta_bulk -= 1
                    # left & right fences bulk: look at cell to the top.
                    if vy > 0 and aland[vy][vx]["plant"] == aland[vy-1][vx]["plant"]:
                        if aland[vy][vx]["f"][1] and aland[vy-1][vx]["f"][1]:
                            delta_bulk -= 1
                        if aland[vy][vx]["f"][3] and aland[vy-1][vx]["f"][3]:
                            delta_bulk -= 1
                    perimeter += delta_per
                    bulk_perimeter += (delta_per + delta_bulk)

        return area * bulk_perimeter

    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        for line in file:
            row = [s for s in line.strip()]
            land.append(row)
        H = len(land)
        W = len(land[0])
        aland = [[{"used": False, "plant": ch, "f": [False, False, False, False]} for ch in row] for row in land]

        for h in range(H):
            for w in range(W):
                if not aland[h][w]["used"]:
                    A += get_fencing_cost(h, w)


        for h in range(H):
            for w in range(W):
                aland[h][w]["used"] = False
        for h in range(H):
            for w in range(W):
                if not aland[h][w]["used"]:
                    B += get_bulk_fencing_cost(h, w)

        print(f"Answer A: {A}")

        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
