#!/usr/bin/env python3

"""
Solution to the challenge #19 of the "Advent of Code 2017" series.

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

import os
import time

class Direction:
    DOWN = "v"
    LEFT = "<"
    RIGHT = ">"
    UP = "^"

class PacketRouter:
    def __init__(self, map_file):
        self.load_map(map_file)
        self.reset()

    def load_map(self, map_file):
        self.map = []
        for line in open(map_file, "r"):
            self.map.append(" " + line.rstrip("\n") + " ")
        self.width = len(self.map[0]) + 2
        self.height = len(self.map) + 2
        self.map.append(" " * self.width)
        self.map.insert(0, " " * self.width)

        self.startx = self.map[1].index("|")
        self.starty = 1


    def reset(self):
        # Current position of the packet:
        self.x, self.y = self.startx, self.starty

        self.steps_taken = 0
        self.visited_letters = []
        self.direction = Direction.DOWN


    def route(self, showoff=False):
        while self.map[self.y][self.x] != " ":
            # Goes one step in the current direction:
            if self.direction == Direction.DOWN:
                self.y += 1
            elif self.direction == Direction.UP:
                self.y -= 1
            elif self.direction == Direction.RIGHT:
                self.x += 1
            elif self.direction == Direction.LEFT:
                self.x -= 1
            self.steps_taken += 1

            # "Visits" the tile:
            if self.map[self.y][self.x] in ("|", "-"):
                pass
            elif self.map[self.y][self.x] == "+":
                # Chooses one of the available directions, other than the one
                # that the packet came from, then steps there.
                if   self.direction != Direction.DOWN \
                    and not self.map[self.y-1][self.x] == " ":
                        self.direction = Direction.UP
                elif self.direction != Direction.LEFT \
                    and not self.map[self.y][self.x+1] == " ":
                        self.direction = Direction.RIGHT
                elif self.direction != Direction.UP \
                    and not self.map[self.y+1][self.x] == " ":
                        self.direction = Direction.DOWN
                elif self.direction != Direction.RIGHT \
                    and not self.map[self.y][self.x-1] == " ":
                        self.direction = Direction.LEFT
                else:
                    raise RuntimeError("Lost on junction at ({},{})".format(
                            self.x, self.y))
            else:
                self.visited_letters.append(self.map[self.y][self.x])

            if showoff:
                os.system("clear")
                self.visualize()
                time.sleep(0.1)

        print("Finished routing the packet.")


    def visualize(self):
        for y in range(len(self.map)):
            for x in range(len(self.map[y])):
                if x == self.x and y == self.y:
                    print(self.direction, end="")
                else:
                    print(self.map[y][x], end="")
            print()


data_files = ["sample_input_19", "input_19"]

for data_file in data_files:
    pr = PacketRouter(data_file)
    pr.route(False)
    print("Visited letters:", "".join(pr.visited_letters))
    print("Steps taken: {}".format(pr.steps_taken))
