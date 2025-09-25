import itertools

DO_EXAMPLE = False

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
        ans = self.generate_moves(pos_from, pos_to)
        (self.posY, self.posX) = pos_to
        return ans

    def get_all_moves_to(self, char_to):
        pos_from = (self.posY, self.posX)
        pos_to = self.LAYOUT[char_to]
        ans = self.generate_moves(pos_from, pos_to)
        (self.posY, self.posX) = pos_to
        return ans

    def generate_moves(self, pos_from, pos_to, path_here=""):
        ans = []
        diffY, diffX = pos_to[0] - pos_from[0], pos_to[1] - pos_from[1]
        if pos_from == self.FORBIDDEN:
            return []

        if diffY == 0 and diffX == 0:
            ans.append(path_here + "A")
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

    def find_input(self, desired_output):
        ans = []
        for key in desired_output:
            sequence = self.get_all_moves_to(key)
            if len(ans) == 0:
                ans = sequence[:]
            else:
                ans = [(z[0] + z[1]) for z in list(itertools.product(ans, sequence))]
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

def get_numeric_part(code):
    return int(code[:-1])

def main():
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        codes = [line.strip() for line in file]

        n = Numpad()
        k1 = Keypad()
        k2 = Keypad()

        A, B = 0, 0

        for i, code in enumerate(codes):
            print(f"\rCalculating code #{i+1}/{len(codes)} [{code}]...", end="", flush=True)
            ways_n = n.find_input(code)
            ways_k1 = []
            for way_n in ways_n:
                ways_k1 += k1.find_input(way_n)
            ways_k2 = []
            for way_k1 in ways_k1:
                ways_k2 += k2.find_input(way_k1)
            min_length = min([len(way_k2) for way_k2 in ways_k2])
            A += min_length * get_numeric_part(code)

        print(f"\nAnswer A: {A}")

if __name__ == "__main__":
    main()
