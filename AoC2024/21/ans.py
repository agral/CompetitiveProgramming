DO_EXAMPLE = True

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


class GenericPad:
    FORBIDDEN = (None, None)
    LAYOUT = {}
    BEST_PATHS = {} # I have this feeling that I'll soon need to add memoization.

    def __init__(self):
        self.posY, self.posX = None, None

    def get_all_moves_for_keys(self, char_from, char_to):
        pos_from, pos_to = self.LAYOUT[char_from], self.LAYOUT[char_to]
        return self.generate_moves(pos_from, pos_to)

    def get_all_moves_to(self, char_to):
        pos_from = (self.posY, self.posX)
        pos_to = self.LAYOUT[char_to]
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

class Numpad(GenericPad):
    FORBIDDEN = (3, 0)
    LAYOUT = {
        "7": (0, 0), "8": (0, 1), "9": (0, 2),
        "4": (1, 0), "5": (1, 1), "6": (1, 2),
        "1": (2, 0), "2": (2, 1), "3": (2, 2),
                     "0": (3, 1), "A": (3, 2),
    }
    def __init__(self):
        self.posY, self.posX = self.LAYOUT["A"][0], self.LAYOUT["A"][1]

    # unused, but amusing :-)
    def old_move_to(self, key):
        ans = ""
        targetY, targetX = self.LAYOUT[key][0], self.LAYOUT[key][1]
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


class Keypad(GenericPad):
    FORBIDDEN = (0, 0)
    LAYOUT = {
                     "^": (0, 1), "A": (0, 2),
        "<": (1, 0), "v": (1, 1), ">": (1, 2),
    }
    def __init__(self):
        self.posY, self.posX = self.LAYOUT["A"][0], self.LAYOUT["A"][1]

    # again unused, but pretty fly (if it would have worked)
    def old_move_to(self, key):
        ans = ""
        targetY, targetX = self.LAYOUT[key][0], self.LAYOUT[key][1]
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

        if False:
            for code in codes:
                print(code)
                for key in code:
                    seq_numpad = n.get_all_moves_to(key)
                    print(seq_numpad)

        if False:
            for code in codes:
                seq_numpad = "".join([n.move_to(key) for key in code])
                seq_k1 = "".join([k1.move_to(key) for key in seq_numpad])
                seq_k2 = "".join([k2.move_to(key) for key in seq_k1])

                print(f"{seq_k2} ({len(seq_k2)})")
                print(f"{seq_k1}")
                print(f"{seq_numpad}")
                print(f"{code}")
                print()

        # Manual validation of move sequence generators:
        # (manual verification passed)
        moves = n.get_all_moves_for_keys("A", "7")
        print(moves)
        moves = n.get_all_moves_for_keys("1", "A")
        print(moves)
        print()
        moves = k1.get_all_moves_for_keys("A", "<")
        print(moves)
        moves = k1.get_all_moves_for_keys("<", "^")
        print(moves)
        moves = k1.get_all_moves_for_keys("^", ">")
        print(moves)
        moves = k1.get_all_moves_for_keys(">", "A")
        print(moves)
        print()
        moves = k1.get_all_moves_to("<")
        print(moves)
        moves = k1.get_all_moves_to("v")
        print(moves)

if __name__ == "__main__":
    main()
