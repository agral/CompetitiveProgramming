DO_EXAMPLE = True
DO_EXAMPLE = False

DIRS = [[1, 0], [0, 1], [-1, 0], [0, -1]]

def print_arr(arr):
    for y in range(len(arr)):
        for x in range(len(arr[0])):
            print(f"{arr[y][x]}\t", end="")
        print()


def walk_a(hmap, H, W):
    reachable_nines = [[set() for _ in row] for row in hmap]
    for y in range(H):
        for x in range(W):
            if hmap[y][x] == 9:
                pos = (y, x)
                reachable_nines[y][x].add(pos)
    for h in range(8, -1, -1): # from 8 to 0 with decrement -1
        for y in range(H):
            for x in range(W):
                for dir in DIRS:
                    yy, xx = y + dir[0], x + dir[1]
                    if hmap[y][x] == h and yy >= 0 and yy < H and xx >= 0 and xx < W and hmap[yy][xx] == hmap[y][x] + 1:
                        for n in reachable_nines[yy][xx]:
                            reachable_nines[y][x].add(n)
    ans = 0
    for y in range(H):
        for x in range(W):
            if hmap[y][x] == 0:
                ans += len(reachable_nines[y][x])
    return ans

def walk_b(hmap, H, W):
    scores = [[0 for _ in row] for row in hmap]
    for y in range(H):
        for x in range(W):
            if hmap[y][x] == 9:
                scores[y][x] = 1
    for h in range(8, -1, -1): # from 8 to 0 with decrement -1
        for y in range(H):
            for x in range(W):
                for dir in DIRS:
                    yy, xx = y + dir[0], x + dir[1]
                    if hmap[y][x] == h and yy >= 0 and yy < H and xx >= 0 and xx < W and hmap[yy][xx] == hmap[y][x] + 1:
                        scores[y][x] += scores[yy][xx]
    ans = 0
    for y in range(H):
        for x in range(W):
            if hmap[y][x] == 0:
                ans += scores[y][x]

    return ans



def main():
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        hmap = []
        for line in file:
            heights = [int(char) for char in line.strip()]
            hmap.append(heights)
        H = len(hmap)
        W = len(hmap[0])
        A = walk_a(hmap, H, W)
        print(f"Answer A: {A}")

        B = walk_b(hmap, H, W)
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
