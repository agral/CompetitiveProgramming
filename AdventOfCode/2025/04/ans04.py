def main():
    for inputfile in ["example.txt", "input.txt"]:
        print(inputfile)
        ifile = open(inputfile, "r")

        lines = [l.rstrip("\n") for l in ifile.readlines()]
        H = len(lines)
        W = len(lines[0])
        floor = [['.'] * (W + 2) for h in range(H+2)]
        for h in range(H):
            for w in range(W):
                floor[h+1][w+1] = lines[h][w]

        # part A: count the rolls of paper that have <= 3 adjacent rolls.
        ansA = 0
        for h in range(1, H+1):
            for w in range(1, W+1):
                if floor[h][w] == '@':
                    numNeighbors = 0
                    if floor[h-1][w-1] == '@':
                        numNeighbors += 1
                    if floor[h-1][w  ] == '@':
                        numNeighbors += 1
                    if floor[h-1][w+1] == '@':
                        numNeighbors += 1
                    if floor[h  ][w-1] == '@':
                        numNeighbors += 1
                    if floor[h  ][w+1] == '@':
                        numNeighbors += 1
                    if floor[h+1][w-1] == '@':
                        numNeighbors += 1
                    if floor[h+1][w  ] == '@':
                        numNeighbors += 1
                    if floor[h+1][w+1] == '@':
                        numNeighbors += 1

                    if numNeighbors <= 3:
                        ansA += 1

        # part B: Remove the rolls of paper that have <= 3 adjacent rolls, until no more progress can be done.
        ansB = 0
        isProgressing = True
        while isProgressing:
            isProgressing = False
            for h in range(1, H+1):
                for w in range(1, W+1):
                    if floor[h][w] == '@':
                        numNeighbors = 0
                        if floor[h-1][w-1] == '@':
                            numNeighbors += 1
                        if floor[h-1][w  ] == '@':
                            numNeighbors += 1
                        if floor[h-1][w+1] == '@':
                            numNeighbors += 1
                        if floor[h  ][w-1] == '@':
                            numNeighbors += 1
                        if floor[h  ][w+1] == '@':
                            numNeighbors += 1
                        if floor[h+1][w-1] == '@':
                            numNeighbors += 1
                        if floor[h+1][w  ] == '@':
                            numNeighbors += 1
                        if floor[h+1][w+1] == '@':
                            numNeighbors += 1

                        if numNeighbors <= 3:
                            ansB += 1
                            floor[h][w] = '.'
                            isProgressing = True

        print(f"A: {ansA}")
        print(f"B: {ansB}")

if __name__ == "__main__":
    main()
