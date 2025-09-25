import math

DO_EXAMPLE = False

# Returns traversible tiles that are direct neighbors of (y, x).
# Note: a tile is traversible iff it's in grid's bounds and not corrupted.
def get_neighbors(memory, y, x):
    ans = []
    if (y > 0 and (y-1, x) not in memory["corrupted"]):
        ans.append((y-1, x))
    if (x > 0 and (y, x-1) not in memory["corrupted"]):
        ans.append((y, x-1))
    if (y < memory["H"]-1 and (y+1, x) not in memory["corrupted"]):
        ans.append((y+1, x))
    if (x < memory["W"]-1 and (y, x+1) not in memory["corrupted"]):
        ans.append((y, x+1))
    return ans

# Returns the length of the shortest path from (0, 0) to (H-1, W-1)
# Does gral's bootleg version of Dijkstra's algo
def dijkstra(memory):
    shortest = {}
    for h in range(memory["H"]):
        for w in range(memory["W"]):
            shortest[(h, w)] = math.inf
    shortest[(0, 0)] = 0
    path_length = 0
    visit_now = [(0, 0)]
    while len(visit_now) > 0:
        visit_next = []
        for tile in visit_now:
            for t in get_neighbors(memory, tile[0], tile[1]):
                if shortest[t] == math.inf:
                    shortest[t] = path_length + 1
                    visit_next.append(t)
        path_length += 1
        visit_now = visit_next[:]
    return shortest[(memory["H"] - 1, memory["W"] - 1)]

def print_memory(memory):
    for h in range(memory["H"]):
        for w in range(memory["W"]):
            print("#" if (h, w) in memory["corrupted"] else ".", end="")
        print()

def main():
    memory = {
        "H": 7 if DO_EXAMPLE else 71,
        "W": 7 if DO_EXAMPLE else 71,
        "corrupted": set(),
    }

    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        fb = [line.strip().split(",") for line in file]
        for i in range(12 if DO_EXAMPLE else 1024):
            # want (y,x), not (x,y) pairs; to be consistent with my previous solutions in AoC2024.
            memory["corrupted"].add((int(fb[i][1]), int(fb[i][0]))) 
        #print_memory(memory)
        print(f"Answer A: {dijkstra(memory)}")

        # Part B: continue the byte shower until (h-1, w-1) is no longer reachable from (0, 0).
        for i in range(12 if DO_EXAMPLE else 1024, len(fb)):
            print(f"\r[Calculating B]: Corrupting memory address ({fb[i][1]}, {fb[i][0]}).", end="", flush=True)
            memory["corrupted"].add((int(fb[i][1]), int(fb[i][0]))) 
            if dijkstra(memory) == math.inf:
                print(f"\nAnswer B: {fb[i][0]},{fb[i][1]}") # this is in the original x, y order that Eric wants.
                return 0

if __name__ == "__main__":
    main()
