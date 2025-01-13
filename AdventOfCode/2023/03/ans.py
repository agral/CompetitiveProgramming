gears = dict()

def is_symbol(char):
    return not (char.isdigit() or char == '.')

def check_for_bordering_symbol(schematic, coords, length):
    h, w = coords
    single_coords = [(h-1, w-1), (h, w-1), (h+1, w-1), (h-1, w+length), (h, w+length), (h+1, w+length)]
    for coord in single_coords:
        if is_symbol(schematic[coord[0]][coord[1]]):
            return (True, schematic[coord[0]][coord[1]], coord)
    for i in range(w, w + length):
        if is_symbol(schematic[h-1][i]):
            return (True, schematic[h-1][i], (h-1, i))
        if is_symbol(schematic[h+1][i]):
            return (True, schematic[h+1][i], (h+1, i))
    return False, 'N', (-1,-1)


def main():
    A = 0
    B = 0
    #with open("example.txt", "r") as file:
    with open("input.txt", "r") as file:
        schematic = [line.strip() for line in file]
        H = len(schematic)
        W = len(schematic[0])
        schematic.insert(0, '.' * (H + 2))
        schematic.append('.' * (H + 2))
        for h in range(1, H + 1):
            schematic[h] = '.' + schematic[h] + '.'

        for h in range(1, H + 1):
            w = 1
            while w < W + 1:
                if schematic[h][w].isdigit():
                    coords = (h, w)
                    length = 1
                    num = int(schematic[h][w])
                    w += 1
                    while (schematic[h][w].isdigit()):
                        num *= 10
                        num += int(schematic[h][w])
                        w += 1
                        length += 1

                    has_symbol, symbol, symbol_coords = check_for_bordering_symbol(schematic, coords, length)
                    if has_symbol:
                        A += num
                        if (symbol == '*'):
                            if not symbol_coords in gears:
                                gears[symbol_coords] = [num,]
                            else:
                                gears[symbol_coords].append(num)
                else:
                    w += 1

        print(A)
        for (_, v) in gears.items():
            if len(v) == 2:
                B += v[0] * v[1]
        print(B)

if __name__ == "__main__":
    main()
