#!/usr/bin/env python3o

"""
Solution to the challenge #21 of the "Advent of Code 2017" series.

MIT License

Copyright (c) 2017 Adam Szczerbiak (szczerbiakadam@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

def Hash(squished_pattern):
    result = 0
    for c in range(len(squished_pattern)):
        result *= 2
        if squished_pattern[c] == "#":
            result += 1
    result += 100000 * len(squished_pattern)
    return result

def unhash(number):
    result = ""
    while number > 0:
        result += "#" if number % 2 == 1 else "."
        number //= 2
    return result[::-1] # returns a reversed string.


patterns = []
expansions = dict()

class Pattern:
    def __init__(self, representation, expansion=None):
        """ Creates a new instance of a Pattern class.

        Parameters:
            size: the size of a tile containing a pattern - 2 or 3.
            representation: a string representing a serialized pattern.
            expansion: a string representing a serialized expansion
                    of the pattern.
        """

        def flip(squished):
            if self.size == 2:
                flipped = [squished[1], squished[0], squished[3], squished[2]]
            else:
                flipped = [squished[2], squished[1], squished[0],
                           squished[5], squished[4], squished[3],
                           squished[8], squished[7], squished[6]]
            return "".join(flipped)


        def rot90(squished):
            if self.size == 2:
                rotated = [squished[3], squished[0], squished[2], squished[1]]
            else:
                rotated = [squished[6], squished[3], squished[0],
                           squished[7], squished[4], squished[1],
                           squished[8], squished[5], squished[2]]
            return "".join(rotated)


        self.size = 2 if len(representation) == 5 else 3

        self.expansion = [list(row) for row in expansion.split("/")]

        self.rows = representation.split("/")
        self.normal = "".join(self.rows)
        flipped = flip(self.normal)
        normal = self.normal
        expansions[flipped] = self.expansion
        expansions[normal] = self.expansion
        for i in range(3):
            flipped = rot90(flipped)
            normal = rot90(normal)
            expansions[flipped] = self.expansion
            expansions[normal] = self.expansion


class Board:
    def __init__(self):
        self.size = 3
        self.b = [[".", "#", "."],
                  [".", ".", "#"],
                  ["#", "#", "#"]]

    def grow(self):
        newsize = self.size * 3 // 2 if self.size % 2 == 0 \
                else self.size * 4 // 3

        print("Newsize: {}".format(newsize))

        skip = 2 if self.size % 2 == 0 else 3
        newb = [["."] * newsize for _ in range(newsize)]

        for y in range(self.size // skip):
            for x in range(self.size // skip):
                chunk = self.serialize_slice(skip*x, skip*y, skip)
                newchunk = expansions[chunk]
                #print("Chunk: ", end="")
                #print(newchunk)

                # Emplaces the expansion chunk on the new board:
                Y = len(newchunk)
                for row in range(Y):
                    X = len(newchunk[row])
                    for col in range(X):
                        _y, _x = y*Y+row, x*X+col
                        #print("self.size=={}, skip={}, X={}, Y={}".format(self.size, skip, X, Y))
                        #print("Row={}, col={}, _y={}, _x={}".format(row, col, _y, _x))
                        newb[y*Y+row][x*X+col] = newchunk[row][col]
                print("Emplaced one chunk.")

        self.size = newsize
        self.b = newb

    def serialize_slice(self, x, y, size):
        if size == 2:
            result = "".join([self.b[y][x], self.b[y][x+1],
                              self.b[y+1][x], self.b[y+1][x+1]])
        elif size == 3:
            result = "".join([self.b[y][x], self.b[y][x+1], self.b[y][x+2],
                              self.b[y+1][x], self.b[y+1][x+1], self.b[y+1][x+2],
                              self.b[y+2][x], self.b[y+2][x+1], self.b[y+2][x+2]
                              ])
        else:
            raise RuntimeError("serialize_slice() with size other than 2 or 3")

        return result

    def print(self):
        print("\n".join("".join(row) for row in self.b))
        print()

    def alive_cells_count(self):
        return sum([row.count("#") for row in self.b])


data_file = "simple_input_21"
data_file = "input_21"
lines = [line.strip() for line in open(data_file, "r")]
for line in lines:
    (pat, exp) = line.split(" => ")
    patterns.append(Pattern(pat, exp))
print(len(expansions), "expansions in the map loaded.")


board = Board()
print("Initial state:")
board.print()
print()

for i in range(5):
    print("Round #{}.".format(i+1))
    board.grow()
    print("Alive cells: {}".format(board.alive_cells_count()))
    board.print()
