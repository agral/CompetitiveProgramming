DO_EXAMPLE = not True
def main():
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        p = [line.strip() for line in file] # p is for word puzzle
        H = len(p)
        W = len(p[0])
        for y in range(H):
            for x in range(W - 3):
                # row-aligned, forwards:
                if (p[y][x] == "X" and p[y][x+1] == "M" and p[y][x+2] == "A" and p[y][x+3] == "S"):
                    A += 1
                # row-aligned, backwards:
                if (p[y][W-1-x] == "X" and p[y][W-2-x] == "M" and p[y][W-3-x] == "A" and p[y][W-4-x] == "S"):
                    A += 1
        for y in range(H - 3):
            for x in range(W):
                # column-aligned, forwards:
                if (p[y][x] == "X" and p[y+1][x] == "M" and p[y+2][x] == "A" and p[y+3][x] == "S"):
                    A += 1
                # column-aligned, backwards:
                if (p[H-1-y][x] == "X" and p[H-2-y][x] == "M" and p[H-3-y][x] == "A" and p[H-4-y][x] == "S"):
                    A += 1
        for y in range(H - 3):
            for x in range(W - 3):
                # diagonal \, to bottom-right:
                if (p[y][x] == "X" and p[y+1][x+1] == "M" and p[y+2][x+2] == "A" and p[y+3][x+3] == "S"):
                    A += 1
                # diagonal \, to top-left:
                if (p[H-1-y][W-1-x] == "X" and p[H-2-y][W-2-x] == "M" and p[H-3-y][W-3-x] == "A" and p[H-4-y][W-4-x] == "S"):
                    A += 1
                # diagonal /, to top-right:
                if (p[H-1-y][x] == "X" and p[H-2-y][x+1] == "M" and p[H-3-y][x+2] == "A" and p[H-4-y][x+3] == "S"):
                    A += 1
                # diagonal /, to bottom-left:
                if (p[y][W-1-x] == "X" and p[y+1][W-2-x] == "M" and p[y+2][W-3-x] == "A" and p[y+3][W-4-x] == "S"):
                    A += 1

        # now searching for X-MAS:
        for y in range(H-2):
            for x in range(W-2):
                # X-shaped, \ to bottom-right, / to top-right
                if ((p[y][x] == "M" and p[y+1][x+1] == "A" and p[y+2][x+2] == "S") and
                   (p[y+2][x] == "M" and p[y+1][x+1] == "A" and p[y][x+2] == "S")):
                    B += 1
                # X-shaped, \ to bottom-right, / to bottom-left
                if ((p[y][x] == "M" and p[y+1][x+1] == "A" and p[y+2][x+2] == "S") and
                   (p[y][x+2] == "M" and p[y+1][x+1] == "A" and p[y+2][x] == "S")):
                    B += 1
                # X-shaped, \ to top-left, / to top-right
                if ((p[y+2][x+2] == "M" and p[y+1][x+1] == "A" and p[y][x] == "S") and
                   (p[y+2][x] == "M" and p[y+1][x+1] == "A" and p[y][x+2] == "S")):
                    B += 1
                # X-shaped, \ to top-left, / to bottom-left
                if ((p[y+2][x+2] == "M" and p[y+1][x+1] == "A" and p[y][x] == "S") and
                   (p[y][x+2] == "M" and p[y+1][x+1] == "A" and p[y+2][x] == "S")):
                    B += 1



        print(f"Answer A: {A}")

        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
