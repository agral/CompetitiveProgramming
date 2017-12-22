#!/usr/bin/env python3

"""
Solution to the challenge #22 of the "Advent of Code 2017" series.

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

class Node:
    INFECTED = "#"
    CLEAN = "."
    WEAKENED = "W"
    FLAGGED = "F"

class Direction:
    DOWN = 0
    LEFT = 1
    UP = 2
    RIGHT = 3
    TOTAL = 4


class ViralBoard:
    def __init__(self, diagram):
        self.board = dict()
        Y = len(diagram)
        self.minY = 0
        self.maxY = Y
        self.minX = 0
        self.maxX = 0
        for y in range(0, Y):
            X = len(diagram[y])
            if X > self.maxX:
                self.maxX = X
            for x in range(0, X):
                self.board[(x, y)] = diagram[y][x]

    def at(self, x, y):
        return self.board.get((x, y), Node.CLEAN)

    def print(self, carrier):
        for y in range(self.minY, self.maxY):
            for x in range(self.minX, self.maxX):
                if y == carrier.y and x == carrier.x:
                    print("*", end="")
                else:
                    print(self.at(x, y), end="")
            print()


class Carrier:
    def __init__(self, viralboard, direction=Direction.UP):
        self.viralboard = viralboard
        self.x = (self.viralboard.maxX - self.viralboard.minX) // 2
        self.y = (self.viralboard.maxY - self.viralboard.minY) // 2
        self.direction = direction
        self.infection_counter = 0

    def move(self):
        if self.direction == Direction.LEFT:
            self.x -= 1
        elif self.direction == Direction.RIGHT:
            self.x += 1
        elif self.direction == Direction.UP:
            self.y -= 1
        elif self.direction == Direction.DOWN:
            self.y += 1

    def extend_board_if_necessary(self):
        self.viralboard.minY = min(self.viralboard.minY, self.y)
        self.viralboard.maxY = max(self.viralboard.maxY, self.y+1)
        self.viralboard.minX = min(self.viralboard.minX, self.x)
        self.viralboard.maxX = max(self.viralboard.maxX, self.x+1)

    def step(self):
        # Turns right if current node is infected, left otherwise. Then
        # toggles the infection in the current cell (infected <-> clean),
        if self.viralboard.at(self.x, self.y) == Node.INFECTED:
            self.direction = (self.direction + Direction.TOTAL + 1) % Direction.TOTAL
            self.viralboard.board[(self.x, self.y)] = Node.CLEAN
        else:
            self.direction = (self.direction + Direction.TOTAL - 1) % Direction.TOTAL
            self.viralboard.board[(self.x, self.y)] = Node.INFECTED
            self.infection_counter += 1

        self.viralboard.minY = min(self.viralboard.minY, self.y)
        self.viralboard.maxY = max(self.viralboard.maxY, self.y+1)
        self.viralboard.minX = min(self.viralboard.minX, self.x)
        self.viralboard.maxX = max(self.viralboard.maxX, self.x+1)

        self.extend_board_if_necessary()
        self.move()


class EvolvedCarrier(Carrier):
    Next = {Node.CLEAN: Node.WEAKENED,
            Node.WEAKENED: Node.INFECTED,
            Node.INFECTED: Node.FLAGGED,
            Node.FLAGGED: Node.CLEAN}
    def step(self):
        # Decides which way to turn based on the value of current node:
        current_node = self.viralboard.at(self.x, self.y)
        if current_node == Node.CLEAN: # Turns left on clean nodes:
            self.direction = (self.direction + Direction.TOTAL - 1) % Direction.TOTAL
        elif current_node == Node.INFECTED: # Turns right on infected nodes:
            self.direction = (self.direction + Direction.TOTAL + 1) % Direction.TOTAL
        elif current_node == Node.FLAGGED: # Reverses direction on flagged nodes:
            self.direction = (self.direction + (Direction.TOTAL // 2)) % Direction.TOTAL
        # (EvolvedCarrier does not change its direction on weakened nodes)

        # Modifies the node according to its evolved logic:
        self.viralboard.board[(self.x, self.y)] = \
                EvolvedCarrier.Next[current_node]
        if current_node == Node.WEAKENED:
            self.infection_counter += 1

        # Finally moves one step forward (inherited from base Carrier class):
        self.extend_board_if_necessary()
        self.move()



data_files = ["sample_input_22", "input_22"]
for data_file in data_files:
    print("Reading input file: {}".format(data_file))
    map_of_memory = [list(line.strip()) for line in open(data_file, "r")]
    board = ViralBoard(map_of_memory)
    carrier = Carrier(board)

    # Challenge A:
    STEPS = 10000
    for i in range(STEPS):
        carrier.step()
    print("[Challenge A] Total infections: {}".format(
            carrier.infection_counter))

    # Challenge B:
    board = ViralBoard(map_of_memory)
    carrier = EvolvedCarrier(board)
    STEPS = 10000000

    for i in range(STEPS):
        if i % (STEPS // 1000) == 0:
            print("\rEmulating... ({:.2f}% done)".format(
                    100 * (i+1) / STEPS), end="")
        carrier.step()
    print("\n[Challenge B] Total infections: {}".format(
            carrier.infection_counter))
