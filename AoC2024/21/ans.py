DO_EXAMPLE = True

DIRS = [[-1, 0], [0, 1], [1, 0], [0, -1]] # North, East, South, West

ARROWS = {
    "y": {
        -1: "^",
        1: "v",
    },
    "x": {
        -1: "<",
        1: ">",
    }
}

NUM_KEYPAD = {
    "7": (0, 0), "8": (0, 1), "9": (0, 2),
    "4": (1, 0), "5": (1, 1), "6": (1, 2),
    "1": (2, 0), "2": (2, 1), "3": (2, 2),
                 "0": (3, 1), "A": (3, 2),
}

DIR_KEYPAD = {
                 "^": (0, 1), "A": (0, 2),
    "<": (1, 0), "v": (1, 1), ">": (1, 2),
}

def manhattan(y_start, x_start, y_end, x_end):
    return abs(y_end - y_start) + abs(x_end - x_start)

class Numpad:
    FORBIDDEN = (3, 0)
    def __init__(self):
        self.posY, self.posX = NUM_KEYPAD["A"][0], NUM_KEYPAD["A"][1]

    def move_to(self, key):
        ans = ""
        targetY, targetX = NUM_KEYPAD[key][0], NUM_KEYPAD[key][1]
        diffY, diffX = targetY - self.posY, targetX - self.posX
        if diffY == 0 and diffX == 0:
            return "A"
        stepY, stepX = 1 if diffY > 0 else -1, 1 if diffX > 0 else -1
        # Forbidden position: (3, 0)
        # Perform the following sequence of moves to never fly over the forbidden position:
        # 1. Equalize vertically, but don't go to 3rd row (stop at 2nd row if applicable).
        # 2. Equalize horizontally.
        # 3. Go to 3rd row if need to go there.
        while (targetY < 3 and diffY != 0) or diffY > 1:
            self.posY += stepY
            diffY -= stepY
            ans += ARROWS["y"][stepY]
        while (diffX != 0):
            self.posX += stepX
            diffX -= stepX
            ans += ARROWS["x"][stepX]
        if (targetY == 3 and self.posY == 2):
            self.posY = 3
            diffY = 0
            ans += ARROWS["y"][1]
        ans += "A"
        return ans

    def get_all_moves_for_keys(self, char_from, char_to):
        pos_from, pos_to = NUM_KEYPAD[char_from], NUM_KEYPAD[char_to]
        return self.generate_moves(pos_from, pos_to)

    def generate_moves(self, pos_from, pos_to, path_here=""):
        ans = []
        diffY, diffX = pos_to[0] - pos_from[0], pos_to[1] - pos_from[1]
        if pos_from == self.FORBIDDEN:
            return []

        if diffY == 0 and diffX == 0:
            ans.append(path_here)
        else:
            if diffY > 0:
                ans += (self.generate_moves((pos_from[0]+1, pos_from[1]), pos_to, path_here + "v"))
            elif diffY < 0:
                ans += (self.generate_moves((pos_from[0]-1, pos_from[1]), pos_to, path_here + "^"))
            if diffX > 0:
                ans += (self.generate_moves((pos_from[0], pos_from[1]+1), pos_to, path_here + ">"))
            elif diffX < 0:
                ans += (self.generate_moves((pos_from[0], pos_from[1]-1), pos_to, path_here + "<"))

        return ans


class Keypad:
    def __init__(self):
        self.posY, self.posX = DIR_KEYPAD["A"][0], DIR_KEYPAD["A"][1]

    def move_to(self, key):
        ans = ""
        targetY, targetX = DIR_KEYPAD[key][0], DIR_KEYPAD[key][1]
        diffY, diffX = targetY - self.posY, targetX - self.posX
        if diffY == 0 and diffX == 0:
            return "A"
        # Forbidden position: (0, 0).
        # This time perform the following maneuvers to never fly over the forbidden position:
        # 1. Equalize vertically, if current pos is on 1st or 2nd row
        if self.posX > 0 and targetY != self.posY:
            if self.posY < targetY:
                self.posY += 1
                ans += ARROWS["y"][1]
            else:
                self.posY -= 1
                ans += ARROWS["y"][-1]
        # 2. Equalize horizontally; this is now safe
        while self.posX != targetX:
            if self.posX < targetX:
                self.posX += 1
                ans += ARROWS["x"][1]
            else:
                self.posX -= 1
                ans += ARROWS["x"][-1]
        # 3. Equalize vertically once again; this is to account for case when first step didn't fire
        # because starting position was on row #0 with forbidden patch.
        if targetY != self.posY:
            if self.posY < targetY:
                self.posY += 1
                ans += ARROWS["y"][1]
            else:
                self.posY -= 1
                ans += ARROWS["y"][-1]
        ans += "A"
        return ans

def main():
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        codes = [line.strip() for line in file]

        n = Numpad()
        k1 = Keypad()
        k2 = Keypad()
        for code in codes:
            seq_numpad = "".join([n.move_to(key) for key in code])
            seq_k1 = "".join([k1.move_to(key) for key in seq_numpad])
            seq_k2 = "".join([k2.move_to(key) for key in seq_k1])

            print(f"{seq_k2} ({len(seq_k2)})")
            print(f"{seq_k1}")
            print(f"{seq_numpad}")
            print(f"{code}")
            print()

        moves = n.get_all_moves_for_keys("A", "1")
        print(moves)
        moves = n.get_all_moves_for_keys("1", "A")
        print(moves)

if __name__ == "__main__":
    main()
