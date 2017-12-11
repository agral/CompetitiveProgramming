#!/usr/bin/env python3

"""
Solution to the challenge #11 of the "Advent of Code 2017" series.

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

def get_distance(moves):
    individual_moves = { "nw": 0, "n": 0, "ne": 0, "sw": 0, "s": 0, "se": 0 }
    for move in moves:
        individual_moves[move] += 1

    while True: # emulates a do-while loop
        # movements on the NW-SE, NE-SW, N-S diagonals cancel out each other:
        common_nwse = min(individual_moves["nw"], individual_moves["se"])
        common_nesw = min(individual_moves["ne"], individual_moves["sw"])
        common_ns = min(individual_moves["n"],  individual_moves["s"])
        individual_moves["nw"] -= common_nwse
        individual_moves["se"] -= common_nwse
        individual_moves["ne"] -= common_nesw
        individual_moves["sw"] -= common_nesw
        individual_moves["n"] -= common_ns
        individual_moves["s"] -= common_ns

        # movements on the NE+NW / SE+SW can be treated as one movement N/S:
        common_nenw = min(individual_moves["ne"], individual_moves["nw"])
        common_sesw = min(individual_moves["se"], individual_moves["sw"])
        individual_moves["ne"] -= common_nenw
        individual_moves["nw"] -= common_nenw
        individual_moves["n"] += common_nenw
        individual_moves["se"] -= common_sesw
        individual_moves["sw"] -= common_sesw
        individual_moves["s"] += common_sesw

        # movements on the NW+S / N+SE can be treated as one movement SW/NE:
        common_nws = min(individual_moves["nw"], individual_moves["s"])
        common_nse = min(individual_moves["n"], individual_moves["se"])
        individual_moves["nw"] -= common_nws
        individual_moves["s"] -= common_nws
        individual_moves["sw"] += common_nws
        individual_moves["n"] -= common_nse
        individual_moves["se"] -= common_nse
        individual_moves["ne"] += common_nse

        # movements on the N+SW / NE+S can be treated as one movement NW/SE:
        common_nsw = min(individual_moves["n"], individual_moves["sw"])
        common_nes = min(individual_moves["ne"], individual_moves["s"])
        individual_moves["n"] -= common_nsw
        individual_moves["sw"] -= common_nsw
        individual_moves["nw"] += common_nsw
        individual_moves["ne"] -= common_nes
        individual_moves["s"] -= common_nes
        individual_moves["se"] += common_nes

        total_optimizations = common_nwse + common_nesw + common_ns + \
                common_nenw + common_sesw + common_nws + common_nes + \
                common_nsw + common_nse
        if total_optimizations == 0:
            break

    moves_count = individual_moves["nw"] + individual_moves["sw"] + \
                  individual_moves["ne"] + individual_moves["se"] + \
                  individual_moves["n"] + individual_moves["s"]

    #print("Moves: ", end="")
    #print(individual_moves)
    return moves_count



data_files = [ "sample_input_11", "input_11" ]

for individual_file in data_files:
    print("Processing file: {}".format(individual_file))
    data = [line.strip() for line in open(individual_file, "r")]
    for line in data:
        moves = line.split(",")
        d = get_distance(moves)
        print("Absolute distance: {}.".format(d))

        # Challenge #02: How many tiles away is the furthest he's ever been?
        # Brute-force solution:
        max_distance = 0
        next_moves = moves[:]
        moves = []
        while len(next_moves) > 0:
            this_move = next_moves.pop(0)
            moves.append(this_move)
            d = get_distance(moves)
            if d > max_distance:
                max_distance = d
        print("Max distance: {}.".format(max_distance))

