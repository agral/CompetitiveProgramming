#!/usr/bin/env python3

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

""" The following code utilizes the nomenclature as stated:

                                .#.
Normal pattern representation:  ..#
                                ###

Sparse pattern representation: ".#./..#/###"
squished pattern representation: ".#...####"
"""

import re

def sparse_to_normal(sparse):
    return [list(row) for row in sparse.split("/")]

def normal_to_sparse(normal):
    return "/".join(["".join(row) for row in normal])

def sparse_to_squished(sparse):
    squished = re.sub("/", "", sparse)
    assert(len(squished) == 4 or len(squished) == 9)
    return squished



def flip(sparse):
    if len(sparse) == 5:
        flipped = [sparse[1], sparse[0], "/", sparse[4], sparse[3]]
    elif len(sparse) == 11:
        flipped =  [sparse[2], sparse[1], sparse[0], "/",
                    sparse[6], sparse[5], sparse[4], "/",
                    sparse[10], sparse[9], sparse[8]]
    else:
        raise ValueError("Illegal length of sparse pattern.")
    return "".join(flipped)


def rot90(sparse):
    if len(sparse) == 5:
        rotated = [sparse[4], sparse[0], "/", sparse[3], sparse[1]]
    elif len(sparse) == 11:
        rotated =  [sparse[8], sparse[4], sparse[0], "/",
                    sparse[9], sparse[5], sparse[1], "/",
                    sparse[10], sparse[6], sparse[2]]
    else:
        raise ValueError("Illegal length of sparse pattern.")
    return "".join(rotated)


def load_expandable_patterns(patterns_file):
    lines = [line.strip() for line in open(patterns_file, "r")]
    expansions = dict()
    for line in lines:
        sparse_pattern, sparse_expansion = line.split(" => ")
        flipped = flip(sparse_pattern)
        expansions[sparse_pattern] = sparse_expansion
        expansions[flipped] = sparse_expansion

        for i in range(3):
            sparse_pattern = rot90(sparse_pattern)
            flipped = rot90(flipped)
            expansions[sparse_pattern] = sparse_expansion
            expansions[flipped] = sparse_expansion

    print("{} expansions loaded.".format(len(expansions)))
    return expansions


class Board:
    def __init__(self, expansions):
        self.size = 3
        self.b = sparse_to_normal(".#./..#/###")
        self.expansions = expansions

    def grow(self):
        new_size = self.size*3//2 if self.size % 2 == 0 else self.size*4//3
        square_size = 2 if self.size % 2 == 0 else 3
        newb = [["."] * new_size for _ in range(new_size)]

        for y in range(self.size // square_size):
            for x in range(self.size // square_size):
                sparse_chunk = self.slice_to_sparse(
                        square_size*x, square_size*y, square_size)
                new_chunk = self.expansions[sparse_chunk]

                # Emplaces the expanded chunk on the new board:
                S = 3 if len(new_chunk) == 11 else 4
                for row in range(S):
                    for col in range(S):
                        newb[S*y+row][S*x+col] = new_chunk[row * (S+1) + col]
        self.size = new_size
        self.b = newb

    def slice_to_sparse(self, x, y, size):
        if size == 2:
            result = "".join([self.b[y][x], self.b[y][x+1], "/",
                              self.b[y+1][x], self.b[y+1][x+1]])
        elif size == 3:
            result = "".join([self.b[y][x], self.b[y][x+1], self.b[y][x+2], "/",
                    self.b[y+1][x], self.b[y+1][x+1], self.b[y+1][x+2], "/",
                    self.b[y+2][x], self.b[y+2][x+1], self.b[y+2][x+2]])
        else:
            raise RuntimeError("slice_to_sparse() with size other than 2 or 3")
        return result

    def print(self):
        print("\n".join("".join(row) for row in self.b))
        print()

    def count_active_cells(self):
        return sum([row.count("#") for row in self.b])


data_files = ["simple_input_21"]
#data_files = ["input_21"]
for data_file in data_files:
    print("Working with input file: {} ...".format(data_file))

    board = Board(load_expandable_patterns(data_file))
    board.print()

    for i in range(5):
        board.grow()
        board.print()

    print("Challenge A: {} cells are alive.".format(board.count_active_cells()))
